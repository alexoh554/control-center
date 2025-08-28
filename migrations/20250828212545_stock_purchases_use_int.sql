-- +goose Up
-- +goose StatementBegin
ALTER TABLE stock_purchases
DROP COLUMN price_cents;

ALTER TABLE stock_purchases
DROP COLUMN total_price_cents;

ALTER TABLE stock_purchases
ADD COLUMN price_cents INT NOT NULL;

ALTER TABLE stock_purchases
ADD COLUMN total_price_cents INT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stock_purchases
DROP COLUMN price_cents;

ALTER TABLE stock_purchases
DROP COLUMN total_price_cents;

ALTER TABLE stock_purchases
ADD COLUMN price_cents NUMERIC(12, 2) NOT NULL NOT NULL;

ALTER TABLE stock_purchases
ADD COLUMN total_price_cents NUMERIC(12, 2) NOT NULL NOT NULL;
-- +goose StatementEnd
