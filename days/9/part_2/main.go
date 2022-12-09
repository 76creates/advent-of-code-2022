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
	mrx := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	mry := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	visited[fmt.Sprintf("%d:%d", 0, 0)] = struct{}{}
	for _, cmd := range commands {
		cmdParts := strings.Split(cmd, " ")
		direction := cmdParts[0]
		distance, _ := strconv.Atoi(cmdParts[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				mry[0]++
			case "D":
				mry[0]--
			case "L":
				mrx[0]--
			case "R":
				mrx[0]++
			default:
				panic("Unknown direction")
			}
			for i := 1; i < 10; i++ {
				if !checkTouching(mrx[i-1], mry[i-1], mrx[i], mry[i]) {
					mrx[i], mry[i] = calculateTailMovement(mrx[i-1], mry[i-1], mrx[i], mry[i])
					if !checkTouching(mrx[i-1], mry[i-1], mrx[i], mry[i]) {
						panic("Tail not touching head")
					}
				}
			}
			visited[fmt.Sprintf("%d:%d", mrx[9], mry[9])] = struct{}{}
		}
	}
	fmt.Println("9 visited:", len(visited))
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
