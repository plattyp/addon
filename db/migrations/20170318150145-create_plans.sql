
-- +migrate Up
CREATE TABLE public.plans (
	"id" serial PRIMARY KEY,
	"code" varchar(100) NOT NULL collate "default",
	"name" varchar NOT NULL COLLATE "default",
	"description" varchar COLLATE "default",
	"ordinal" int DEFAULT 99,
	"created_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"updated_at" timestamp(6) WITH TIME ZONE NOT NULL DEFAULT (now() AT TIME ZONE 'UTC'),
	"deleted_at" timestamp(6) WITH TIME ZONE
);

-- +migrate Down
DROP TABLE IF EXISTS public.plans;
