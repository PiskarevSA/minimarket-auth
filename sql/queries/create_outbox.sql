-- name: CreateOutbox :exec
INSERT INTO outbox (
    eventName,
    status,
    payload,
    created_at,
    created_by,
    updated_at,
    updated_by
)
SELECT
    @eventName,
    @status,
    @payload,
    @created_at,
    @created_by,
    @updated_at,
    @updated_by;
