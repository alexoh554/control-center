-- +goose Up
-- +goose StatementBegin

ALTER TABLE tasks 
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE tasks
ADD COLUMN deleted_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE tasks
DROP COLUMN created_at;

ALTER TABLE tasks
DROP COLUMN deleted_at

-- +goose StatementEnd
