-- +goose Up
INSERT INTO currency_exchange (conversion_direction, exchange_rate, last_updated)
VALUES ("usd-cad", 1.39, DATETIME()),
	("usd-eur", 0.86, DATETIME()),
	("usd-mxn", 17.58, DATETIME()),
	("cad-usd", 0.72, DATETIME()),
	("cad-eur", 0.62, DATETIME()),
	("cad-mxn", 12.68, DATETIME()),
	("eur-usd", 1.16, DATETIME()),
	("eur-cad", 1.61, DATETIME()),
	("eur-mxn", 20.47, DATETIME()),
	("mxn-usd", 0.057, DATETIME()),
	("mxn-cad", 0.079, DATETIME()),
	("mxn-eur", 0.049, DATETIME());

-- +goose Down
DELETE FROM currency_exchange;	
