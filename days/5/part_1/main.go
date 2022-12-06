package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()

	commandLines := false
	var cratesRaw []string
	var commands []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			commandLines = true
		} else if commandLines {
			commands = append(commands, line)
		} else {
			cratesRaw = append(cratesRaw, line)
		}
	}

	var cratesWidth int
	cratesWidth = len(strings.Split(cratesRaw[len(cratesRaw)-1], "   "))
	fmt.Println("Width:", cratesWidth)

	cratesRaw = cratesRaw[:len(cratesRaw)-1]
	var stackCols [][]string
	for i := 0; i < cratesWidth; i++ {
		stackCols = append(stackCols, []string{})
	}
	chunkSize := 4
	for _, crateRaw := range cratesRaw {
		for chunkNum := 0; (chunkNum*chunkSize)+chunkSize < len(crateRaw)+2; chunkNum++ {
			chunkStart := chunkNum * chunkSize
			if crateRaw[chunkStart:chunkStart+3] == "   " {
				continue
			} else {
				stackCols[chunkNum] = append(stackCols[chunkNum], crateRaw[chunkStart+1:chunkStart+2])
			}
		}
	}

	for _, command := range commands {
		// move 1 from 2 to 1
		commandSlice := strings.Split(command, " ")
		moveNum, _ := strconv.Atoi(commandSlice[1])
		fromCol, _ := strconv.Atoi(commandSlice[3])
		fromCol--
		toCol, _ := strconv.Atoi(commandSlice[5])
		toCol--
		for i := 0; i < moveNum; i++ {
			move := stackCols[fromCol][0]
			stackCols[toCol] = append([]string{move}, stackCols[toCol]...)
			stackCols[fromCol] = stackCols[fromCol][1:]
		}
	}

	for _, stackCol := range stackCols {
		fmt.Print(stackCol[0])
	}

	//println("less elfs: ", lessElfs)
}
