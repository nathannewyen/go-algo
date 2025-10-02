package greedy

// GreedyChoiceTracker8 monitors greedy decisions made during execution
type GreedyChoiceTracker8 struct {
	choicesMade     int
	optimalChoices  int
}

// TrackChoice8 records a greedy choice and whether it was locally optimal
func (tracker *GreedyChoiceTracker8) TrackChoice8(isOptimal bool) {
	tracker.choicesMade++
	if isOptimal {
		tracker.optimalChoices++
	}
}

// OptimalityRatio8 returns the fraction of choices that were locally optimal
func (tracker *GreedyChoiceTracker8) OptimalityRatio8() float64 {
	if tracker.choicesMade == 0 {
		return 0
	}
	return float64(tracker.optimalChoices) / float64(tracker.choicesMade)
}
