package main

import (
	"bufio"
	"os"
	"strconv"
)

var dirs = make(map[string]map[string]int)
var currentDir []string

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var rows [][]int
	for s.Scan() {
		line := s.Text()
		var row []int
		for _, c := range line {
			i, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, i)
		}
		rows = append(rows, row)
	}

	greatestScenicScore := 0
	for y, row := range rows {
		if y == 0 || y == len(rows)-1 {
			continue
		}
		for x, _ := range row {
			if x == 0 || x == len(row)-1 {
				continue
			}
			score := getScenicScore(rows, x, y)
			println("x:", x, "y:", y, "score:", score)
			if score > greatestScenicScore {
				greatestScenicScore = score
			}
		}
	}
	println("greatest scenic score:", greatestScenicScore)

}

func getScenicScore(rows [][]int, x, y int) (score int) {
	return visibleTreesFromLeft(rows, x, y) * visibleTreesFromRight(rows, x, y) * visibleTreesFromAbove(rows, x, y) * visibleTreesFromBellow(rows, x, y)
}

func visibleTreesFromLeft(rows [][]int, x, y int) (visibleTrees int) {
	lookX := x - 1
	referenceTree := rows[y][x]
	for {
		if lookX < 0 { // out of bounds
			break
		}
		visibleTrees++
		if rows[y][lookX] >= referenceTree {
			return
		}
		lookX--
	}
	return
}

func visibleTreesFromRight(rows [][]int, x, y int) (visibleTrees int) {
	lookX := x + 1
	referenceTree := rows[y][x]
	rowLen := len(rows[y])
	for {
		if lookX >= rowLen { // out of bounds
			break
		}
		visibleTrees++
		if rows[y][lookX] >= referenceTree {
			return
		}
		lookX++
	}
	return
}

func visibleTreesFromAbove(rows [][]int, x, y int) (visibleTrees int) {
	lookY := y - 1
	referenceTree := rows[y][x]

	for {
		if lookY < 0 { // out of bounds
			break
		}
		visibleTrees++
		if rows[lookY][x] >= referenceTree {
			return
		}
		lookY--
	}
	return
}

func visibleTreesFromBellow(rows [][]int, x, y int) (visibleTrees int) {
	lookY := y + 1
	referenceTree := rows[y][x]
	verticalLen := len(rows)
	for {
		if lookY >= verticalLen { // out of bounds
			break
		}
		visibleTrees++
		if rows[lookY][x] >= referenceTree {
			return
		}
		lookY++
	}
	return
}
