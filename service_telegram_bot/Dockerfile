FROM golang:1.17 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./service_telegram_bot/build/app ./service_telegram_bot/cmd/main.go


FROM alpine:3.16

WORKDIR /app
COPY --from=build /app .

CMD ["./service_telegram_bot/build/app"]