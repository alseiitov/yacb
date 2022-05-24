-- +goose Up
-- +goose StatementBegin
CREATE TABLE crypto_currencies
(
    id     SERIAL      NOT NULL UNIQUE PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL UNIQUE,
    name   VARCHAR(64) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS crypto_currencies;
-- +goose StatementEnd
