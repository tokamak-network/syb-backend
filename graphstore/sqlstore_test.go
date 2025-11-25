package graphstore

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	// apply schema
	_, err = db.Exec(`
        CREATE TABLE edges (
            ilo BIGINT NOT NULL,
            ihi BIGINT NOT NULL,
            PRIMARY KEY (ilo, ihi)
        );
    `)
	if err != nil {
		t.Fatalf("creating schema: %v", err)
	}
	return db
}

func TestSQLStore_AddEdgeAndNeighbors(t *testing.T) {
    ctx := context.Background()
    db := newTestDB(t)
    defer db.Close()

    s := NewSQLStore(db)

    // same triangle test as MemoryStore
    if err := s.AddEdge(ctx, 1, 2); err != nil {
        t.Fatalf("AddEdge(1,2): %v", err)
    }
    if err := s.AddEdge(ctx, 2, 3); err != nil {
        t.Fatalf("AddEdge(2,3): %v", err)
    }
    if err := s.AddEdge(ctx, 3, 1); err != nil {
        t.Fatalf("AddEdge(3,1): %v", err)
    }
    if err := s.AddEdge(ctx, 2, 1); err != nil { // dup
        t.Fatalf("AddEdge(2,1) dup: %v", err)
    }

    // self-edge MUST error
    if err := s.AddEdge(ctx, 1, 1); err == nil {
        t.Fatalf("expected error on self-edge, got nil")
        // optionally:
        // if !strings.Contains(err.Error(), "self edge not allowed") { ... }
    }

    cases := []struct {
        v    uint64
        want []uint64
    }{
        {1, []uint64{2, 3}},
        {2, []uint64{1, 3}},
        {3, []uint64{1, 2}},
        {4, nil},
    }

    for _, tc := range cases {
        got, err := s.Neighbors(ctx, tc.v)
        if err != nil {
            t.Fatalf("Neighbors(%d) err: %v", tc.v, err)
        }
        if !reflect.DeepEqual(got, tc.want) {
            t.Fatalf("Neighbors(%d) = %v, want %v", tc.v, got, tc.want)
        }
    }
}

}
