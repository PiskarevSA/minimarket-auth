CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY,
    login VARCHAR(32) UNIQUE NOT NULL,
    password_hash CHAR(60) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);