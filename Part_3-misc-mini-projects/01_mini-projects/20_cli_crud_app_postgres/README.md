# ToDo using Go and Postgres
> minimal ToDo app using Go and Postgres using the CLI

This project uses portions from https://dev.to/limaleandro1999/building-a-simple-crud-application-with-go-and-postgresql-27dk

## v0: setting up the stage

In this version we start setting up the Postgres database using Docker and making sure we can connect to it.

Follow the instructions of the [README.md](./postgres15-specs/README.md) to build and run the image.

Then, connect to it using a SQL client tool (such as DBeaver) and create the database using the tool:

![Create the db](./images/dbeaver.png)

Then, you can proceed to create and query a simple table to check everything is working as expected:

```sql
-- Create the todos table
CREATE TABLE todos (
	id serial PRIMARY KEY,
	title text NOT NULL,
	description text NOT NULL,
	done bool NOT NULL
);

-- Insert a few records
INSERT INTO todos
  (title, description)
VALUES
  ('Say hello to Jason', 'Say hello to Jason Isaacs!');

INSERT INTO todos
  (title, description, done)
VALUES
  ('Meet friends on Friday', 'Meet my friends on Friday evening', false);

INSERT INTO todos
  (title, description, done)
VALUES
  ('Set up Postgres 15', 'Use Docker to setup a Postgres instance', true);

-- Query results
SELECT * FROM todos
```

If everything goes according to the plan you'll be able to see some results on your SQL client tool and therefore can proceed to the application development.

## v1: refactoring into functions

In this version we need to start refactoring the code into functions, creating the servie layer.

## v2: adding tests

In this version we start adding tests using [sqlmock](https://github.com/DATA-DOG/go-sqlmock).

This will require a bit of refactoring to make sure we use DI, so that we can effectively mock the database access.

## v3: reviewing coverage and removing hardcoded conn

In this iteration we review the test coverage and add a few more tests to bump up the coverage beyond 90%.

Also, the connection string to Postgres is externalized in an environment variable: "POSTGRES_CONN_STRING".

Therefore, you can run the main program doing something like:

```bash
POSTGRES_CONN_STRING="postgres://postgres:postgres@localhost/tododb?sslmode=disable" go run .
```

## v4: Creating a CLI app

In this iteration, we create a CLI application so that records can be queried, created, retrieved, updated and deleted from the terminal window.

## v5: Wrapping it up

As the idea is not to create a production app, but rather understand how to access a DB in a Go program, etc., this last iteration simply cleans and polishes the example.

+ Better project organization
+ Better choice of packages to improve DX for things such as `todo.NewRepository(db)` and `todo.NewService(repo)`.


Let's start with the project organization. According to the notes on [02: Go Project Layout](../../../Part_2-advanced-concepts/02_go_prj_layout/concepts.ipynb), we're dealing with a quite complex CLI tool. As such, it contains certain libraries such as the Service and Repository layer, and also a CLI tool implemented in the `main.go`.

Eventually we will want to create a web app out of this code, so it makes sense to use the *Everything together* section of that document which proposes something like:

```
compress/                       # Root project directory
├── compress.go                 # Project's most relevant file
├── go.mod
├── go.sum
├── encode/                     # Library code for certain prj capability
│   └── encode.go
├── decode/                     # Another library for another prj capability
│   └── decode.go
├── internal/                   # Internal lib not to be exposed outside the prj
│   └── deflate/
|        └── deflate.go
├── cmd/                        # CLI application, uses same name as the prj
│   └── compress/
|        └── main.go
├── .gitignore
├── Makefile
└── README.md
```

However, we cannot create separate libraries for the Services and the Repositories because right now they are defined in the same package and use non-exported fields such as `ToDo.id`.

As making the `id` field public seems problematic, because we don't want to expose write permissions to the outside, it's better to keep them in the same package, so the reorganization will be minimal.

Also, because of the same reason, we won't be able to do something like `todorepo.New`, and will keep the previous approach: `todo.NewRepo`.


## Todo

- [ ] Add more tests
