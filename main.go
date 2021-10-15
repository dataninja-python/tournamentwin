package main

import (
	"fmt"
	"log"
	"time"
	"tournamentwin/win"
)

func combineSlices(holder [][]string, inputs ...[]string) [][]string {
	holder = append(holder, inputs...)
	return holder
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func executeCode() string {
	defer timeTrack(time.Now(), "tournament1")
	comp := [][]string{}
	r1 := []string{"Bulls", "Eagles"}
	r2 := []string{"Bulls", "Bears"}
	r3 := []string{"Bears", "Eagles"}
	results1 := []int{0, 0, 1}
	comp = combineSlices(comp, r1, r2, r3)
	output := win.TournamentWinner(comp, results1)
	fmt.Println(output)
	return output
}

func executeCode2() string {
	defer timeTrack(time.Now(), "tournament1")
	comp := [][]string{}
	r1 := []string{"Bulls", "Eagles"}
	r2 := []string{"Bulls", "Bears"}
	r3 := []string{"Bears", "Eagles"}
	results1 := []int{0, 0, 0}
	comp = combineSlices(comp, r1, r2, r3)
	output := win.TournamentWinner2(comp, results1)
	fmt.Println(output)
	return output
}

func main() {
	executeCode()
	executeCode2()
}
