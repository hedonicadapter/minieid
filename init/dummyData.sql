CREATE TABLE users (
    id serial PRIMARY KEY,
    name text NOT NULL,
    created_on timestamptz
);

INSERT INTO users (name)
SELECT
    'Dylan number ' || i
FROM generate_series(1, 5) as i;

