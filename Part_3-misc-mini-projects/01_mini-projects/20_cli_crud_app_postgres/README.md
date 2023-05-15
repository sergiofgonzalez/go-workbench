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

This will require a bit of refactoring to make sure we use DI, so that we can effectively mock.



## Todo
- [] All error use case
- [ ] Start refactoring, forget about the db and create a service and entity layers, then maybe it can be simplified. Services should be stateless.

- [ ] Include tests with sqlmock
- [ ] Add CLI with commands that invoke the services
- [ ] Add config so that db connection and stuff is externalized