
-- +migrate Up
CREATE TABLE public.users (
	"id" serial PRIMARY KEY,
	"email" varchar(255) NOT NULL COLLATE "default",
	"encrypted_password" text COLLATE "default",
	"password_salt" text COLLATE "default",
  "created_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"updated_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"deleted_at" timestamp(6) WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS public.users;
