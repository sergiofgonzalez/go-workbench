package todo_test

import (
	"database/sql/driver"
	"testing"

	"example.com/todo/todo"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSerialize(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error while opening the stubbed db connection", err)
	}
	defer db.Close()

	// Set the expectations for the DB interaction first
	sql := "CREATE TABLE (.+) todos"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute the method under test
	r := todo.New(db)

	// Assert method response and expectations
	if err = r.Serialize(); err != nil {
		t.Errorf("no errors expected but got %v", err)
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
			r := todo.New(db)

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
