include .env
export

start_prod:
	@./bin

start-server:
	@cd cmd && go run main.go

migrate-up:
	@migrate -path internal/database/migrate -database ${CONN_STRING} up

migrate-down:
	@migrate -path internal/database/migrate -database ${CONN_STRING} down