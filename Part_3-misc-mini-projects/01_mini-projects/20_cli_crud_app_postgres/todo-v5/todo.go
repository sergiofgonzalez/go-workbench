package todo

import (
	"database/sql"
	"fmt"
)

// ToDo models a task to be done
type ToDo struct {
	Title       string
	Description string
	Done        bool
	id          sql.NullInt16
}

// ID returns the associated identifier of the item if persisted, or an error otherwise
func (t *ToDo) ID() (int, error) {
	if !t.id.Valid {
		return -1, ErrUnpersisted
	}
	return int(t.id.Int16), nil
}

func (t *ToDo) String() string {
	var idStr string
	if t.id.Valid {
		idStr = fmt.Sprintf("%d", t.id.Int16)
	} else {
		idStr = "<NULL>"
	}
	return fmt.Sprintf("{id: %s, title: %q, description: %q, done: %v}", idStr, t.Title, t.Description, t.Done)
}
