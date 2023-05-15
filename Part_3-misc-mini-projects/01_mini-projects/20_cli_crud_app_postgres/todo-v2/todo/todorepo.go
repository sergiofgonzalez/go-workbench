package todo

import (
	"database/sql"
	"fmt"
)

// Repository is a data structure that connects the ToDo model with the database
type Repository struct {
	db *sql.DB
}

// New returns a new repository for the ToDo model
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Serialize performs the creation of the table that holds ToDo tasks if it doesn't exist
func (r *Repository) Serialize() error {
	sql := `CREATE TABLE if not EXISTS todos (
		id serial PRIMARY KEY,
		title text NOT NULL,
		description text NOT NULL,
		done bool NULL
	)`
	if _, err := r.db.Exec(sql); err != nil {
		return (fmt.Errorf("Serialize: couldn't perform table init: %v", err))
	}
	return nil
}

// All returns all the ToDo tasks available in the database
func (r *Repository) All() ([]ToDo, error) {
	query := `SELECT id, title, description, done
	FROM todos`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("All failed: SELECT didn't complete correctly: %w", err)
	}
	defer rows.Close()

	todos := []ToDo{}
	for rows.Next() {
		todo := &ToDo{}
		err := rows.Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("All failed: couldn't scan row from result set: %w", err)
		}
		todos = append(todos, *todo)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("All failed: couldn't scan all rows: %v", rows.Err())
	}
	return todos, nil
}

// FindByID retrieves the corresponding ToDo task identified by the given id
func (r *Repository) FindByID(id int) (*ToDo, error) {
	query := `SELECT id, title, description, done
	FROM todos
 WHERE ID = $1`

	todo := ToDo{}
	err := r.db.QueryRow(query, id).Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
	switch {
	case err == sql.ErrNoRows:
		return &todo, ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("FindById failed: couldn't retrieve record from ToDo table: %w", err)
	default:
		return &todo, nil
	}
}

// FindByTitle retrieves the corresponding ToDo tasks identified by the given title
func (r *Repository) FindByTitle(title string) ([]ToDo, error) {
	query := `SELECT id, title, description, done
						  FROM todos
						 WHERE title = $1`

	rows, err := r.db.Query(query, title)
	if err != nil {
		return nil, fmt.Errorf("FindByTitle failed: SELECT didn't complete correctly: %w", err)
	}
	defer rows.Close()

	todos := []ToDo{}
	for rows.Next() {
		todo := &ToDo{}
		err := rows.Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("FindByTitle failed: couldn't scan row from result set: %w", err)
		}
		todos = append(todos, *todo)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("FindByTitle failed: couldn't scan all rows: %v", rows.Err())
	}
	return todos, nil
}

// Save inserts or updates the given ToDo record in the database
func (r *Repository) Save(t *ToDo) error {
	switch t.id.Valid {
	case true:
		err := r.doUpdate(t)
		if err != nil {
			return fmt.Errorf("Save failed: cannot update: %w", err)
		}
	case false:
		err := r.doInsert(t)
		if err != nil {
			return fmt.Errorf("Save failed: cannot insert: %w", err)
		}
	}
	return nil
}

// Delete deletes the corresponding ToDo record in the database
func (r *Repository) Delete(t *ToDo) error {
	query := `DELETE FROM todos
	           WHERE id = $1`

	_, err := r.db.Exec(query, &t.id)
	return err
}

func (r *Repository) doUpdate(t *ToDo) error {
	query := `UPDATE todos
	             SET title = $1,
							     description = $2,
									 done = $3
						 WHERE id = $4`

	_, err := r.db.Exec(query, &t.Title, &t.Description, &t.Done, &t.id)
	return err
}

func (r *Repository) doInsert(t *ToDo) error {
	query := `INSERT INTO todos
	            (title, description, done)
						VALUES
						  ($1, $2, $3)
						RETURNING id`

	err := r.db.QueryRow(query, &t.Title, &t.Description, &t.Done).Scan(&t.id)
	return err
}
