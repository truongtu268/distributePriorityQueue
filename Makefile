start_database:
	docker compose up

stop_database:
	docker compose down

migrate_up:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/advertisement?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/advertisement?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: start_database stop_database migrate_up migrate_down sqlc