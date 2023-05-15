package main

import (
	"fmt"
	"os"

	"example.com/todo/todo"
)

func main() {
	// make sure the database in which persist the entities is closed
	defer todo.Close()

	// Save a new ToDo task
	greetToDo := &todo.ToDo{
		Title:       "Say Hello",
		Description: "Start the morning saying hello to Jason Isaacs!",
		Done:        false,
	}

	fmt.Println(greetToDo)
	err := greetToDo.Save()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't save task: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greetToDo)

	// Update an existing task
	greetToDo.Done = true
	err = greetToDo.Save()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't save updated task: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greetToDo)

	// Find all tasks
	all, err := todo.All()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved stored tasks: %v\n", err)
		os.Exit(1)
	}

	var lastID int
	for _, t := range all {
		fmt.Println(&t)
		lastID = int(t.ID.Int16)
	}

	// Find a task by Id: found
	t, err := todo.FindByID(lastID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved task with id %d: %v\n", lastID, err)
		os.Exit(1)
	}
	fmt.Println(t)

	// Find a task by Id: not found
	t, err = todo.FindByID(-1)
	if err == todo.ErrNotFound {
		fmt.Printf("couldn't find task with id %d in the db\n", -1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved task with id %d: %v\n", lastID, err)
	}

	// Find by title: found some
	todos, err := todo.FindByTitle("Say Hello")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved tasks by title: %q\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d tasks with title %q\n", len(todos), "Say Hello")

	// Find by title: found none
	todos, err = todo.FindByTitle("Say goodbye")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved tasks by title: %q\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d tasks with title %q\n", len(todos), "Say Goodbye")

	// Clean the db
	for _, t := range all {
		err := t.Delete()
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't delete task with id %d: %q\n", t.ID.Int16, err)
			os.Exit(1)
		}
	}

	// // Connect to the database
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/tododb?sslmode=disable")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't connect to the database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer db.Close()

	// // Insert a new record
	// query := "INSERT INTO todos (title, description, done) VALUES ($1, $2, $3)"
	// _, err = db.Exec(query, "Write Go program", "Write a Go program using Postgres", true)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't insert a new ToDo task: %v\n", err)
	// 	os.Exit(1)
	// }

	// // Retrieving single result from db
	// query = "SELECT * FROM todos WHERE title = $1"

	// todo := ToDo{}
	// err = db.QueryRow(query, "Say hello to Jason").Scan(&todo.id, &todo.title, &todo.description, &todo.done)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't retrieve a task record from ToDo table: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("%+v\n", todo)

	// // Retrieving multiple results from db
	// query = "SELECT * FROM todos"

	// rows, err := db.Query(query)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't retrieve task records from ToDo table: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&todo.id, &todo.title, &todo.description, &todo.done)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "couldn't scan row from ToDo table: %v\n", err)
	// 	}
	// 	fmt.Printf("%+v\n", todo)
	// }

	// // Fixing the error: sql: Scan error on column index 3, name "done": sql/driver: couldn't convert <nil> (<nil>) into type bool
	// query = "SELECT * FROM todos WHERE done is NULL"
	// var safeBool sql.NullBool

	// rows, err = db.Query(query)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't retrieve task records where done is null from ToDo table: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&todo.id, &todo.title, &todo.description, &safeBool)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "couldn't scan row from ToDo table (will continue): %v\n", err)
	// 	}
	// 	if safeBool.Valid {
	// 		todo.done = safeBool.Bool
	// 	}
	// 	fmt.Printf("%+v\n", todo)
	// }

	// // Updating database records
	// query = "UPDATE todos SET done = false WHERE done is NULL"

	// _, err = db.Exec(query)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't update table: %v\n", err)
	// 	os.Exit(1)
	// }

	// // Deleting database records
	// query = "DELETE FROM todos WHERE title like $1"

	// _, err = db.Exec(query, "%Jason%")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "couldn't update table: %v\n", err)
	// 	os.Exit(1)
	// }
}
