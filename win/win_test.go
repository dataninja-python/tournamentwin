package win

import (
	"testing"
)

// create a constant of winners
func getWinnerResults() []string {
	return []string{"Python", "Java", "C#", "C#", "Bulls", "Eagles", "Bulls", "AlgoMasters", "SQL", "B"}
}

func TestTournamentWinner(t *testing.T) {
	winners := getWinnerResults()
	comp := [][]string{}
	r1 := []string{"HTML", "C#"}
	r2 := []string{"C#", "Python"}
	r3 := []string{"Python", "HTML"}
	results := []int{0, 0, 1}
	comp = combineSlices(comp, r1, r2, r3)
	testOutcome := TournamentWinner(comp, results)
	if testOutcome != winners[0] {
		t.Errorf("Winner was incorrect. Got: %s instead of: %s", testOutcome, winners[0])
	}
}

func combineSlices(holder [][]string, inputs ...[]string) [][]string {
	holder = append(holder, inputs...)
	return holder
}
