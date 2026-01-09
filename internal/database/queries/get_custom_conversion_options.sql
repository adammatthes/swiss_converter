-- name: GetCustomConversionOptions :many
SELECT end_type FROM user_conversions
	WHERE start_type = ?;
