package blog

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// API is the data structure that binds the APIs and Repositories of the application
type API struct {
	db *sql.DB
}

// Post models a blog post
type Post struct {
	ID    int
	Title string
	Body  string
}

// NewAPI returns a new API pointer with the given db connection
func NewAPI(db *sql.DB) *API {
	return &API{db: db}
}

// Posts returns an HTTP response with all the posts
func (a *API) Posts(w http.ResponseWriter, r *http.Request) {
	rows, err := a.db.Query("SELECT id, title, body FROM posts")
	if err != nil {
		a.fail(w, "Failed to fetch posts: " + err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		p := &Post{}
		if err := rows.Scan(&p.ID, &p.Title, &p.Body); err != nil {
			a.fail(w, "Failed to scan post rows: " + err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}
	if rows.Err() != nil {
		a.fail(w, "Failed to read all posts: " + rows.Close().Error(), http.StatusInternalServerError)
		return
	}

	data := struct{
		Posts []*Post
	}{posts}
	a.ok(w, data)
}

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost/blogdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	api := NewAPI(db)
	http.HandleFunc("/posts", api.Posts)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func (a *API) fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct{
		Error string
	}{
		Error: msg,
	}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(resp)
}

func (a *API) ok(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(data)
	if err != nil {
		a.fail(w, "Could not serialize response", http.StatusInternalServerError)
		return
	}
	w.Write(resp)	// if not calling WriteHeader StatusOK is returned
}
