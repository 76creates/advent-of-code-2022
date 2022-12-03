package main

import (
	"bufio"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	var backpacks [][][]int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var backpack []int
		for _, c := range s.Text() {
			backpack = append(backpack, int(c))
		}
		var firstHalf, secondHalf []int
		firstHalf, secondHalf = backpack[:len(backpack)/2], backpack[len(backpack)/2:]
		sort.Ints(firstHalf)
		sort.Ints(secondHalf)
		backpacks = append(backpacks, [][]int{firstHalf, secondHalf})
	}

	var total int
	for _, backpack := range backpacks {
		var mutual int
	FindMutual:
		for _, firstHalf := range backpack[0] {
			for _, secondHalf := range backpack[1] {
				if firstHalf == secondHalf {
					mutual = firstHalf
					break FindMutual
				} else if firstHalf < secondHalf {
					// its sorted so we can skip the rest
					continue FindMutual
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
