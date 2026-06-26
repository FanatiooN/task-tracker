-- +goose Up
CREATE TABLE oauth_credentials (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    provider TEXT NOT NULL,
    provider_id TEXT NOT NULL,
    UNIQUE(provider, provider_id)
);

-- +goose Down
DROP TABLE oauth_credentials;