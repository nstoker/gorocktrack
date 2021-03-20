# gorocktrack

A project aimed at creating a Golang based tool to aid a guitarist to help focus on songs they are learning, without forgetting about songs they haven't played in long time.

## Database migration

Migrations are located in the `internal/pkg/migrations/files` directory.
Migrations are handled with use of the [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool.

To install the command line tool run:

```bash
go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
```

### Generate a new migration

Migrations are created using the `migrate` tool. You generate the migration files, then need to edit the files to add the actual migration.

```bash
migrate create -ext sql -dir internal/pkg/migrations/files create_users_table
ls internal/pkg/migrations/files
# 20191119133115_create_users_table.down.sql  20191119133115_create_users_table.up.sql

```

Edit the files (this example adds basic user details):

```bash
# internal/pkg/migrations/files/20191119110032_create_table.down.sql
DROP TABLE IF EXISTS users;
```

```bash
# internal/pkg/migrations/files/20191119110032_create_table.up.sql
CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(300),
    email VARCHAR(300) UNIQUE NOT NULL,
    password VARCHAR(500),
    updated_at BIGINT
);
```

### Applying migrations

This should #justWork on startup, however...

```bash
export POSTGRESQL_URL=postgres://postgres:password@localhost:5432/go_api?sslmode=disable
migrate -database ${POSTGRESQL_URL} -path internal/pkg/migrations/files up
```

and to undo all migrations (be very,very sure):

```bash
migrate -database ${POSTGRESQL_URL} -path internal/pkg/migrations/files down
```

## Environment variables

The project is intended to run on a Heroku server, and expects the following environment files:

|name|purpose|
|--|--|
| DATABASE_URL | The address the postgres server is listening at |
| PORT | The port the server will listen on |
| ADMIN_EMAIL | The admin's email address |
| ADMIN_NAME | The admin's name |
| ADMIN_PASS | The admin's password |

In development and test, copy the `example.env` to `.env` and `.test.env` and set the values appropriate for your systems.
