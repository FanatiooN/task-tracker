-- +goose Up
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    type TEXT NOT NULL,
    provider TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
DROP TABLE notifications;