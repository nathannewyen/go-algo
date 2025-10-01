package backtracking

// BacktrackState10 tracks the current state of a backtracking search
type BacktrackState10 struct {
	depth        int
	nodesVisited int
	pruneCount   int
}

// NewState10 initializes a fresh backtracking state tracker
func NewState10() *BacktrackState10 {
	return &BacktrackState10{depth: 0, nodesVisited: 0, pruneCount: 0}
}

// IncrementDepth10 moves one level deeper in the search tree
func (state *BacktrackState10) IncrementDepth10() {
	state.depth++
	state.nodesVisited++
}
