-- +goose Up
ALTER TABLE tasks ALTER COLUMN status SET DEFAULT 'in_progress';

-- +goose Down
ALTER TABLE tasks ALTER COLUMN status DROP DEFAULT;
