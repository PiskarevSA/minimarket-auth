-- name: CreateAccount :exec
INSERT INTO accounts (
    id, 
    login, 
    password_hash,
    created_at,
    updated_at
)
VALUES (
    @id, 
    @login, 
    @password_hash,
    @created_at,
    @created_at
);