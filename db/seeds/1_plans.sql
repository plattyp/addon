INSERT INTO "plans" (id, code, name, description, ordinal)
VALUES (1, 'test', 'Test', 'This is the test plan', 1)
ON CONFLICT (id) DO UPDATE SET
  id = excluded.id,
  code = excluded.code,
  name = excluded.name,
  description = excluded.description,
  ordinal = excluded.ordinal;
