package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dirs = make(map[string]map[string]int)
var currentDir []string

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var commands []string
	for s.Scan() {
		commands = append(commands, s.Text())
	}

	var xVal, cycle, summedSignalStrenght int
	xVal = 1
	for _, command := range commands {
		var f func()
		var execCycles int
		if command == "noop" {
			execCycles = 1
			f = func() {}
		} else {
			execCycles = 2
			addXVal, err := strconv.Atoi(strings.Split(command, " ")[1])
			if err != nil {
				panic(err)
			}
			f = func() {
				xVal += addXVal
			}
		}
		for i := 0; i < execCycles; i++ {
			cycle++
			if cycle == 20 || (cycle-20)%40 == 0 {
				fmt.Println("cycle: ", cycle, ",xVal: ", xVal, ",signalStrength: ", xVal*cycle)
				summedSignalStrenght += xVal * cycle
			}
		}
		f()
	}
	fmt.Println("summedSignalStrenght: ", summedSignalStrenght)
}
