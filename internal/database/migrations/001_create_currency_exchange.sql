-- +goose Up
CREATE TABLE currency_exchange (
	conversion_direction TEXT PRIMARY KEY,
	exchange_rate DECIMAL NOT NULL,
	last_updated DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE currency_exchange;
