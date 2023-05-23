package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"example.com/todo/todo"
	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver
)

// Command-line flags
var (
	action = flag.String("action", "retrieve", "Action to perform")
	title = flag.String("title", "", "Title for your ToDo task")
	description = flag.String("description", "", "Description for your ToDo task")
	done = flag.Bool("done", false, "Status of your ToDo task")
	id = flag.Int64("id", -1, "The ToDo task ID to be acted upon")
)


func main() {
	db := dbConn()

	// make sure the database in which persist the entities is closed
	defer db.Close()

	// wire the database driver into the software layer
	repo := todo.NewRepository(db)
	svc := todo.NewService(repo)

	// set up the table if it doesn't exist
	repo.Serialize()

	flag.Parse()

	switch *action {
	case "create":
		handleCreate(svc)
	case "retrieve":
		handleRetrieve(svc)
	case "update":
		handleUpdate(svc)
	case "delete":
		handleDelete(svc)
	default:
		fmt.Fprintf(os.Stderr, "invalid action: must be one of create, retrieve, update, delete")
		os.Exit(1)
	}
}

func dbConn() (db *sql.DB) {
	postgresConnString, ok := os.LookupEnv("POSTGRES_CONN_STRING")
	if !ok {
		fmt.Fprint(os.Stderr, "error: missing POSTGRES_CONN_STRING environment variable\n")
		os.Exit(1)
	}

	db, err := sql.Open("pgx", postgresConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't connect to the database: %v\n", err)
		os.Exit(1)
	}
	return
}

func handleCreate(svc *todo.Service) {
	t := todo.ToDo{Title: *title, Description: *description, Done: *done}
	err := svc.Create(t)
	switch {
	case err == todo.ErrEmptyTitleAndDescription:
		doError("missing parameters --title and/or --description are required")
	case err != nil:
		doError(fmt.Sprintf("could not create ToDo task:\n%v\n", err))
	}
}

func handleRetrieve(svc *todo.Service) {
	var items []todo.ToDo
	var err error
	if *title == "" {
		items, err = svc.RetrieveAll()
		if err != nil {
			doError(fmt.Sprintf("could not retrieve ToDo tasks:\n%v", err))
		}
	} else {
		items, err = svc.Retrieve(*title)
		if err != nil {
			doError(fmt.Sprintf("could not retrieve ToDo tasks by title:\n%v", err))
		}
	}
	if len(items) == 0 {
		fmt.Println("-- nothing to retrieve")
		return
	}
	for _, item := range items {
		fmt.Println(&item)
	}
}

func handleUpdate(svc *todo.Service) {
	item := &todo.ToDo{
		Title: *title,
		Description: *description,
		Done: *done,
	}
	err := svc.Update(int(*id), item)
	switch {
	case err == todo.ErrEmptyID:
		doError("missing parameters: --id must be populated")
	case err == todo.ErrAtLeastOneFieldShouldBeUpdated:
		doError("nothing to change")
	case err == todo.ErrNotFound:
		doError("could not find ToDo task with given ID")
	case err != nil:
		doError(fmt.Sprintf("could not update ToDo task:\n%v", err))
	}
}

func handleDelete(svc *todo.Service) {
	err := svc.Delete(int(*id))
	switch {
	case err == todo.ErrEmptyID:
		doError("missing parameters: --id must be populated")
	case err == todo.ErrNotFound:
		doError("could not find ToDo task with given ID")
	case err != nil:
		doError(fmt.Sprintf("could not delete ToDo task:\n%v", err))
	}
}

func doError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}