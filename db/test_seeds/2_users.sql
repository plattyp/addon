INSERT INTO "users" (id, email, plan_id)
VALUES (123, 'username@example.com', 1)
ON CONFLICT (id) DO UPDATE SET
  id = excluded.id,
  email = excluded.email,
  plan_id = excluded.plan_id;
