-- name: SetCurrencyRate :exec
UPDATE currency_exchange
SET exchange_rate = ?, last_updated = CURRENT_TIMESTAMP
WHERE conversion_direction = ?

RETURNING *;
