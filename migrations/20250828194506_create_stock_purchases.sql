-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock_purchases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- ID of the purchase
    symbol TEXT NOT NULL,
    price_cents NUMERIC(12, 2) NOT NULL,
    quantity INT NOT NULL,
    purchased_at TIMESTAMPTZ NOT NULL,
    total_price_cents NUMERIC(12, 2) NOT NULL, -- price_cents * quantity
    broker TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock_purchases;
-- +goose StatementEnd
