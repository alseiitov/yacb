FROM golang:1.17 as build

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./service_crypto_currency/build/app ./service_crypto_currency/cmd/main.go


FROM alpine:3.16

WORKDIR /app
COPY --from=build /app .

CMD ["./service_crypto_currency/build/app"]