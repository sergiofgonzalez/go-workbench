package todo

type todoErr string

func (e todoErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound is returned when the corresponding record could not be found int he database
	ErrNotFound = todoErr("can't find task in the database")

	// ErrUnpersisted is returned when client code tries to read the ID from an item that has not been persisted
	ErrUnpersisted = todoErr("can't obtain ID from unpersisted item")
)