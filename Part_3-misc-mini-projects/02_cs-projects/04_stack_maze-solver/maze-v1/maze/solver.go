package maze

import "fmt"

func expandNode(node Node, maze Matrix, frontier *NodeStack, exploredSet *ExploredSet) {
	findTransitionNodes := func(maze [][]string, state State) []Node {
		var transitionNodes []Node
		candidateMoves := []Action{
			{RowMove: -1, ColMove: 0},
			{RowMove: 0, ColMove: 1},
			{RowMove: 1, ColMove: 0},
			{RowMove: 0, ColMove: -1},
		}

		for _, m := range candidateMoves {
			nextRow := state.Row + m.RowMove
			nextCol := state.Col + m.ColMove
			if nextRow >= 0 && nextRow < len(maze) &&
				nextCol >= 0 && nextCol < len(maze[0]) &&
				maze[nextRow][nextCol] != Wall {
				transitionNodes = append(
					transitionNodes,
					Node{
						State:  State{nextRow, nextCol},
						Parent: &node,
						Action: &Action{m.RowMove, m.ColMove},
						Cost:   node.Cost + 1,
					})
			}
		}
		return transitionNodes
	}

	for _, transitionNode := range findTransitionNodes(maze, node.State) {
		if !exploredSet.contains(transitionNode) {
			frontier.Push(transitionNode)
		}
	}
}

// Solve takes a maze represented as a Matrix and returns its solution (if found)
func (m *Matrix) Solve() (solution []string) {
	initialState := m.findInitialState()

	startNode := Node{
		State:  initialState,
		Parent: nil,
		Action: nil,
		Cost:   0,
	}

	frontier := NewNodeStack()
	frontier.Push(startNode)
	exploredSet := &ExploredSet{}

	node := startNode
	var err error
	for {
		if frontier.Empty() {
			break
		}
		node, err = frontier.Pop()
		if err != nil {
			panic(fmt.Sprintf("Unexpected situation while searching: %v", err))
		}
		if m.testGoal(node.State) {
			break
		}
		exploredSet.add(node)
		expandNode(node, *m, frontier, exploredSet)
	}

	if !m.testGoal(node.State) {
		return nil
	}
	return buildSolution(node)
}

func buildSolution(node Node) (solution []string) {
	reversedMoves := NewStringStack()

	for node.Parent != nil {
		reversedMoves.Push(Elem(fmt.Sprintf("from (row=%d, col=%d), move (%d, %d) to (%d, %d) (cost=%d)", node.Parent.State.Row, node.Parent.State.Col, node.Action.RowMove, node.Action.ColMove, node.State.Row, node.State.Col, node.Cost)))
		node = *node.Parent
	}
	reversedMoves.Push("Solution found!!!")

	for !reversedMoves.Empty() {
		move, err := reversedMoves.Pop()
		if err != nil {
			panic("Error found while building the solution")
		}
		solution = append(solution, string(move))
	}
	return
}
