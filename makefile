DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}
migrate:
	$(MIGRATE) up
migrate-down:
	$(MIGRATE) down
run:
	go run cmd/app/main.go
migrate-force:
	migrate -database $(DB_DSN) -path migrations force 20250216082214

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go