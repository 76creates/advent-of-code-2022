package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	var elfPairs [][][]int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var elfPair [][]int
		elfs := strings.SplitN(s.Text(), ",", 2)
		for _, elf := range elfs {
			elfRange := strings.SplitN(elf, "-", 2)
			from, err := strconv.Atoi(elfRange[0])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(elfRange[1])
			if err != nil {
				panic(err)
			}
			elfPair = append(elfPair, []int{from, to})
		}
		elfPairs = append(elfPairs, elfPair)
	}

	lessElfs := 0
	for _, elfPair := range elfPairs {
		if elfPair[0][0] >= elfPair[1][0] && elfPair[0][1] <= elfPair[1][1] {
			lessElfs++
		} else if elfPair[1][0] >= elfPair[0][0] && elfPair[1][1] <= elfPair[0][1] {
			lessElfs++
		}
	}
	println("less elfs: ", lessElfs)
}
