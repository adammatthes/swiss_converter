-- name: DeleteCustomConversion :exec
DELETE FROM user_conversions
WHERE (start_type = :start_type AND end_type = :end_type)
	OR (start_type = :end_type AND end_type = :start_type)
RETURNING *;
