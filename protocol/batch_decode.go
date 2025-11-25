package protocol

// Edge is a decoded edge from edgesPacked (ilo, ihi).
type Edge struct {
	Ilo uint64
	Ihi uint64
}

// Batch is a decoded, protocol-level representation of a BatchSubmitted event.
//
// For now we only need BatchID, Count and Edges to test ApplyBatch / ApplyEdge.
// You can extend this later with StorageHash, NewGraphRoot, NewScoreRoot, etc.
type Batch struct {
	BatchID uint64
	Count   uint32
	Edges   []Edge
}
