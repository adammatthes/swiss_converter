-- +goose Up
INSERT INTO currency_exchange (conversion_direction, exchange_rate, last_updated)
VALUES ("US DollarCanadian Dollar", 1.39, DATETIME()),
	("US DollarEuro", 0.86, DATETIME()),
	("US DollarPeso", 17.58, DATETIME()),
	("Canadian DollarUS Dollar", 0.72, DATETIME()),
	("Canadian DollarEuro", 0.62, DATETIME()),
	("Canadian DollarPeso", 12.68, DATETIME()),
	("EuroUS Dollar", 1.16, DATETIME()),
	("EuroCanadian Dollar", 1.61, DATETIME()),
	("EuroPeso", 20.47, DATETIME()),
	("PesoUS Dollar", 0.057, DATETIME()),
	("PesoCanadian Dollar", 0.079, DATETIME()),
	("PesoEuro", 0.049, DATETIME());

-- +goose Down
DELETE FROM currency_exchange;	
