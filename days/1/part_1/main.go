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

	var elfWithMostCallories int
	var mostCalloriesCarried int
	for elf, foodItems := range foodCarriedByEachElf {
		calloriesCarried := 0
		for _, callories := range foodItems {
			calloriesCarried += callories
		}
		if calloriesCarried > mostCalloriesCarried {
			elfWithMostCallories = elf
			mostCalloriesCarried = calloriesCarried
		}
	}

	println("elf ", elfWithMostCallories, " carries most callories: ", mostCalloriesCarried)
}
