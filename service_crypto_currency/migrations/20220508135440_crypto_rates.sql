-- +goose Up
-- +goose StatementBegin
CREATE TABLE crypto_rates
(
    crypto_currency_id INTEGER          NOT NULL REFERENCES crypto_currencies (id) ON DELETE CASCADE,
    price              DOUBLE PRECISION NOT NULL,
    date               TIMESTAMP        NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS crypto_rates;
-- +goose StatementEnd
