CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    email citext UNIQUE NOT NULL,
    password VARCHAR,
    admin BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);
