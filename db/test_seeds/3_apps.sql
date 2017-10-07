INSERT INTO "apps" (id, name, user_id)
VALUES (1, 'myapp', 123)
ON CONFLICT (id) DO UPDATE SET
  id = excluded.id,
  name = excluded.name,
  user_id = excluded.user_id;
