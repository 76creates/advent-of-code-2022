package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type terrainSpot struct {
	id        int
	height    int
	distance  int
	proximity []*terrainSpot
}

func main() {
	fmt.Println("________ SAMPLE PART ONE ________")
	part1("sample")
	fmt.Println("\n________ SAMPLE PART TWO ________")
	part2("sample")
	fmt.Println("\n--------------------------------------------\n")

	fmt.Println("________ PART ONE ________")
	part1("input")
	fmt.Println("\n________ PART TWO ________")
	part2("input")
}

func parse(file string) [][]any {
	f, _ := os.Open(file)
	defer f.Close()

	s := bufio.NewScanner(f)
	var packetPairs [][]any
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		var pair1, pair2 any
		if err := json.Unmarshal([]byte(s.Text()), &pair1); err != nil {
			panic(err)
		}
		s.Scan()
		if err := json.Unmarshal([]byte(s.Text()), &pair2); err != nil {
			panic(err)
		}
		packetPairs = append(packetPairs, []any{pair1, pair2})
	}
	return packetPairs
}

func isSame(pair1, pair2 []any) int {
	for i := 0; i < len(pair1); i++ {
		if i >= len(pair2) {
			return 0
		}
		if val1, ok := pair1[i].(float64); ok {
			if val2, ok := pair2[i].(float64); ok {
				if val1 == val2 {
					continue
				} else if val1 < val2 {
					return 1
				}
				return 0
			}
		}
		_, okArr1 := pair1[i].([]any)
		_, okArr2 := pair2[i].([]any)
		if okArr1 || okArr2 {
			if !okArr1 {
				if s := isSame([]any{pair1[i]}, pair2[i].([]any)); s < 2 {
					return s
				}
			} else if !okArr2 {
				if s := isSame(pair1[i].([]any), []any{pair2[i]}); s < 2 {
					return s
				}
			} else {
				if s := isSame(pair1[i].([]any), pair2[i].([]any)); s < 2 {
					return s
				}
			}
		}
	}
	if len(pair1) == len(pair2) {
		return 2
	}
	return 1
}

func findNum(arr any) int {
	if val, ok := arr.(float64); ok {
		return int(val)
	} else if arr2, ok := arr.([]any); ok {
		if len(arr2) == 0 {
			return 0
		}
		return findNum(arr2[0])
	} else {
		panic("what is this sorcery")
	}
}

func findLen(arr []any) int {
	len := len(arr)
	for _, v := range arr {
		if _, ok := v.([]any); ok {
			len += findLen(v.([]any))
		} else if _, ok := v.(float64); ok {
		} else {
			panic("what is this sorcery")
		}
	}
	return len
}

func part1(file string) {
	content := parse(file)
	sum := 0
	for i, pair := range content {
		if isSame(pair[0].([]any), pair[1].([]any)) == 1 {
			sum += i + 1
		}
	}
	println(sum)
}

func part2(file string) {
	// let nobody tell you how you solve this problem <3

	content := parse(file)
	var ints []int
	ints = append(ints, 202)
	ints = append(ints, 602)
	for _, pair := range content {
		for _, p := range pair {
			ints = append(ints, findNum(p)*100+findLen(p.([]any)))
		}
	}

	sort.Ints(ints)
	sum := 1
	for index, val := range ints {
		if val == 202 || val == 602 {
			sum *= index + 1
		}
	}

	fmt.Println(sum)
}
