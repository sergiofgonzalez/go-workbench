package todo

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("pgx", "postgres://postgres:postgres@localhost/tododb?sslmode=disable")
	if err != nil {
		panic(ErrDBConnection)
	}

	sql := `CREATE TABLE if not EXISTS todos (
		        id serial PRIMARY KEY,
		        title text NOT NULL,
		        description text NOT NULL,
		        done bool NULL
	        );`
	_, err = db.Exec(sql)
	if err != nil {
		panic(fmt.Errorf("couldn't confirm ToDo table existence: %w", err))
	}
}

// Close closes the underlying database connection used to persist the model entities
func Close() {
	db.Close()
}

// ToDo is a data structure modeling a task to be done
type ToDo struct {
	ID          sql.NullInt16
	Title       string
	Description string
	Done        bool
}

// All returns all todos available
func All() ([]ToDo, error) {
	query := `SELECT id, title, description, done
	            FROM todos`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("FindAll failed: SELECT didn't complete correctly: %w", err)
	}
	defer rows.Close()

	todos := []ToDo{}
	for rows.Next() {
		todo := &ToDo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("FindAll failed: couldn't scan row from result set: %w", err)
		}
		todos = append(todos, *todo)
	}
	return todos, nil
}

// FindByID retrieves the corresponding ToDo task identified by the given id
func FindByID(id int) (*ToDo, error) {
	query := `SELECT id, title, description, done
						  FROM todos
						 WHERE ID = $1`

	todo := ToDo{}
	err := db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	switch {
	case err == sql.ErrNoRows:
		return nil, ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("FindById failed: couldn't retrieve record from ToDo table: %w", err)
	default:
		return &todo, nil
	}
}

// FindByTitle retrieves the corresponding ToDo tasks identified by the given title
func FindByTitle(title string) ([]ToDo, error) {
	query := `SELECT id, title, description, done
						  FROM todos
						 WHERE title = $1`

	rows, err := db.Query(query, title)
	if err != nil {
		return nil, fmt.Errorf("FindByTitle failed: SELECT didn't complete correctly: %w", err)
	}
	defer rows.Close()

	todos := []ToDo{}
	for rows.Next() {
		todo := &ToDo{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("FindByTitle failed: couldn't scan row from result set: %w", err)
		}
		todos = append(todos, *todo)
	}
	return todos, nil
}

// Save inserts or updates a ToDo record in the database
func (t *ToDo) Save() error {
	switch t.ID.Valid {
	case true:
		err := t.doUpdate()
		if err != nil {
			return fmt.Errorf("Save failed: cannot update: %w", err)
		}
	case false:
		err := t.doInsert()
		if err != nil {
			return fmt.Errorf("Save failed: cannot insert: %w", err)
		}
	}
	return nil
}

// Delete deletes the corresponding ToDo record in the database
func (t *ToDo) Delete() error {
	query := `DELETE FROM todos
	           WHERE id = $1`

	_, err := db.Exec(query, &t.ID)
	return err
}

func (t *ToDo) doUpdate() error {
	query := `UPDATE ToDos
	             SET title = $1,
							     description = $2,
									 done = $3
						 WHERE id = $4`

	_, err := db.Exec(query, &t.Title, &t.Description, &t.Done, &t.ID)
	return err
}

func (t *ToDo) doInsert() error {
	query := `INSERT INTO ToDos
	            (title, description, done)
						VALUES
						  ($1, $2, $3)
						RETURNING id`

	err := db.QueryRow(query, &t.Title, &t.Description, &t.Done).Scan(&t.ID)
	return err
}

func (t *ToDo) String() string {
	var idStr string
	if t.ID.Valid {
		idStr = fmt.Sprintf("%d", t.ID.Int16)
	} else {
		idStr = "<NULL>"
	}
	return fmt.Sprintf("{id: %s, title: %q, description: %q, done: %v}", idStr, t.Title, t.Description, t.Done)
}
