package backtracking

// BacktrackState9 tracks the current state of a backtracking search
type BacktrackState9 struct {
	depth        int
	nodesVisited int
	pruneCount   int
}

// NewState9 initializes a fresh backtracking state tracker
func NewState9() *BacktrackState9 {
	return &BacktrackState9{depth: 0, nodesVisited: 0, pruneCount: 0}
}

// IncrementDepth9 moves one level deeper in the search tree
func (state *BacktrackState9) IncrementDepth9() {
	state.depth++
	state.nodesVisited++
}
