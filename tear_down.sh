cd internal/database/migrations

for i in {1..4}; do
	goose sqlite3 ../../../swiss.db down
	if [ $? -ne 0 ]; then
		break
	fi
done
