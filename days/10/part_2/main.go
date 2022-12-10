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

	var screen []string
	var spritePos, cycle int
	spritePos = 1
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
				spritePos += addXVal
			}
		}
		for i := 0; i < execCycles; i++ {
			if (spritePos%40)-1 <= cycle%40 && (spritePos%40)+1 >= cycle%40 {
				screen = append(screen, "#")
			} else {
				screen = append(screen, " ")
			}
			cycle++
		}
		f()
	}
	for i, pixel := range screen {
		fmt.Print(pixel)
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}
