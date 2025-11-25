package protocol

import (
	"math/big"
	"testing"

	poseidon "github.com/iden3/go-iden3-crypto/v2/poseidon"
)

// TestPadLenFromMaxDegreeAndZeroArray checks the math for padLen and zeroArray.
func TestPadLenFromMaxDegreeAndZeroArray(t *testing.T) {
	tests := []struct {
		maxDeg  uint64
		wantLen int
	}{
		// For 0, we defined a special case: still at least one full block.
		{maxDeg: 0, wantLen: 1 + 15},
		{maxDeg: 1, wantLen: 1 + 15},  // ceil(1/15) = 1
		{maxDeg: 15, wantLen: 1 + 15}, // ceil(15/15) = 1
		{maxDeg: 16, wantLen: 1 + 30}, // ceil(16/15) = 2
		{maxDeg: 30, wantLen: 1 + 30}, // ceil(30/15) = 2
		{maxDeg: 31, wantLen: 1 + 45}, // ceil(31/15) = 3
	}

	for _, tt := range tests {
		gotLen := padLenFromMaxDegree(tt.maxDeg)
		if gotLen != tt.wantLen {
			t.Fatalf("padLenFromMaxDegree(%d) = %d, want %d",
				tt.maxDeg, gotLen, tt.wantLen)
		}

		za := zeroArray(tt.maxDeg)
		if len(za) != tt.wantLen {
			t.Fatalf("zeroArray(%d) len = %d, want %d",
				tt.maxDeg, len(za), tt.wantLen)
		}
		for i, v := range za {
			if v != 0 {
				t.Fatalf("zeroArray(%d)[%d] = %d, want 0",
					tt.maxDeg, i, v)
			}
		}
	}
}

// TestNbrArrayHasherSingleRound checks that for maxDeg <= 15 (numR = 1)
// nbrArrayHasher is equivalent to a single Poseidon.Hash on [deg, u0..u14].
func TestNbrArrayHasherSingleRound(t *testing.T) {
	maxDeg := uint64(15) // numR = 1, padLen = 16
	padLen := padLenFromMaxDegree(maxDeg)
	if padLen != 16 {
		t.Fatalf("expected padLen 16 for maxDeg=15, got %d", padLen)
	}

	// Neighbours: [5, 9] => deg=2.
	compact := buildNbrDataCompact([]uint64{5, 9}) // [2,5,9]
	nbrData := padNbrData(compact, maxDeg)         // len=16, rest zeros

	// Hash using nbrArrayHasher (production code).
	h1 := nbrArrayHasher(nbrData)
	if h1 == nil {
		t.Fatalf("nbrArrayHasher returned nil hash")
	}

	// Manually compute Poseidon_16([deg, u0..u14]) directly using iden3 poseidon.
	block := make([]*big.Int, 16)
	for i := 0; i < 16; i++ {
		block[i] = new(big.Int).SetUint64(nbrData[i])
	}
	h2, err := poseidon.Hash(block)
	if err != nil {
		t.Fatalf("poseidon.Hash failed: %v", err)
	}

	if h1.Cmp(h2) != 0 {
		t.Fatalf("single-round mismatch:\n  nbrArrayHasher = %s\n  poseidon.Hash   = %s",
			h1.Text(16), h2.Text(16))
	}

	// Determinism check.
	h3 := nbrArrayHasher(nbrData)
	if h1.Cmp(h3) != 0 {
		t.Fatalf("hash not deterministic: h1=%s, h3=%s", h1.Text(16), h3.Text(16))
	}
}

// TestNbrArrayHasherMultiRound checks that when numR > 1, the chaining logic
// [deg, u0..u14] then [acc, next 15] matches a manual computation using poseidon.Hash.
func TestNbrArrayHasherMultiRound(t *testing.T) {
	maxDeg := uint64(16) // ceil(16/15)=2 => numR=2, padLen=1+30=31
	padLen := padLenFromMaxDegree(maxDeg)
	if padLen != 31 {
		t.Fatalf("expected padLen 31 for maxDeg=16, got %d", padLen)
	}

	// Let's make deg=16, neighbours=[1..16].
	nbrs := make([]uint64, 16)
	for i := 0; i < 16; i++ {
		nbrs[i] = uint64(i + 1)
	}
	compact := buildNbrDataCompact(nbrs) // [16,1,2,...,16]
	nbrData := padNbrData(compact, maxDeg)
	if len(nbrData) != padLen {
		t.Fatalf("padNbrData length = %d, want %d", len(nbrData), padLen)
	}

	// Production hash.
	hProd := nbrArrayHasher(nbrData)

	// Manual multi-round Poseidon chaining:
	//  - B0 = [deg, u0..u14]
	//  - acc0 = Poseidon(B0)
	//  - B1 = [acc0, u15, 0,0,...] (since padLen is only 31)
	block := make([]*big.Int, 16)

	// B0
	for i := 0; i < 16; i++ {
		block[i] = new(big.Int).SetUint64(nbrData[i])
	}
	acc0, err := poseidon.Hash(block)
	if err != nil {
		t.Fatalf("poseidon.Hash(B0) failed: %v", err)
	}

	// B1
	block[0] = new(big.Int).Set(acc0)
	for j := 1; j < 16; j++ {
		idx := 16 + (j - 1) // 16..30
		if idx < len(nbrData) {
			block[j] = new(big.Int).SetUint64(nbrData[idx])
		} else {
			block[j] = big.NewInt(0)
		}
	}
	acc1, err := poseidon.Hash(block)
	if err != nil {
		t.Fatalf("poseidon.Hash(B1) failed: %v", err)
	}

	if hProd.Cmp(acc1) != 0 {
		t.Fatalf("multi-round mismatch:\n  nbrArrayHasher = %s\n  manual chained = %s",
			hProd.Text(16), acc1.Text(16))
	}
}

// TestNbrArrayHasherPanicsOnBadLen makes sure we catch invalid lengths.
func TestNbrArrayHasherPanicsOnBadLen(t *testing.T) {
	assertPanics := func(name string, f func()) {
		t.Helper()
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s: expected panic, got none", name)
			}
		}()
		f()
	}

	// Too short: len < 16.
	assertPanics("short", func() {
		nbrArrayHasher(make([]uint64, 10))
	})

	// Wrong shape: (len-1) not divisible by 15.
	assertPanics("bad-mod", func() {
		// padLen = 17 => (17-1)=16, 16%15 != 0
		nbrArrayHasher(make([]uint64, 17))
	})

	// Sanity: a valid length should NOT panic.
	padLen := padLenFromMaxDegree(15) // 16
	nbrData := make([]uint64, padLen)
	nbrData[0] = 0
	_ = nbrArrayHasher(nbrData) // should not panic
}
