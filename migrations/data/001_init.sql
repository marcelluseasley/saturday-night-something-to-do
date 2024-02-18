CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    street_address TEXT,
    city TEXT,
    state TEXT,
    zip_code TEXT
);

---- create above / drop below ----

DROP TABLE users;
