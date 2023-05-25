package todo_test

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"example.com/todo"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSerializeHappyPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	sql := "CREATE TABLE (.+) todos"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute the method under test
	r := todo.NewRepository(db)

	// Assert method response and expectations
	if err = r.Serialize(); err != nil {
		t.Errorf("no errors expected but got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestSerializeSadPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	sql := "CREATE TABLE (.+) todos"
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("some fabricated error"))

	// Execute the method under test
	r := todo.NewRepository(db)

	// Assert method response and expectations
	if err = r.Serialize(); err == nil {
		t.Error("an error was expected, but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestAllHappyPath(t *testing.T) {
	tests := map[string]struct {
		expectedRows [][]driver.Value
	}{
		"No rows": {
			expectedRows: nil,
		},
		"Single row": {
			expectedRows: [][]driver.Value{
				{1, "Title 1", "Description 1", true},
			},
		},
		"Multiple rows": {
			expectedRows: [][]driver.Value{
				{1, "Title 1", "Description 1", true},
				{2, "Title 2", "Description 2", true},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal("error while opening the stubbed db connection", err)
			}
			defer db.Close()

			// Set the expectations for the DB interaction first
			sql := "SELECT (.+) FROM todos"

			mockRows := sqlmock.NewRows([]string{"id", "title", "description", "done"})
			for _, row := range testCase.expectedRows {
				mockRows.AddRow(row...)
			}

			mock.ExpectQuery(sql).WillReturnRows(mockRows)

			// Execute the method under test
			r := todo.NewRepository(db)

			// Assert method response and expectations
			todos, err := r.All()
			if err != nil {
				t.Errorf("no errors expected but got %v", err)
			}

			assertResult(t, todos, testCase.expectedRows)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
			}
		})
	}
}

func TestAllSadPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	sql := "SELECT (.+) FROM todos"
	mock.ExpectQuery(sql).WillReturnError(fmt.Errorf("fabricated error"))

	// Execute the method under test
	r := todo.NewRepository(db)

	// Assert method response and expectations
	if _, err = r.All(); err == nil {
		t.Error("an error was expected but got none", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestFindByIdHappyPath(t *testing.T) {
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
}

func TestFindByIdSadPath(t *testing.T) {

	t.Run("no rows found", func(t *testing.T) {
		// Open the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		query := "SELECT (.+) FROM todos WHERE ID = (.+)"
		mock.ExpectQuery(query).WithArgs(1).WillReturnError(sql.ErrNoRows)

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if _, err = r.FindByID(1); err != todo.ErrNotFound {
			t.Errorf("error %v was expected but got %v", todo.ErrNotFound, err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("error was found", func(t *testing.T) {
		// Open the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		query := "SELECT (.+) FROM todos WHERE ID = (.+)"
		mock.ExpectQuery(query).WithArgs(1).WillReturnError(fmt.Errorf("fabricated error"))

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if _, err = r.FindByID(1); err == nil {
			t.Error("error was expected but got none")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestFindByTitleHappyPath(t *testing.T) {
	tests := map[string]struct {
		expectedRows [][]driver.Value
	}{
		"No rows": {
			expectedRows: nil,
		},
		"Single row": {
			expectedRows: [][]driver.Value{
				{1, "This is the title I was looking for", "Description 1", true},
			},
		},
		"Multiple rows": {
			expectedRows: [][]driver.Value{
				{1, "This is the title I was looking for", "Description 1", true},
				{2, "This is the title I was looking for", "Description 2", true},
			},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal("error while opening the stubbed db connection", err)
			}
			defer db.Close()

			// Set the expectations for the DB interaction first
			sql := "SELECT (.+) FROM todos WHERE title = (.+)"

			mockRows := sqlmock.NewRows([]string{"id", "title", "description", "done"})
			for _, row := range testCase.expectedRows {
				mockRows.AddRow(row...)
			}

			mock.ExpectQuery(sql).
				WithArgs("This is the title I was looking for").
				WillReturnRows(mockRows)

			// Execute the method under test
			r := todo.NewRepository(db)

			// Assert method response and expectations
			todos, err := r.FindByTitle("This is the title I was looking for")
			if err != nil {
				t.Errorf("no errors expected but got %v", err)
			}

			assertResult(t, todos, testCase.expectedRows)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %v", err)
			}
		})
	}
}

func TestFindByTitleSadPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	query := "SELECT (.+) FROM todos WHERE title = (.+)"
	mock.ExpectQuery(query).
		WithArgs("Some Title").
		WillReturnError(fmt.Errorf("fabricated error"))

	// Execute the method under test
	r := todo.NewRepository(db)

	// Assert method response and expectations
	if _, err = r.FindByTitle("Some Title"); err == nil {
		t.Error("an error was expected but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestSaveHappyPath(t *testing.T) {
	t.Run("Save a new entry", func(t *testing.T) {
		// Scenario 1: Saving a new entry
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		todoItem := todo.ToDo{
			Title:       "New Title",
			Description: "New Description",
			Done:        false,
		}

		// Set the expectations for the DB interaction first
		sql := "INSERT INTO todos (.+)"
		mock.ExpectQuery(sql).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).FromCSVString("1"))

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if err = r.Save(&todoItem); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		assertResult(t, []todo.ToDo{todoItem}, [][]driver.Value{{1, "New Title", "New Description", false}})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Update an existing entry", func(t *testing.T) {
		// Scenario 1: Saving a new entry
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		todoItem := todo.ToDo{
			Title:       "Title To Update",
			Description: "Description To Update",
			Done:        false,
		}

		// Set the expectations for the DB interaction first
		mock.ExpectQuery("INSERT INTO todos (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).FromCSVString("1"))
		mock.ExpectExec("UPDATE todos SET (.+) WHERE id = (.+)").
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Execute the method under test
		r := todo.NewRepository(db)
		if err = r.Save(&todoItem); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		todoItem.Title = "Updated Title"
		todoItem.Description = "Updated Description"
		todoItem.Done = true

		// Assert method response and expectations
		if err = r.Save(&todoItem); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		assertResult(t, []todo.ToDo{todoItem}, [][]driver.Value{{1, "Updated Title", "Updated Description", true}})

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestSaveSadPath(t *testing.T) {
	t.Run("Save a new entry fails", func(t *testing.T) {
		// Scenario 1: Saving a new entry
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		todoItem := todo.ToDo{
			Title:       "New Title",
			Description: "New Description",
			Done:        false,
		}

		// Set the expectations for the DB interaction first
		sql := "INSERT INTO todos (.+)"
		mock.ExpectQuery(sql).WillReturnError(fmt.Errorf("some fabricated error"))

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if err = r.Save(&todoItem); err == nil {
			t.Error("an error was expected but found none")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Update an existing entry fails", func(t *testing.T) {
		// Scenario 1: Saving a new entry
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		todoItem := todo.ToDo{
			Title:       "Title To Update",
			Description: "Description To Update",
			Done:        false,
		}

		// Set the expectations for the DB interaction first
		mock.ExpectQuery("INSERT INTO todos (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).FromCSVString("1"))
		mock.ExpectExec("UPDATE todos SET (.+) WHERE id = (.+)").
			WillReturnError(fmt.Errorf("some fabricated error"))

		// Execute the method under test
		r := todo.NewRepository(db)
		if err = r.Save(&todoItem); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		todoItem.Title = "Updated Title"
		todoItem.Description = "Updated Description"
		todoItem.Done = true

		// Assert method response and expectations
		if err = r.Save(&todoItem); err == nil {
			t.Error("an error was expected, but got none")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestDeleteHappyPath(t *testing.T) {
	t.Run("Delete existing instance", func(t *testing.T) {
		// Establish the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		sql := "DELETE FROM todos WHERE id = (.+)"
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 1))

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if err = r.Delete(&todo.ToDo{}); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})

	t.Run("Delete a non existing instance", func(t *testing.T) {
		// Establish the mock connection
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal("error while opening the stubbed db connection", err)
		}
		defer db.Close()

		// Set the expectations for the DB interaction first
		sql := "DELETE FROM todos WHERE id = (.+)"
		mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 0))

		// Execute the method under test
		r := todo.NewRepository(db)

		// Assert method response and expectations
		if err = r.Delete(&todo.ToDo{}); err != nil {
			t.Errorf("no errors expected but got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	})
}

func TestDeleteSadPath(t *testing.T) {
	// Establish the mock connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	sql := "DELETE FROM todos WHERE id = (.+)"
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("some fabricated error"))

	// Execute the method under test
	r := todo.NewRepository(db)

	// Assert method response and expectations
	if err = r.Delete(&todo.ToDo{}); err == nil {
		t.Error("an error was expected but got none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func assertResult(t testing.TB, gotItems []todo.ToDo, expectedItems [][]driver.Value) bool {
	t.Helper()

	if len(gotItems) != len(expectedItems) {
		return false
	}

	for i := 0; i < len(gotItems); i++ {
		id, err := gotItems[i].ID()
		if err != nil {
			return false
		}

		if gotItems[i].Title != expectedItems[i][1] ||
			gotItems[i].Description != expectedItems[i][2] ||
			gotItems[i].Done != expectedItems[i][3] ||
			id != expectedItems[i][0] {
			return false
		}
	}

	return true
}
