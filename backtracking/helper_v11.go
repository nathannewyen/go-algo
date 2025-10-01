package backtracking

// BacktrackState11 tracks the current state of a backtracking search
type BacktrackState11 struct {
	depth        int
	nodesVisited int
	pruneCount   int
}

// NewState11 initializes a fresh backtracking state tracker
func NewState11() *BacktrackState11 {
	return &BacktrackState11{depth: 0, nodesVisited: 0, pruneCount: 0}
}

// IncrementDepth11 moves one level deeper in the search tree
func (state *BacktrackState11) IncrementDepth11() {
	state.depth++
	state.nodesVisited++
}
