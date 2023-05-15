package todo

type todoErr string

func (e todoErr) Error() string {
	return string(e)
}

const (
	// ErrDBConnection is returned when you can't establish a connection to the database
	ErrDBConnection = todoErr("can't connect to ToDo database")

	// ErrNotFound is returned when the corresponding record could not be found int he database
	ErrNotFound = todoErr("can't find task in the database")
)