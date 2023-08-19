migrateup:
	migrate -path pkg/lib/database/migration -database "postgresql://@192.168.1.100:5432/alpha-db?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/lib/database/migration -database "postgresql://@192.168.1.100:5432/alpha-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate