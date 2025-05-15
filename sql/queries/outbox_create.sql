-- name: CreateOutbox :exec
INSERT INTO outbox (
    event,
    payload,
    created_at,
    created_by,
    updated_at,
    updated_by
)
SELECT
    @event,
    @payload,
    @created_at,
    @created_by,
    @updated_at,
    @updated_by;
