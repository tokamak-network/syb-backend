package protocol

import (
	"encoding/binary"
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

	// IMPORTANT: for events we don't use abi.Pack("BatchSubmitted", ...).
	// We pack only the NON-INDEXED arguments, in order, using the event's
	// Inputs.NonIndexed() helper.
	nonIndexed := batchSubmittedEvent.Inputs.NonIndexed()
	data, err := nonIndexed.Pack(
		count,
		storageHash,
		newGraphRoot,
		newScoreRoot,
		edgesPacked,
	)
	if err != nil {
		t.Fatalf("NonIndexed.Pack: %v", err)
	}

	log := types.Log{
		Topics: []common.Hash{
			BatchSubmittedEventID(),      // topic[0]: event signature
			encodeUint64Topic(batchID),   // topic[1]: indexed batchId
		},
		Data: data,
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

	// Sanity checks.
	if len(edgesPacked) != int(decoded.Count)*8 {
		t.Fatalf("edgesPacked length mismatch: got %d, want %d",
			len(edgesPacked), int(decoded.Count)*8)
	}

	if BatchSubmittedEventID() != batchSubmittedEvent.ID {
		t.Fatalf("BatchSubmittedEventID != batchSubmittedEvent.ID")
	}
}
