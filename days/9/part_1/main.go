package main

import (
	"bufio"
	"fmt"
	"log"
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

	visited := make(map[string]struct{})
	hx, hy := 0, 0
	tx, ty := 0, 0
	visited[fmt.Sprintf("%d:%d", tx, ty)] = struct{}{}
	for _, cmd := range commands {
		cmdParts := strings.Split(cmd, " ")
		direction := cmdParts[0]
		distance, _ := strconv.Atoi(cmdParts[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				hy++
			case "D":
				hy--
			case "L":
				hx--
			case "R":
				hx++
			default:
				panic("Unknown direction")
			}
			if !checkTouching(hx, hy, tx, ty) {
				tx, ty = calculateTailMovement(hx, hy, tx, ty)
				visited[fmt.Sprintf("%d:%d", tx, ty)] = struct{}{}
				if !checkTouching(hx, hy, tx, ty) {
					panic("Tail not touching head")
				}
			}
		}
	}
	fmt.Println("visited:", len(visited))
}

func checkTouching(hx, hy, tx, ty int) bool {
	if hx == tx && hy == ty { // same spot
		return true
	} else if (hx == tx+1 || hx == tx-1) && (hy >= ty-1 && hy <= ty+1) { // H left or right of T
		return true
	} else if hx == tx && (hy == ty+1 || hy == ty-1) { // H above or below T
		return true
	}
	return false
}

func calculateTailMovement(hx, hy, tx, ty int) (nextTx int, nextTy int) {
	if hx == tx && hy == ty { // H over T
		return tx, ty
	} else if hx == tx && hy > ty+1 { // H above T
		return tx, ty + 1
	} else if hx == tx && hy < ty-1 { // H below T
		return tx, ty - 1
	} else if hy == ty && hx > tx+1 { // H right of T
		return tx + 1, ty
	} else if hy == ty && hx < tx-1 { // H left of T
		return tx - 1, ty
	} else if (hx > tx+1 && hy > ty) || (hx > tx && hy > ty+1) { // H up right of T, but not touching
		return tx + 1, ty + 1
	} else if (hx < tx-1 && hy > ty) || (hx < tx && hy > ty+1) { // H up left of T, but not touching
		return tx - 1, ty + 1
	} else if (hx > tx+1 && hy < ty) || (hx > tx && hy < ty-1) { // H down right of T, but not touching
		return tx + 1, ty - 1
	} else if (hx < tx-1 && hy < ty) || (hx < tx && hy < ty-1) { // H down left of T, but not touching
		return tx - 1, ty - 1
	}
	log.Fatalf("Unknown movement: hx=%d, hy=%d, tx=%d, ty=%d\n", hx, hy, tx, ty)
	return
}
