package protocol

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Edge is a decoded edge from edgesPacked (ilo, ihi) where ilo < ihi.
type Edge struct {
	Ilo uint64
	Ihi uint64
}

// Batch is the decoded, protocol-level representation of a BatchSubmitted event.
type Batch struct {
	BatchID      uint64
	Count        uint32
	StorageHash  [32]byte
	NewGraphRoot [32]byte
	NewScoreRoot [32]byte
	Edges        []Edge
}

// -----------------------------------------------------------------------------
// ABI setup for BatchSubmitted
// -----------------------------------------------------------------------------

// This JSON is just the event ABI for BatchSubmitted from your contract.
const batchSubmittedEventABIJSON = `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint64",
        "name": "batchId",
        "type": "uint64"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "count",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "storageHash",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "newGraphRoot",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "bytes32",
        "name": "newScoreRoot",
        "type": "bytes32"
      },
      {
        "indexed": false,
        "internalType": "bytes",
        "name": "edgesPacked",
        "type": "bytes"
      }
    ],
    "name": "BatchSubmitted",
    "type": "event"
  }
]`

var (
	batchABI           abi.ABI
	batchSubmittedEvent abi.Event
)

func init() {
	a, err := abi.JSON(strings.NewReader(batchSubmittedEventABIJSON))
	if err != nil {
		panic(fmt.Errorf("protocol: failed to parse BatchSubmitted ABI: %w", err))
	}
	batchABI = a

	ev, ok := batchABI.Events["BatchSubmitted"]
	if !ok {
		panic("protocol: BatchSubmitted event not found in ABI")
	}
	batchSubmittedEvent = ev
}

// BatchSubmittedEventID returns the keccak256 topic for the BatchSubmitted event.
// You can use this in syncer.Config.BatchSubmittedTopic instead of hardcoding.
func BatchSubmittedEventID() common.Hash {
	return batchSubmittedEvent.ID
}

// -----------------------------------------------------------------------------
// Decoding
// -----------------------------------------------------------------------------

// DecodeBatchFromLog decodes a types.Log (which must be a BatchSubmitted event)
// into a Batch struct.
//
// It assumes:
//   - topic[0] == BatchSubmittedEventID()
//   - topic[1] is the indexed batchId (uint64)
//   - data encodes (count, storageHash, newGraphRoot, newScoreRoot, edgesPacked)
//     using standard Solidity ABI rules.
func DecodeBatchFromLog(lg types.Log) (*Batch, error) {
	if len(lg.Topics) < 2 {
		return nil, fmt.Errorf("protocol: DecodeBatchFromLog: expected at least 2 topics, got %d", len(lg.Topics))
	}
	if lg.Topics[0] != batchSubmittedEvent.ID {
		return nil, fmt.Errorf("protocol: DecodeBatchFromLog: topic[0] mismatch (not BatchSubmitted)")
	}

	// Decode indexed batchId from topic[1] (uint64 stored as 32-byte big-endian).
	bidBig := new(big.Int).SetBytes(lg.Topics[1].Bytes())
	if !bidBig.IsUint64() {
		return nil, fmt.Errorf("protocol: batchId does not fit in uint64")
	}
	batchID := bidBig.Uint64()

	// Decode non-indexed fields from lg.Data.
	// We define a mirror struct for the event's non-indexed parameters.
	var decoded struct {
		Count        uint32
		StorageHash  [32]byte
		NewGraphRoot [32]byte
		NewScoreRoot [32]byte
		EdgesPacked  []byte
	}

	if err := batchABI.UnpackIntoInterface(&decoded, "BatchSubmitted", lg.Data); err != nil {
		return nil, fmt.Errorf("protocol: abi unpack BatchSubmitted: %w", err)
	}

	// Sanity: edgesPacked length should be Count * 8 bytes.
	if len(decoded.EdgesPacked)%8 != 0 {
		return nil, fmt.Errorf("protocol: edgesPacked length %d not multiple of 8", len(decoded.EdgesPacked))
	}
	nEdges := len(decoded.EdgesPacked) / 8
	if uint32(nEdges) != decoded.Count {
		return nil, fmt.Errorf("protocol: edgesPacked count %d != declared count %d", nEdges, decoded.Count)
	}

	edges, err := decodeEdgesPacked(decoded.EdgesPacked)
	if err != nil {
		return nil, err
	}

	return &Batch{
		BatchID:      batchID,
		Count:        decoded.Count,
		StorageHash:  decoded.StorageHash,
		NewGraphRoot: decoded.NewGraphRoot,
		NewScoreRoot: decoded.NewScoreRoot,
		Edges:        edges,
	}, nil
}

// decodeEdgesPacked interprets edgesPacked as N*8 bytes: [ilo(4)||ihi(4)] big-endian
// per edge, exactly matching the Solidity packing in _buildEdgesAndStorageHash.
func decodeEdgesPacked(b []byte) ([]Edge, error) {
	if len(b)%8 != 0 {
		return nil, fmt.Errorf("protocol: decodeEdgesPacked: length %d not multiple of 8", len(b))
	}
	n := len(b) / 8
	edges := make([]Edge, 0, n)

	for i := 0; i < n; i++ {
		off := i * 8
		ilo := binary.BigEndian.Uint32(b[off : off+4])
		ihi := binary.BigEndian.Uint32(b[off+4 : off+8])

		if ilo == 0 || ihi == 0 {
			// Not strictly required, but if you treat 0 as "invalid" idx you can
			// enforce that here. Otherwise, remove this check.
			// We'll be strict for now but easy to relax later.
			// return nil, fmt.Errorf("protocol: decodeEdgesPacked: got zero index at edge %d", i)
		}

		edges = append(edges, Edge{
			Ilo: uint64(ilo),
			Ihi: uint64(ihi),
		})
	}

	return edges, nil
}
