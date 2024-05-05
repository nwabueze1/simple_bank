migratedown:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations down

migratedown1:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations down 1

migrateup:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations up

migrateup1:
	migrate -database postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable -path db/migrations up 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .

.PHONY: migratedown migrateup migratedown1 migrateup1 sqlc test server