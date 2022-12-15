package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
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
	part1("sample", 10)
	fmt.Println("\n________ SAMPLE PART TWO ________")
	part2("sample", 20)
	fmt.Println("\n--------------------------------------------\n")

	fmt.Println("________ PART ONE ________")
	part1("input", 2000000)
	fmt.Println("\n________ PART TWO ________")
	part2("input", 4000000)
}

func parse(file string) (sensors [][]int, beacons [][]int) {
	lineXp, _ := regexp.Compile(`Sensor at x=(\d+), y=(\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	f, _ := os.Open(file)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := lineXp.FindAllStringSubmatch(s.Text(), -1)
		sensorX, _ := strconv.Atoi(m[0][1])
		sensorY, _ := strconv.Atoi(m[0][2])
		beaconX, _ := strconv.Atoi(m[0][3])
		beaconY, _ := strconv.Atoi(m[0][4])
		sensors = append(sensors, []int{sensorX, sensorY})
		beacons = append(beacons, []int{beaconX, beaconY})
	}
	return sensors, beacons
}

func getDistance2D(a, b []int) int {
	return getDistance(a[0], b[0]) + getDistance(a[1], b[1])
}
func getDistance(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func getXSpan(referenceY int, a, b []int) (present bool, min, max int) {
	distance := getDistance2D(a, b)
	distanceY := getDistance(a[1], referenceY)
	if distance-distanceY < 0 {
		return false, 0, 0
	}
	return true, a[0] - (distance - distanceY), a[0] + (distance - distanceY)
}

func findFilledCoordinates(spans [][]int) map[int]struct{} {
	filled := make(map[int]struct{})

	sort.SliceStable(spans, func(i, j int) bool {
		if spans[i][0] == spans[j][0] {
			return spans[i][1] < spans[j][1]
		}
		return spans[i][0] < spans[j][0]
	})

	for i := 0; i < len(spans); i++ {
		for j := spans[i][0]; j <= spans[i][1]; j++ {
			if _, ok := filled[j]; !ok {
				filled[j] = struct{}{}
			}
		}
	}

	return filled
}

func findEmptyCoordinate(spans [][]int) int {
	empty := -1
	sort.SliceStable(spans, func(i, j int) bool {
		if spans[i][0] == spans[j][0] {
			return spans[i][1] < spans[j][1]
		}
		return spans[i][0] < spans[j][0]
	})
	cursor := 0
	for i := 0; i < len(spans); i++ {
		if cursor > spans[i][0] && cursor > spans[i][1] {
			continue
		} else if cursor >= spans[i][0] && cursor <= spans[i][1] {
			cursor = spans[i][1]
			continue
		} else if cursor < spans[i][0] {
			if cursor+1 == spans[i][0] {
				cursor = spans[i][1]
				continue
			} else if cursor+2 < spans[i][0] {
				panic("too many empty")
			}
			if empty != -1 {
				panic("too many empty")
			}
			empty = cursor + 1
			cursor = spans[i][1]
		}
	}
	return empty
}

func part1(file string, row int) {
	sensors, beacons := parse(file)

	var spans [][]int
	for i := 0; i < len(sensors); i++ {
		if present, min, max := getXSpan(row, sensors[i], beacons[i]); present {
			spans = append(spans, []int{min, max})
		}
	}
	filled := len(findFilledCoordinates(spans))

	uniqueBeacon := make(map[string]struct{})
	for i := 0; i < len(beacons); i++ {
		if beacons[i][1] == row {
			uniqueBeacon[fmt.Sprintf("%d:%d", beacons[i][0], beacons[i][1])] = struct{}{}
		}
	}
	fmt.Println("Filled", filled-len(uniqueBeacon))
}

func part2(file string, depth int) {
	sensors, beacons := parse(file)

	for d := 0; d < depth; d++ {
		var spans [][]int
		for i := 0; i < len(sensors); i++ {
			if present, min, max := getXSpan(d, sensors[i], beacons[i]); present {
				spans = append(spans, []int{min, max})
			}
		}
		emptyCoordinate := findEmptyCoordinate(spans)
		if emptyCoordinate < 0 {
			continue
		}
		fmt.Println((emptyCoordinate * 4000000) + d)

	}

	parse(file)
}
