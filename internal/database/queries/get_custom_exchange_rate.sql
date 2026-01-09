-- name: GetCustomExchangeRate :one
SELECT exchange_rate FROM user_conversions
	WHERE start_type = ? AND end_type = ?;
