FROM postgres

RUN echo hello

# RUN psql -U $POSTGRES_USER -c '
#         CREATE TABLE users (
#          id serial PRIMARY KEY,
#          name text NOT NULL,
#          created_on timestamptz
#         );
# '
#
# RUN psql -U $POSTGRES_USER -c '
#         INSERT INTO users (name)
#         SELECT
#             'Dylan number ' || i
#         FROM generate_series(1, 5) as i;
# '
