-- +goose Up
CREATE TABLE refresh_tokens(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL ,
    is_revoked BOOL DEFAULT FALSE
);

-- +goose Down
DROP TABLE refresh_tokens;
