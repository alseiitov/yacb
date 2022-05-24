-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_subscriptions
(
    id                     SERIAL      NOT NULL UNIQUE PRIMARY KEY,
    user_id                INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    crypto_currency_symbol VARCHAR(32) NOT NULL,
    update_interval        INTERVAL    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_subscriptions;
-- +goose StatementEnd
