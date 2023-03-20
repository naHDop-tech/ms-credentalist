-- +goose Up
-- +goose StatementBegin
CREATE TABLE "customer_auth" (
    "id" uuid PRIMARY KEY,
    "customer_id" uuid NOT NULL,
    "is_verified" boolean NOT NULL,
    "otp" varchar(50) NOT NULL,
    "channel" varchar(50) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "customer_auth" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "customer_auth" DROP CONSTRAINT IF EXISTS "customer_id";
DROP TABLE IF EXISTS "auth_settings";
-- +goose StatementEnd
