package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// ToDo is a data structure modeling a task to be done
type ToDo struct {
	id          int
	title       string
	description string
	done        bool
}

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/tododb?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't connect to the database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Insert a new record
	query := "INSERT INTO todos (title, description, done) VALUES ($1, $2, $3)"
	_, err = db.Exec(query, "Write Go program", "Write a Go program using Postgres", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't insert a new ToDo task: %v\n", err)
		os.Exit(1)
	}

	// Retrieving single result from db
	query = "SELECT * FROM todos WHERE title = $1"

	todo := ToDo{}
	db.QueryRow(query, "Say hello to Jason").Scan(&todo.id, &todo.title, &todo.description, &todo.done)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieve a task record from ToDo table: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", todo)


	// Retrieving multiple results from db
	query = "SELECT * FROM todos"


	rows, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieve task records from ToDo table: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&todo.id, &todo.title, &todo.description, &todo.done)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't scan row from ToDo table: %v\n", err)
		}
		fmt.Printf("%+v\n", todo)
	}

	// Fixing the error: sql: Scan error on column index 3, name "done": sql/driver: couldn't convert <nil> (<nil>) into type bool
	query = "SELECT * FROM todos WHERE done is NULL"
	var safeBool sql.NullBool

	rows, err = db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieve task records where done is null from ToDo table: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&todo.id, &todo.title, &todo.description, &safeBool)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't scan row from ToDo table (will continue): %v\n", err)
		}
		if safeBool.Valid {
			todo.done = safeBool.Bool
		}
		fmt.Printf("%+v\n", todo)
	}

	// Updating database records
	query = "UPDATE todos SET done = false WHERE done is NULL"

	_, err = db.Exec(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't update table: %v\n", err)
		os.Exit(1)
	}

	// Deleting database records
	query = "DELETE FROM todos WHERE title like $1"

	_, err = db.Exec(query, "%Jason%")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't update table: %v\n", err)
		os.Exit(1)
	}
}
