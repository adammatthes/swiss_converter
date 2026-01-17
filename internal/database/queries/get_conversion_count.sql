-- name: GetConversionCount :one
SELECT COUNT(*) FROM query_log
	WHERE start_type = ? AND end_type = ?;
