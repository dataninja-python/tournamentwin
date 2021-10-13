package win

import "fmt"

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	output := "Good Bye"
	fmt.Println(competitions[0][0])
	fmt.Println(results[1])
	fmt.Println(HOME_TEAM_WON)
	return output
}
