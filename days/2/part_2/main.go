package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	R    = 1
	P    = 2
	S    = 3
	WIN  = 6
	LOSE = 0
	DRAW = 3
)

var (
	symbols = map[string]int{
		"A": R,
		"B": P,
		"C": S,
	}
	desiredOutcomes = map[string]string{
		"X": "LOSE",
		"Y": "DRAW",
		"Z": "WIN",
	}
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	var rounds [][]string
	s := bufio.NewScanner(f)
	for s.Scan() {
		round := strings.SplitN(s.Text(), " ", 2)
		rounds = append(rounds, round)
	}

	var outcome int
	for _, roundHand := range rounds {
		var myHand int
		opponent := roundHand[0]

		desiredOutcome := roundHand[1]
		if desiredOutcomes[desiredOutcome] == "DRAW" {
			myHand = symbols[opponent]
		} else if desiredOutcomes[desiredOutcome] == "LOSE" {
			myHand = symbols[opponent] - 1
			if myHand < 1 {
				myHand = 3
			}
		} else {
			myHand = symbols[opponent] + 1
			if myHand > 3 {
				myHand = 1
			}
		}

		if symbols[opponent] == myHand {
			outcome += DRAW
		} else if symbols[opponent] == myHand+1 || symbols[opponent] == myHand-2 {
			outcome += LOSE
		} else {
			outcome += WIN
		}
		outcome += myHand
	}

	println("score: ", outcome)
}
