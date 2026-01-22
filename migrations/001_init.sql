CREATE TABLE blocks (
    number        BIGINT PRIMARY KEY,
    hash          TEXT NOT NULL,
    parent_hash   TEXT NOT NULL,
    indexed_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE events (
    chain_id        BIGINT NOT NULL,
    block_number    BIGINT NOT NULL,
    block_hash      TEXT NOT NULL,
    tx_hash         TEXT NOT NULL,
    log_index       INTEGER NOT NULL,
    contract        TEXT NOT NULL,
    event_name      TEXT NOT NULL,
    payload         JSONB NOT NULL,

    PRIMARY KEY (chain_id, tx_hash, log_index)
);

CREATE TABLE indexer_state (
    id              INTEGER PRIMARY KEY CHECK (id = 1),
    last_indexed    BIGINT NOT NULL
);

INSERT INTO indexer_state (id, last_indexed) VALUES (1, 0);
