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

	visibleTrees := 0
	for y, row := range rows {
		for x, _ := range row {
			if visibleFromLeft(rows, x, y) || visibleFromRight(rows, x, y) || visibleFromAbove(rows, x, y) || visibleFromBellow(rows, x, y) {
				println("visible")
				visibleTrees++
			}
		}
	}
	println("there are", visibleTrees, "visible trees")

}

func visibleFromLeft(rows [][]int, x, y int) bool {
	lookX := x - 1
	referenceTree := rows[y][x]
	for {
		if lookX < 0 { // out of bounds
			break
		}
		if rows[y][lookX] >= referenceTree {
			return false
		}
		lookX--
	}
	return true
}

func visibleFromRight(rows [][]int, x, y int) bool {
	lookX := x + 1
	referenceTree := rows[y][x]
	rowLen := len(rows[y])
	for {
		if lookX >= rowLen { // out of bounds
			break
		}
		if rows[y][lookX] >= referenceTree {
			return false
		}
		lookX++
	}
	return true
}

func visibleFromAbove(rows [][]int, x, y int) bool {
	lookY := y - 1
	referenceTree := rows[y][x]

	for {
		if lookY < 0 { // out of bounds
			break
		}
		if rows[lookY][x] >= referenceTree {
			return false
		}
		lookY--
	}
	return true
}

func visibleFromBellow(rows [][]int, x, y int) bool {
	lookY := y + 1
	referenceTree := rows[y][x]
	verticalLen := len(rows)
	for {
		if lookY >= verticalLen { // out of bounds
			break
		}
		if rows[lookY][x] >= referenceTree {
			return false
		}
		lookY++
	}
	return true
}
