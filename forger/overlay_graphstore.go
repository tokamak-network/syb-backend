type OverlayGraphStore struct {
	base GraphStore

	mu  sync.RWMutex
	adj map[uint64]map[uint64]struct{} // overlay edges only
}

func NewOverlayGraphStore(base GraphStore) *OverlayGraphStore {
	return &OverlayGraphStore{
		base: base,
		adj:  make(map[uint64]map[uint64]struct{}),
	}
}

func (o *OverlayGraphStore) AddEdge(ctx context.Context, u, v uint64) error {
	if u == v {
		return nil
	}
	if u > v {
		u, v = v, u
	}

	o.mu.Lock()
	defer o.mu.Unlock()

	if o.adj[u] == nil {
		o.adj[u] = make(map[uint64]struct{})
	}
	if o.adj[v] == nil {
		o.adj[v] = make(map[uint64]struct{})
	}
	o.adj[u][v] = struct{}{}
	o.adj[v][u] = struct{}{}
	return nil
}

func (o *OverlayGraphStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	baseNbrs, err := o.base.Neighbors(ctx, v)
	if err != nil {
		return nil, err
	}

	o.mu.RLock()
	defer o.mu.RUnlock()

	m := make(map[uint64]struct{}, len(baseNbrs))
	for _, n := range baseNbrs {
		m[n] = struct{}{}
	}
	if overlaySet, ok := o.adj[v]; ok {
		for n := range overlaySet {
			m[n] = struct{}{}
		}
	}

	out := make([]uint64, 0, len(m))
	for n := range m {
		out = append(out, n)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out, nil
}

