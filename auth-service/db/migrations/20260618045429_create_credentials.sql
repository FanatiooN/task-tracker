-- +goose Up
CREATE TABLE credentials (
    user_id UUID PRIMARY KEY NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

-- +goose Down
DROP TABLE credentials
