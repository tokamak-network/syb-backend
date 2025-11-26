package protocol

import (
	"encoding/binary"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// helper to encode edges into the same packed format as the contract:
// N * 8 bytes: [ilo(4) || ihi(4)] big-endian.
func encodeEdgesPacked(edges []Edge) []byte {
	out := make([]byte, 0, len(edges)*8)
	tmp := make([]byte, 8)
	for _, e := range edges {
		binary.BigEndian.PutUint32(tmp[0:4], uint32(e.Ilo))
		binary.BigEndian.PutUint32(tmp[4:8], uint32(e.Ihi))
		out = append(out, tmp...)
	}
	return out
}

// helper to build a topic for uint64 batchId (uint256 big-endian).
func encodeUint64Topic(x uint64) common.Hash {
	var buf [32]byte
	binary.BigEndian.PutUint64(buf[24:], x) // last 8 bytes
	return common.BytesToHash(buf[:])
}

func TestDecodeEdgesPacked_Basic(t *testing.T) {
	edgesIn := []Edge{
		{Ilo: 1, Ihi: 2},
		{Ilo: 7, Ihi: 42},
		{Ilo: 100, Ihi: 200},
	}

	packed := encodeEdgesPacked(edgesIn)

	edgesOut, err := decodeEdgesPacked(packed)
	if err != nil {
		t.Fatalf("decodeEdgesPacked error: %v", err)
	}

	if !reflect.DeepEqual(edgesOut, edgesIn) {
		t.Fatalf("decodeEdgesPacked round-trip mismatch:\n got  %#v\n want %#v",
			edgesOut, edgesIn)
	}
}

func TestDecodeBatchFromLog_Basic(t *testing.T) {
	// Construct a synthetic BatchSubmitted log:
	//
	//   topics[0] = event ID
	//   topics[1] = batchId (indexed uint64)
	//   data      = abi-encoded (count, storageHash, newGraphRoot, newScoreRoot, edgesPacked)

	batchID := uint64(123)
	edges := []Edge{
		{Ilo: 1, Ihi: 2},
		{Ilo: 2, Ihi: 3},
	}
	edgesPacked := encodeEdgesPacked(edges)
	count := uint32(len(edges))

	var storageHash [32]byte
	var newGraphRoot [32]byte
	var newScoreRoot [32]byte

	// Give them some non-zero values so we can assert they come back intact.
	copy(storageHash[:], []byte("storage-hash-1234567890abcdef"))
	copy(newGraphRoot[:], []byte("graph-root-1234567890abcdef"))
	copy(newScoreRoot[:], []byte("score-root-1234567890abcdef"))

	// Use the same batchABI we defined in batch_decode.go to pack the data.
	data, err := batchABI.Pack(
		"BatchSubmitted",
		count,
		storageHash,
		newGraphRoot,
		newScoreRoot,
		edgesPacked,
	)
	if err != nil {
		t.Fatalf("batchABI.Pack: %v", err)
	}

	log := types.Log{
		Topics: []common.Hash{
			BatchSubmittedEventID(),  // topic[0]: event signature
			encodeUint64Topic(batchID), // topic[1]: indexed batchId
		},
		Data: data,
		// BlockNumber, Index, etc. can stay zero for this unit test.
	}

	decoded, err := DecodeBatchFromLog(log)
	if err != nil {
		t.Fatalf("DecodeBatchFromLog error: %v", err)
	}

	// Check basic fields.
	if decoded.BatchID != batchID {
		t.Fatalf("BatchID mismatch: got %d, want %d", decoded.BatchID, batchID)
	}
	if decoded.Count != count {
		t.Fatalf("Count mismatch: got %d, want %d", decoded.Count, count)
	}

	if decoded.StorageHash != storageHash {
		t.Fatalf("StorageHash mismatch:\n got  %x\n want %x",
			decoded.StorageHash, storageHash)
	}
	if decoded.NewGraphRoot != newGraphRoot {
		t.Fatalf("NewGraphRoot mismatch:\n got  %x\n want %x",
			decoded.NewGraphRoot, newGraphRoot)
	}
	if decoded.NewScoreRoot != newScoreRoot {
		t.Fatalf("NewScoreRoot mismatch:\n got  %x\n want %x",
			decoded.NewScoreRoot, newScoreRoot)
	}

	// Check edges.
	if !reflect.DeepEqual(decoded.Edges, edges) {
		t.Fatalf("Edges mismatch:\n got  %#v\n want %#v",
			decoded.Edges, edges)
	}

	// Small sanity check: edgesPacked length vs Count.
	if len(edgesPacked) != int(decoded.Count)*8 {
		t.Fatalf("edgesPacked length mismatch: got %d, want %d",
			len(edgesPacked), int(decoded.Count)*8)
	}

	// And that BatchSubmittedEventID is consistent with the ABI.
	if BatchSubmittedEventID() != batchSubmittedEvent.ID {
		t.Fatalf("BatchSubmittedEventID != batchSubmittedEvent.ID")
	}

	// Bonus: confirm that topic[1] decodes to batchID the same way DecodeBatchFromLog does.
	bidBig := new(big.Int).SetBytes(log.Topics[1].Bytes())
	if !bidBig.IsUint64() || bidBig.Uint64() != batchID {
		t.Fatalf("topic[1] roundtrip mismatch: got %s, want %d",
			bidBig.String(), batchID)
	}
}
