-- +goose Up
ALTER TABLE refresh_tokens ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE refresh_tokens ALTER COLUMN id DROP DEFAULT;