-- +goose Up
INSERT INTO currency_exchange (conversion_direction, exchange_rate, last_updated)
VALUES ("US DollarCanadian Dollar", 1.2, DATETIME()),
	("US DollarEuro", 0.7, DATETIME()),
	("US DollarPeso", 34.2, DATETIME()),
	("Canadian DollarUS Dollar", 0.8, DATETIME()),
	("Canadian DollarEuro", 0.64, DATETIME()),
	("Canadian DollarPeso", 23.6, DATETIME()),
	("Euro US Dollar", 1.4, DATETIME()),
	("EuroCanadian Dollar", 1.63, DATETIME()),
	("EuroPeso", 76.1, DATETIME()),
	("PesoUS Dollar", 0.38, DATETIME()),
	("PesoCanadian Dollar", 0.51, DATETIME()),
	("PesoEuro", 0.23, DATETIME());

-- +goose Down
DELETE FROM currency_exchange;	
