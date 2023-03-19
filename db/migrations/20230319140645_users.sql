-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
    "id" uuid PRIMARY KEY,
    "email" varchar(30) UNIQUE NOT NULL,
    "user_name" varchar(50) UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
