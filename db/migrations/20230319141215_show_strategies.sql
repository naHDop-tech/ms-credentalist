-- +goose Up
-- +goose StatementBegin
CREATE TABLE "show_strategies" (
    "id" uuid PRIMARY KEY,
    "show_immediately" boolean NOT NULL,
    "send_to_email" boolean NOT NULL,
    "send_to_phone" boolean NOT NULL,
    "credential_id" uuid UNIQUE NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);
ALTER TABLE "show_strategies" ADD FOREIGN KEY ("credential_id") REFERENCES "credentials" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "show_strategies" DROP CONSTRAINT IF EXISTS "credential_id";
DROP TABLE IF EXISTS "show_strategies";
-- +goose StatementEnd
