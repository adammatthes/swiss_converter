-- +goose Up
CREATE TABLE query_log (
	id INTEGER PRIMARY KEY,
	start_type TEXT NOT NULL,
	end_type TEXT NOT NULL,
	amount TEXT NOT NULL,
	dtg DATE DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE query_log;
