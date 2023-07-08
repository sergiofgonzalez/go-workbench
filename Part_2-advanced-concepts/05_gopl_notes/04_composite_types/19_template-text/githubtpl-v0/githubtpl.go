package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

// IssuesURL is the endpoint for querying issues at Github
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult models the response obtained from Github
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue models the information we want to retrieve from the items returned from Github
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User models the information we want to retrieve from a Github user
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}


// SearchIssues performs a search for the given terms on Github issues endpoint
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(fmt.Sprintf("%s?q=%s", IssuesURL, q))
	if err != nil {
		return nil, err
	}

	// it would have been better to use defer
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	// everything went well
	resp.Body.Close()
	return &result, nil
}

// daysAgo convert the given time into the equivalent elapsed days
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}
`

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))


func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("No args received: assuming 'repo:golang/go is:open json decoder'")
		args = []string{"repo:golang/go", "is:open", "json", "decoder"}
	}
	result, err := SearchIssues(args)
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
