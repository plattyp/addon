
-- +migrate Up
CREATE TABLE public.apps (
	"id" serial PRIMARY KEY,
	"name" varchar NOT NULL COLLATE "default",
	"user_id" int NOT NULL,
	"created_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"updated_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"deleted_at" timestamp(6) WITH TIME ZONE,
	CONSTRAINT "fk_apps_user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION NOT DEFERRABLE INITIALLY IMMEDIATE
);

-- +migrate Down
DROP TABLE IF EXISTS public.apps;
