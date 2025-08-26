-- +goose Up
-- +goose StatementBegin
-- Table IDs should be UUIDs
CREATE EXTENSION "pgcrypto";

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    status TEXT CHECK (status IN ('BACKLOG', 'TODO', 'IN_PROGRESS', 'DONE'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION "pgcrypto";

DROP TABLE tasks;
-- +goose StatementEnd
