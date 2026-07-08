-- +goose Up
CREATE TABLE user_contacts(
    user_id UUID NOT NULL,
    provider TEXT NOT NULL,
    contact TEXT NOT NULL,
    PRIMARY KEY(user_id, provider)
);

-- +goose Down
DROP TABLE user_contacts;