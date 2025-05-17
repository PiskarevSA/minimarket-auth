CREATE TABLE IF NOT EXISTS outbox (
    id BIGSERIAL NOT NULL,
    eventName VARCHAR(64) NOT NULL,
    status VARCHAR(16) NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    updated_by VARCHAR(32) NOT NULL,
    
    PRIMARY KEY (id, created_at)
);

CREATE INDEX IF NOT EXISTS outbox_created_at_not_completed_idx
    ON outbox (created_at, status) 
    WHERE (status != 'PROCESSED');

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

