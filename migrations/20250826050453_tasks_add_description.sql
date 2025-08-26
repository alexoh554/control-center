-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks
ADD COLUMN description TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tasks
DROP COLUMN description;
-- +goose StatementEnd
