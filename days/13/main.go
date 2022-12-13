package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func parse(file string) [][]string {
	f, _ := os.Open(file)
	defer f.Close()

	s := bufio.NewScanner(f)
	var packetPairs [][]string
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		packetPair1 := s.Text()
		s.Scan()
		packetPair2 := s.Text()
		packetPairs = append(packetPairs, []string{packetPair1, packetPair2})
	}
	return packetPairs
}

func getPacketStructure(input string) []any {
	var items []any
	for {
		if len(input) == 0 {
			break
		}
		var item any
		if input == "[]" {
			item = []any{"nil"}
			input = ""
		} else if matchedBracket, err := regexp.Match(`^\[(.*?)\]$`, []byte(input)); err != nil {
			panic(err)
		} else if matchedBracket {
			item = getPacketStructure(input[1 : len(input)-1])
		} else if matchedNum, err := regexp.Match(`^[0-9],.*`, []byte(input)); err != nil {
			panic(err)
		} else if matchedNum {
			item, err = strconv.Atoi(string(input[0]))
			if err != nil {
				panic(err)
			}
			input = input[2:]
		} else {
			println(input)
			panic("Invalid input")
		}

		items = append(items, item)
	}
	return items
}

func part1(file string) {
	parse(file)
}

func part2(file string) {
	parse(file)
}
