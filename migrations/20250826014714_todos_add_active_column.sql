-- +goose Up
-- +goose StatementBegin

ALTER TABLE todos 
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE todos
ADD COLUMN deleted_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE todos
DROP COLUMN created_at;

ALTER TABLE todos
DROP COLUMN deleted_at

-- +goose StatementEnd
