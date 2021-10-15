package win

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	var output string
	// add 3 to the score of the winner
	const WIN_SCORE = 3
	// create the object to track the results of thr competition
	competitionTracker := make(map[string]int)
	// functions that add names to a map
	nameMaker := func(tracker map[string]int,
		compNames [][]string) map[string]int {
		// loop over the 2d array of competitors
		for i, teams := range compNames {
			for j, team := range teams {
				// if this is the first row and first column
				// then add the name to the tracker
				if i == 0 && j == 0 {
					// fmt.Println("adding ", team, " to tracker")
					tracker[team] = 0
				} else {
					// check to see if each name is in map
					_, ok := tracker[team]
					if ok {
						// fmt.Println(team, " is in tracker ", ok)
					} else {
						// fmt.Println("adding ", team, " to tracker")
						tracker[team] = 0
					}
				}
			}
		}
		return tracker
	}
	// enter all unique names into map object with value of 0
	competitionTracker = nameMaker(competitionTracker, competitions)

	scoreTracker := func(tracker map[string]int,
		compRounds [][]string, result []int) map[string]int {
		for index, round := range compRounds {
			// fmt.Println("In this round: ", round, " the winner is: ", result[index])
			if result[index] == HOME_TEAM_WON {
				tracker[round[0]] = tracker[round[0]] + WIN_SCORE
			} else {
				tracker[round[1]] = tracker[round[1]] + WIN_SCORE
			}
		}
		return tracker
	}
	// add scores to winners at each round
	competitionTracker = scoreTracker(competitionTracker, competitions, results)

	// select the key with the highest score
	selectWinner := func(tracker map[string]int) string {
		var final string
		k := []string{}
		for key, _ := range tracker {
			k = append(k, key)
		}
		for ind, value := range k {
			// fmt.Println(tracker[value])
			if ind == 0 {
				final = value
			}
			if tracker[value] > tracker[final] {
				final = value
			}
		}
		return final
	}
	output = selectWinner(competitionTracker)
	return output
}

func TournamentWinner2(competitions [][]string, results []int) string {
	currentBestTeam := ""
	scores := map[string]int{currentBestTeam: 0}

	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}

	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}
