-- name: AddCustomConversion :exec
INSERT INTO user_conversions(start_type, end_type, exchange_rate)
	VALUES (?, ?, ?);
