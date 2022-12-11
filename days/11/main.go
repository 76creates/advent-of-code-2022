package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	index          int
	items          []int
	testDivisible  int
	opString       []string
	trowsToTrue    int
	trowsToFalse   int
	inspectedItems int
}

var (
	ops = map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}
)

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

func operation(s int, op []string) int {
	var a, b int
	if op[0] == "old" {
		a = s
	} else {
		a, _ = strconv.Atoi(op[0])
	}
	if op[2] == "old" {
		b = s
	} else {
		b, _ = strconv.Atoi(op[2])
	}
	return ops[op[1]](a, b)
}

func parse(file string) map[int]*monkey {
	f, _ := os.Open(file)
	defer f.Close()

	monkeys := make(map[int]*monkey)
	monkeyIndex := -1
	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}

		if !strings.HasPrefix(s.Text(), "Monkey ") {
			panic("n... nani?")
		}
		monkeyIndex++
		monkeys[monkeyIndex] = &monkey{
			index:          monkeyIndex,
			inspectedItems: 0,
		}

		s.Scan()
		startingItems := strings.Split(s.Text()[len("  Starting items: "):], ", ")
		for _, item := range startingItems {
			itemStressLvl, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			monkeys[monkeyIndex].items = append(monkeys[monkeyIndex].items, itemStressLvl)
		}

		s.Scan()
		monkeys[monkeyIndex].opString = strings.Split(s.Text()[len("  Operation: new = "):], " ")

		s.Scan()
		monkeys[monkeyIndex].testDivisible, _ = strconv.Atoi(s.Text()[len("  Test: divisible by "):])

		s.Scan()
		monkeys[monkeyIndex].trowsToTrue, _ = strconv.Atoi(s.Text()[len("    If true: throw to monkey "):])

		s.Scan()
		monkeys[monkeyIndex].trowsToFalse, _ = strconv.Atoi(s.Text()[len("    If false: throw to monkey "):])
	}
	return monkeys
}

func part1(file string) int {
	monkeys := parse(file)

	throwAround(monkeys, 20, func(s int) int {
		return int(math.Floor(float64(s / 3)))
	})

	return calculateMonkeyBusiness(monkeys)
}

func part2(file string) int {
	monkeys := parse(file)

	gcd := 1
	for m := 0; m < len(monkeys); m++ {
		gcd *= monkeys[m].testDivisible
	}

	throwAround(monkeys, 10000, func(s int) int {
		return s % gcd
	})

	return calculateMonkeyBusiness(monkeys)
}

func throwAround(monkeys map[int]*monkey, rounds int, worryLvlFn func(int) int) {
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].items {
				monkeys[m].inspectedItems++
				worryLevel := worryLvlFn(operation(item, monkeys[m].opString))
				if worryLevel%monkeys[m].testDivisible == 0 {
					monkeys[monkeys[m].trowsToTrue].items = append(monkeys[monkeys[m].trowsToTrue].items, worryLevel)
				} else {
					monkeys[monkeys[m].trowsToFalse].items = append(monkeys[monkeys[m].trowsToFalse].items, worryLevel)
				}
			}
			monkeys[m].items = []int{}
		}
	}
}

func calculateMonkeyBusiness(monkeys map[int]*monkey) int {
	var inspected []int
	for m := 0; m < len(monkeys); m++ {
		fmt.Printf("monkey %d inspected %d items\n", m, monkeys[m].inspectedItems)
		inspected = append(inspected, monkeys[m].inspectedItems)
	}
	sort.Ints(inspected)
	monkeyBusiness := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	fmt.Println("monkey business: ", monkeyBusiness)
	return monkeyBusiness
}
