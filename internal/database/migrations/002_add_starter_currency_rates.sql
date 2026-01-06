-- +goose Up
INSERT INTO currency_exchange (conversion_direction, exchange_rate, last_updated)
VALUES ("USDollarCANDollar", 1.2, DATETIME()),
	("USDollarEuro", 0.7, DATETIME()),
	("USDollarPeso", 34.2, DATETIME()),
	("CANDollarUSDollar", 0.8, DATETIME()),
	("CANDollarEuro", 0.64, DATETIME()),
	("CANDollarPeso", 23.6, DATETIME()),
	("EuroUSDollar", 1.4, DATETIME()),
	("EuroCANDollar", 1.63, DATETIME()),
	("EuroPeso", 76.1, DATETIME()),
	("PesoUSDollar", 0.38, DATETIME()),
	("PesoCANDollar", 0.51, DATETIME()),
	("PesoEuro", 0.23, DATETIME());

-- +goose Down
DELETE FROM currency_exchange;	
