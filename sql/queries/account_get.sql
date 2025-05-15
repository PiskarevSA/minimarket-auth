-- name: GetAccountCredentials :one
SELECT
    id AS user_id,
    password_hash
FROM accounts
WHERE login = @login;
