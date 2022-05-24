-- +goose Up
-- +goose StatementBegin
INSERT INTO crypto_currencies (symbol, name)
VALUES ('BTC', 'Bitcoin'),
       ('ETH', 'Ethereum'),
       ('SOL', 'Solana'),
       ('ADA', 'Cardano');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM crypto_currencies;
-- +goose StatementEnd
