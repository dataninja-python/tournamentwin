package main

import (
	"fmt"
	"tournamentwin/win"
)

func main() {
	comp := [][]string{}
	r1 := []string{"HTML", "C#"}
	r2 := []string{"C#", "Python"}
	r3 := []string{"Python", "Html"}
	results1 := []int{0, 0, 1}
	fmt.Println(results1)
	comp = combineSlices(comp, r1, r2, r3)
	fmt.Println(comp)
	fmt.Println(win.TournamentWinner(comp, results1))
}

func combineSlices(holder [][]string, inputs ...[]string) [][]string {
	holder = append(holder, inputs...)
	return holder
}
