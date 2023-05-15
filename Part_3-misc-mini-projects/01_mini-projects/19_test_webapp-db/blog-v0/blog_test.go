package blog_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/blog"
	"github.com/DATA-DOG/go-sqlmock"
)

func assertJSON(actual []byte, data interface{}, t testing.TB) {
	t.Helper()
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("unexpected error while marshaling expected json data: %v", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("wanted %s, but got %s", expected, actual)
	}
}

func TestGetPostsHappyPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error was found opening the stubbed db connection: %v", err)
	}
	defer db.Close()

	// create the api with mocked db, request and response
	api := blog.NewAPI(db)
	req, err := http.NewRequest("GET", "http://localhosts/posts", nil)
	if err != nil {
		t.Fatalf("An error was found when creating test request: %v", err)
	}
	w := httptest.NewRecorder()

	// we create the expectations of the DB actions before executing the API
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello 1").
		AddRow(2, "post 2", "hello 2")

	mock.ExpectQuery("^SELECT (.+) FROM posts").WillReturnRows(rows)

	// we execute the request
	api.Posts(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status code to be %d, but was %d", http.StatusOK, w.Code)
	}

	data := struct {
		Posts []*blog.Post
	}{
		Posts: []*blog.Post{
			{ID: 1, Title: "post 1", Body: "hello 1"},
			{ID: 2, Title: "post 2", Body: "hello 2"},
		},
	}
	assertJSON(w.Body.Bytes(), data, t)

	// make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations on db test: %v", err)
	}
}

func TestGetPostsSadPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error was found opening the stubbed db connection: %v", err)
	}
	defer db.Close()

	// create the api with mocked db, request and response
	api := blog.NewAPI(db)
	req, err := http.NewRequest("GET", "http://localhosts/posts", nil)
	if err != nil {
		t.Fatalf("An error was found when creating test request: %v", err)
	}
	w := httptest.NewRecorder()

	// we create the expectations of the DB actions before executing the API
	mock.ExpectQuery("^SELECT (.+) FROM posts").WillReturnError(fmt.Errorf("fabricated error for testing"))

	// now we execute our request
	api.Posts(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status code to be %d, but was %d", http.StatusInternalServerError, w.Code)
	}

	data := struct {
		Error string
	}{
		Error: "Failed to fetch posts: fabricated error for testing",
	}

	assertJSON(w.Body.Bytes(), data, t)

	// make sure the DB expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations on db test: %v", err)
	}
}
