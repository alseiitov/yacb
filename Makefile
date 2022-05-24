include .env

.PHONY: build migrate_up migrate_down protoc install-goose
.SILENT: migrate-up migrate-down

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./service_crypto_currency/build/app ./service_crypto_currency/cmd/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./service_telegram_bot/build/app ./service_telegram_bot/cmd/main.go

migrate-up:
	goose --dir service_crypto_currency/migrations postgres "user=${POSTGRES_USER_CC} password=${POSTGRES_PASSWORD_CC} dbname=${POSTGRES_DB_CC} port=5432 sslmode=disable" up
	goose --dir service_telegram_bot/migrations postgres "user=${POSTGRES_USER_TB} password=${POSTGRES_PASSWORD_TB} dbname=${POSTGRES_DB_TB} port=5433 sslmode=disable" up

migrate-down:
	goose --dir service_crypto_currency/migrations postgres "user=${POSTGRES_USER_CC} password=${POSTGRES_PASSWORD_CC} dbname=${POSTGRES_DB_CC} port=5432 sslmode=disable" down
	goose --dir service_telegram_bot/migrations postgres "user=${POSTGRES_USER_TB} password=${POSTGRES_PASSWORD_TB} dbname=${POSTGRES_DB_TB} port=5433 sslmode=disable" down

protoc:
	protoc -I ./service_crypto_currency/proto \
		--go_out=./service_crypto_currency/proto  \
		--go-grpc_out=./service_crypto_currency/proto \
		--grpc-gateway_out=./service_crypto_currency/proto \
		--openapiv2_out=./service_crypto_currency/docs \
		./service_crypto_currency/proto/*.proto

	protoc -I ./service_telegram_bot/proto \
		--go_out=./service_telegram_bot/proto  \
		--go-grpc_out=./service_telegram_bot/proto \
		--grpc-gateway_out=./service_telegram_bot/proto \
		--openapiv2_out=./service_telegram_bot/docs \
		./service_telegram_bot/proto/*.proto

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest
