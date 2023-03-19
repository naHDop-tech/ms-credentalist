-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customers" (
    "id" uuid PRIMARY KEY,
    "password" varchar(255) NOT NULL,
    "user_name" varchar(200) UNIQUE NOT NULL,
    "user_id" uuid UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);
ALTER TABLE "customers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "customers" DROP CONSTRAINT IF EXISTS "user_id";
DROP TABLE IF EXISTS "customers";
-- +goose StatementEnd
