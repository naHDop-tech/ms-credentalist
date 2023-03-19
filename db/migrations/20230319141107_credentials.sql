-- +goose Up
-- +goose StatementBegin
CREATE TABLE "credentials" (
    "id" uuid PRIMARY KEY,
    "title" varchar(30) NOT NULL,
    "login_name" varchar(255) NOT NULL,
    "secret" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL,
    "customer_id" uuid UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);
ALTER TABLE "credentials" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "credentials" DROP CONSTRAINT IF EXISTS "customer_id";
DROP TABLE IF EXISTS "credentials";
-- +goose StatementEnd
