package main

import (
	"bufio"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	var backpacks [][]int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var backpack []int
		for _, c := range s.Text() {
			backpack = append(backpack, int(c))
		}
		sort.Ints(backpack)
		backpacks = append(backpacks, backpack)
	}
	var groups [][][]int
	for g := 1; g < len(backpacks); g += 3 {
		groups = append(groups, backpacks[g-1:g+2])
	}

	var total int
	for _, group := range groups {
		var mutual int
	FindMutual:
		for _, firstGroup := range group[0] {
		SecondGroup:
			for _, secondGroup := range group[1] {
				if firstGroup < secondGroup {
					// its sorted so we can skip the rest
					continue FindMutual
				}
				for _, thirdGroup := range group[2] {
					if firstGroup == secondGroup && firstGroup == thirdGroup {
						mutual = firstGroup
						break FindMutual
					} else if firstGroup < thirdGroup {
						// its sorted so we can skip the rest
						continue SecondGroup
					}
				}
			}
		}

		if mutual > 90 {
			mutual -= 96
		} else {
			mutual -= 38
		}
		total += mutual
	}

	println("total: ", total)
}
