
-- +migrate Up
CREATE TABLE public.users (
	"id" serial PRIMARY KEY,
	"email" varchar(255) COLLATE "default",
	"encrypted_password" text COLLATE "default",
	"password_salt" text COLLATE "default",
	"plan_id" int NOT NULL,
	"region" varchar(100) COLLATE "default",
	"heroku_uuid" varchar(60) COLLATE "default",
	"created_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"updated_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"deleted_at" timestamp(6) WITH TIME ZONE,
	CONSTRAINT "fk_users_plan_id" FOREIGN KEY ("plan_id") REFERENCES "public"."plans" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE
);

-- +migrate Down
DROP TABLE IF EXISTS public.users;
