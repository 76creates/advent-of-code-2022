package main

import (
	"bufio"
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

// abcdefghijklmnopqrstuvwxyz

func parse(file string) [][]int {
	f, _ := os.Open(file)
	defer f.Close()

	s := bufio.NewScanner(f)
	var terrain [][]int
	for s.Scan() {
		var row []int
		for _, c := range s.Text() {
			row = append(row, int(c))
		}
		terrain = append(terrain, row)
	}
	return terrain
}

func createGraph(terrain [][]int) (spots map[int]*terrainSpot, start, end int) {
	vtxMap := make(map[int]*terrainSpot)
	// e = 69 s = 83
	for r, row := range terrain {
		for i, height := range row {
			id := r*len(row) + i
			h := height
			if height == 83 {
				start = id
				h = int('a')
			}
			if height == 69 {
				end = id
				h = int('z')
			}
			vtxMap[id] = &terrainSpot{
				id:       id,
				height:   h,
				distance: 0,
			}
		}
	}

	terrainLn := len(terrain[0])
	// add proximity
	for i, v := range vtxMap {
		proxy := []*terrainSpot{}
		if i-1 >= 0 {
			proxy = append(proxy, vtxMap[i-1])
		}
		if i+1 < len(vtxMap) {
			proxy = append(proxy, vtxMap[i+1])
		}
		if i+terrainLn < len(vtxMap) {
			proxy = append(proxy, vtxMap[i+terrainLn])
		}
		if i-terrainLn >= 0 {
			proxy = append(proxy, vtxMap[i-terrainLn])
		}

		for _, spot := range proxy {
			if v.height+1 >= spot.height {
				v.proximity = append(v.proximity, spot)
			}
		}
	}
	fmt.Printf("vtxMap: %+v", vtxMap[end])
	return vtxMap, start, end
}

func shortestPath(spots map[int]*terrainSpot, start, end int) int {
	fmt.Println("finding shortest path from", start, "to", end)
	visitedSpots := make(map[int]struct{})
	distanceMap := make(map[int]int)

	for spotId := range spots {
		distanceMap[spotId] = 999999
	}

	visitedSpots[start] = struct{}{}
	distanceMap[start] = 0
	toVisit := []int{start}

	for {
		if len(toVisit) == 0 {
			break
		}
		currentSpot := toVisit[0]
		if len(toVisit) > 1 {
			toVisit = toVisit[1:]
		} else {
			toVisit = []int{}
		}

		for _, spot := range spots[currentSpot].proximity {
			if _, ok := visitedSpots[spot.id]; !ok {
				visitedSpots[spot.id] = struct{}{}
				toVisit = append(toVisit, spot.id)
				if distanceMap[currentSpot]+1 < distanceMap[spot.id] {
					distanceMap[spot.id] = distanceMap[currentSpot] + 1
				}
				toVisit = append(toVisit, spot.id)
			}
		}
	}
	return distanceMap[end]
}

func part1(file string) {
	terrain := parse(file)
	fmt.Println(shortestPath(createGraph(terrain)))
}

func part2(file string) {
	terrain := parse(file)
	paths := []int{}
	spots, _, end := createGraph(terrain)
	for _, spot := range spots {
		if spot.height == int('a') {
			paths = append(paths, shortestPath(spots, spot.id, end))
		}
	}
	sort.Ints(paths)
	fmt.Println(paths[0])
}
