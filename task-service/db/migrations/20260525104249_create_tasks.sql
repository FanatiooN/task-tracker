-- +goose Up
CREATE TYPE task_status AS ENUM ('in_progress', 'done', 'cancelled');

CREATE TABLE tasks (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   user_id UUID NOT NULL,
   title TEXT NOT NULL,
   description TEXT,
   status task_status NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
   deleted_at TIMESTAMPTZ
);
-- +goose Down
DROP TABLE tasks;
DROP TYPE task_status;