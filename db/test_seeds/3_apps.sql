INSERT INTO "apps" (name, user_id)
VALUES ('myapp', 123)
ON CONFLICT (id) DO UPDATE SET
  name = excluded.name,
  user_id = excluded.user_id;
