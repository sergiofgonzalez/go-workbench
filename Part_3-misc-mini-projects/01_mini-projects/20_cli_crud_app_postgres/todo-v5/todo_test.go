package todo_test

import (
	"database/sql/driver"
	"testing"

	"example.com/todo"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestID(t *testing.T) {
	t.Run("On unpersisted entity", func(t *testing.T) {
		todoItem := todo.ToDo{
			Title:       "A Title",
			Description: "A Description",
			Done:        false,
		}

		_, err := todoItem.ID()
		if err != todo.ErrUnpersisted {
			t.Errorf("expected error to be %v, but was %v", todo.ErrUnpersisted, err)
		}
	})

	t.Run("On persistend entity", func(t *testing.T) {
		// Establish the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		sql := "SELECT (.+) FROM todos WHERE ID = (.+)"

		expectedRow := []driver.Value{1, "Title 1", "Description 1", true}
		mockRows := sqlmock.NewRows([]string{"id", "title", "description", "done"}).
			AddRow(expectedRow...)

		mock.ExpectQuery(sql).WithArgs(1).WillReturnRows(mockRows)

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		todoItem, err := r.FindByID(1)
		if err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		assertResult(t, []todo.ToDo{*todoItem}, [][]driver.Value{expectedRow})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}

		// Now the ID test on persisted entity
		id, err := todoItem.ID()
		if err != nil {
			t.Errorf("expected no errors, but found %v", err)
		}
		if id != 1 {
			t.Errorf("expected id to be %d, but was %d", 1, id)
		}

	})
}

func TestString(t *testing.T) {
	t.Run("unpersisted entity", func(t *testing.T) {
		todoItem := todo.ToDo{
			Title:       "A Title",
			Description: "A Description",
			Done:        false,
		}

		expected := `{id: <NULL>, title: "A Title", description: "A Description", done: false}`
		if todoItem.String() != expected {
			t.Errorf("expected %s, but got %s", expected, todoItem.String())
		}
	})

	t.Run("persisted entity", func(t *testing.T) {
		// Establish the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		sql := "SELECT (.+) FROM todos WHERE ID = (.+)"

		expectedRow := []driver.Value{1, "Title 1", "Description 1", true}
		mockRows := sqlmock.NewRows([]string{"id", "title", "description", "done"}).
			AddRow(expectedRow...)

		mock.ExpectQuery(sql).WithArgs(1).WillReturnRows(mockRows)

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		todoItem, err := r.FindByID(1)
		if err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		assertResult(t, []todo.ToDo{*todoItem}, [][]driver.Value{expectedRow})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}

		// Now the ID test on persisted entity
		expected := `{id: 1, title: "Title 1", description: "Description 1", done: true}`
		if todoItem.String() != expected {
			t.Errorf("expected %s, but got %s", expected, todoItem.String())
		}
	})
}
