package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	var foodCarriedByEachElf [][]int
	var elfItems []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Text()) == 0 {
			foodCarriedByEachElf = append(foodCarriedByEachElf, elfItems)
			elfItems = []int{}
		}
		i, _ := strconv.Atoi(s.Text())
		elfItems = append(elfItems, i)
	}
	foodCarriedByEachElf = append(foodCarriedByEachElf, elfItems)

	top3 := []int{0, 0, 0}
	for _, foodItems := range foodCarriedByEachElf {
		calloriesCarried := 0
		for _, callories := range foodItems {
			calloriesCarried += callories
		}
		top3 = placeInTop3(top3, calloriesCarried)
	}

	println("top calories carried: ", top3[0], ", ", top3[1], ", ", top3[2])
	println("summed top 3: ", top3[0]+top3[1]+top3[2])
}

func placeInTop3(top3 []int, check int) []int {
	for i := 0; i < 3; i++ {
		if check > top3[i] {
			replaced := top3[i]
			top3[i] = check
			return placeInTop3(top3, replaced)
		}
	}
	return top3
}
