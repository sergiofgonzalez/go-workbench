package gordle

import "strings"

type hint byte

// hint enumeration
const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		return "💔"	// this should not happen
	}
}

type feedback []hint

func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, hint := range fb {
		sb.WriteString(hint.String())
	}
	return sb.String()
}

func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for i := 0; i < len(fb); i++ {
		if fb[i] != other[i] {
			return false
		}
	}
	return true
}