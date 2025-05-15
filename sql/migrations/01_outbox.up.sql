CREATE TABLE IF NOT EXISTS outbox (
    id BIGSERIAL NOT NULL,
    event VARCHAR(64) NOT NULL,
    status VARCHAR(16) NOT NULL DEFAULT 'UNPROCESSED',
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by VARCHAR(32) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by VARCHAR(32) NOT NULL,
    
    PRIMARY KEY (id, created_at)
);

CREATE INDEX IF NOT EXISTS outbox_created_at_not_completed_idx
    ON outbox (created_at, status) 
    WHERE (status != 'COMPLETED');

SELECT create_hypertable(
    'outbox',
    'created_at',
    chunk_time_interval => INTERVAL '3 day',
    if_not_exists => TRUE
);

SELECT add_retention_policy(
    'outbox',
    INTERVAL '6 day',
  	if_not_exists => TRUE
);

