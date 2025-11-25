-- Edges table for SQLStore. Undirected edges are stored canonically (ilo < ihi).

CREATE TABLE IF NOT EXISTS edges (
    ilo BIGINT NOT NULL,
    ihi BIGINT NOT NULL,
    PRIMARY KEY (ilo, ihi)
);

CREATE INDEX IF NOT EXISTS idx_edges_ilo ON edges (ilo);
CREATE INDEX IF NOT EXISTS idx_edges_ihi ON edges (ihi);
