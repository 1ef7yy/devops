-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
