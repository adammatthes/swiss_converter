-- name: AddQueryLogEntry :exec
INSERT INTO query_log (start_type, end_type, amount, dtg)
	VALUES (?, ?, ?, CURRENT_TIMESTAMP)
RETURNING *;
