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

	// ErrEmptyTitleAndDescription is returned when Title and Description must be populated but were found to be empty
	ErrEmptyTitleAndDescription = todoErr("Title and Description must be populated")

	// ErrEmptyID is returned when ID must be populated but wasn't
	ErrEmptyID = todoErr("ID must be populated")

	// ErrAtLeastOneFieldShouldBeUpdated is returned when an update was requested but no changes with the save entity is detected
	ErrAtLeastOneFieldShouldBeUpdated = todoErr("at least one field should be updated")
)
