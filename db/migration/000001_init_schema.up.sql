CREATE TABLE "tokens" (
  "jti" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "audience" varchar NOT NULL,
  "token" varchar NOT NULL,
  "expires_in" int NOT NULL,
  "scope" varchar[],
  "authorities" varchar[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "tokens" ("audience");

COMMENT ON COLUMN "tokens"."jti" IS 'uuid primary key';
COMMENT ON COLUMN "tokens"."expires_in" IS 'expires in should be in seconds';