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
		"X": R,
		"Y": P,
		"Z": S,
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
		opponent := roundHand[0]
		myHand := roundHand[1]
		if symbols[opponent] == symbols[myHand] {
			outcome += DRAW
		} else if symbols[opponent] == symbols[myHand]+1 || symbols[opponent] == symbols[myHand]-2 {
			outcome += LOSE
		} else {
			outcome += WIN
		}
		outcome += symbols[myHand]
	}

	println("score: ", outcome)
}
