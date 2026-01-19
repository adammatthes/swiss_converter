cd internal/database/migrations

goose sqlite3 ../../../swiss.db up

cd ../../..
sqlc generate
