INSERT INTO "plans" (id, code, name, description, ordinal)
VALUES (1, 'test', 'Test', 'This is the test plan', 1)
ON CONFLICT (id) DO UPDATE SET
  id = excluded.id,
  code = excluded.code,
  name = excluded.name,
  description = excluded.description,
  ordinal = excluded.ordinal;

INSERT INTO "plans" (id, code, name, description, ordinal)
VALUES (2, 'foo', 'Foo', 'This is the test plan for upgrades/downgrades', 2)
ON CONFLICT (id) DO UPDATE SET
  id = excluded.id,
  code = excluded.code,
  name = excluded.name,
  description = excluded.description,
  ordinal = excluded.ordinal;
