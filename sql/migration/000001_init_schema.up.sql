
CREATE TABLE "person" (
  "id" bigserial PRIMARY KEY,
  "telegram_id" bigint NOT NULL UNIQUE,
  "first_name" varchar(128) NOT NULL,
  "last_name" varchar(128) NOT NULL,
  "language" varchar(16) NOT NULL,
  "telegram_language" varchar(16) NOT NULL,
  "main_password_hash" varchar(128) NOT NULL,
  "fake_password_hash" varchar(128) NOT NULL,
  "location_lat" bigint NOT NULL,
  "location_lon" bigint NOT NULL,
  "is_admin" boolean,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "person" ("telegram_id");