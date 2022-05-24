-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id       INTEGER      NOT NULL NOT NULL UNIQUE PRIMARY KEY,
    password VARCHAR(250) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
