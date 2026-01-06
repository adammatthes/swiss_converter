-- name: GetExchangeRate :one
SELECT exchange_rate FROM currency_exchange
	WHERE conversion_direction = ?;
