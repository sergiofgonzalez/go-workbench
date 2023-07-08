package maze

// ExploredSet keeps track of the already visited states to prevent loops
type ExploredSet []State

func (e *ExploredSet) add(n Node) {

	if !e.contains(n) {
		*e = append(*e, n.State)
	}
}

func (e ExploredSet) contains(n Node) bool {
	for _, s := range e {
		if s == n.State {
			return true
		}
	}
	return false
}