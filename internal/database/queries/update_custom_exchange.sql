-- name: UpdateCustomExchange :exec
UPDATE user_conversions
SET exchange_rate = ?
WHERE start_type = ? AND end_type = ?;
