-- +goose Up
CREATE TABLE user_conversions (
	conversion_id INT PRIMARY KEY,
	start_type TEXT NOT NULL,
	end_type TEXT NOT NULL,
	exchange_rate DECIMAL NOT NULL,
	last_updated DATE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE user_conversion;
