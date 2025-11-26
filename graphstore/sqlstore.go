package graphstore

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
)

// SQLStore is a GraphStore backed by a simple edges(ilo, ihi) table.
// It implicitly satisfies protocol.GraphStore:
//
//   type GraphStore interface {
//       AddEdge(ctx context.Context, u, v uint64) error
//       Neighbors(ctx context.Context, v uint64) ([]uint64, error)
//   }
//
// No import of protocol is needed here – method signatures match.
type SQLStore struct {
	DB *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{DB: db}
}

// AddEdge inserts an undirected edge {u,v}, stored canonically (ilo < ihi).
// It is idempotent thanks to the PRIMARY KEY + ON CONFLICT DO NOTHING.
func (s *SQLStore) AddEdge(ctx context.Context, u, v uint64) error {
	if u == v {
		// You can decide whether to treat this as error or no-op.
		return fmt.Errorf("self edge not allowed: %d", u)
	}
	if u > v {
		u, v = v, u
	}
	_, err := s.DB.ExecContext(ctx,
		`INSERT INTO edges (ilo, ihi) VALUES ($1, $2)
         ON CONFLICT (ilo, ihi) DO NOTHING`,
		u, v,
	)
	return err
}

// Neighbors returns the sorted list of neighbors of v.
func (s *SQLStore) Neighbors(ctx context.Context, v uint64) ([]uint64, error) {
	rows, err := s.DB.QueryContext(ctx, `
        SELECT
          CASE WHEN ilo = $1 THEN ihi ELSE ilo END AS nbr
        FROM edges
        WHERE ilo = $1 OR ihi = $1
        ORDER BY nbr`,
		v,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []uint64
	for rows.Next() {
		var nbr uint64
		if err := rows.Scan(&nbr); err != nil {
			return nil, err
		}
		out = append(out, nbr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Should already be sorted and unique due to PK + ORDER BY,
	// but we defensively sort anyway.
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })

	return out, nil
}
