package backtracking

// BacktrackState8 tracks the current state of a backtracking search
type BacktrackState8 struct {
	depth        int
	nodesVisited int
	pruneCount   int
}

// NewState8 initializes a fresh backtracking state tracker
func NewState8() *BacktrackState8 {
	return &BacktrackState8{depth: 0, nodesVisited: 0, pruneCount: 0}
}

// IncrementDepth8 moves one level deeper in the search tree
func (state *BacktrackState8) IncrementDepth8() {
	state.depth++
	state.nodesVisited++
}
