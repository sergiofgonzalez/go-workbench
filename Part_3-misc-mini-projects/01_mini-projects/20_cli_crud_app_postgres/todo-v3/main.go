package main

import (
	"database/sql"
	"fmt"
	"os"

	"example.com/todo/todo"
	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver
)

func main() {
	postgresConnString, ok := os.LookupEnv("POSTGRES_CONN_STRING")
	if !ok {
		panic("cannot continue without a POSTGRES_CONN_STRING to connect to the database")
	}

	db, err := sql.Open("pgx", postgresConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't connect to the database: %v\n", err)
		os.Exit(1)
	}

	// make sure the database in which persist the entities is closed
	defer db.Close()

	// wire the database driver into the repository
	repo := todo.New(db)

	// set up the table if it doesn't exist
	repo.Serialize()

	// Find all tasks before inserting anything
	_, err = repo.All()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved stored tasks: %v\n", err)
		os.Exit(1)
	}

	// Save a new ToDo task
	greetToDo := &todo.ToDo{
		Title:       "Say Hello",
		Description: "Start the morning saying hello to Jason Isaacs!",
		Done:        false,
	}
	fmt.Println(greetToDo)
	err = repo.Save(greetToDo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't save task: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greetToDo)

	// Update an existing task
	greetToDo.Done = true
	err = repo.Save(greetToDo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't save updated task: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(greetToDo)

	// Find all tasks
	all, err := repo.All()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved stored tasks: %v\n", err)
		os.Exit(1)
	}

	var lastID int
	for _, t := range all {
		fmt.Println(&t)
		lastID, err = t.ID()
	}

	// Find a task by Id: found
	t, err := repo.FindByID(lastID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved task with id %d: %v\n", lastID, err)
		os.Exit(1)
	}
	fmt.Println(t)

	// Find a task by Id: not found
	t, err = repo.FindByID(-1)
	if err == todo.ErrNotFound {
		fmt.Printf("couldn't find task with id %d in the db\n", -1)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved task with id %d: %v\n", lastID, err)
	}

	// Find by title: found some
	todos, err := repo.FindByTitle("Say Hello")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved tasks by title: %q\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d tasks with title %q\n", len(todos), "Say Hello")

	// Find by title: found none
	todos, err = repo.FindByTitle("Say goodbye")
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved tasks by title: %q\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d tasks with title %q\n", len(todos), "Say Goodbye")

	// Delete an existing entry
	t, err = repo.FindByID(lastID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't retrieved task with id %d: %v\n", lastID, err)
		os.Exit(1)
	}
	err = repo.Delete(t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't delete existing task with id %d: %v", lastID, err)
	}

	// Delete a non-existing entry
	err = repo.Delete(t) // we delete the same entry again
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't delete existing task with id %d: %v", lastID, err)
	}

	// Clean the db
	for _, t := range all {
		err := repo.Delete(&t)
		if err != nil {
			id, _ := t.ID()
			fmt.Fprintf(os.Stderr, "couldn't delete task with id %d: %q\n", id, err)
			os.Exit(1)
		}
	}
}
