package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func parse(file string) [][][]int {
	f, _ := os.Open(file)
	defer f.Close()
	rockVectors := [][][]int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		var vector [][]int
		for _, point := range strings.Split(s.Text(), " -> ") {
			x, _ := strconv.Atoi(strings.SplitN(point, ",", 2)[0])
			y, _ := strconv.Atoi(strings.SplitN(point, ",", 2)[1])
			vector = append(vector, []int{x, y})
		}
		rockVectors = append(rockVectors, vector)
	}
	return rockVectors
}

func drawStage(rockVectors [][][]int) ([][]string, int) {
	var scene [][]string
	minX, maxX := findXRange(rockVectors)
	minY, maxY := findYRange(rockVectors)
	for y := minY; y <= maxY; y++ {
		var layer []string
		for x := minX; x <= maxX; x++ {
			layer = append(layer, ".")
		}
		scene = append(scene, layer)
	}
	for _, vectorSet := range rockVectors {
		scene = drawVectorOnStage(vectorSet, scene, minX)
	}
	return scene, minX
}

func drawVectorOnStage(vector [][]int, stage [][]string, normalizeX int) [][]string {
	var storeX, storeY int
	for i, point := range vector {
		currentX := point[0] - normalizeX
		currentY := point[1]
		stage[currentY][currentX] = "#"
		if i > 0 {
			if currentX == storeX {
				sgn := int(math.Copysign(1, float64(storeY-currentY)))
				for distance := 1; distance <= int(math.Abs(float64(storeY-currentY))); distance++ {
					stage[currentY+(distance*sgn)][currentX] = "#"
				}
			} else if currentY == storeY {
				sgn := int(math.Copysign(1, float64(storeX-currentX)))
				for distance := 1; distance <= int(math.Abs(float64(storeX-currentX))); distance++ {
					stage[currentY][currentX+(distance*sgn)] = "#"
				}
			} else {
				panic("Not a straight line")
			}
		}
		storeX = currentX
		storeY = currentY
	}

	return stage
}

func findXRange(rockVectors [][][]int) (int, int) {
	minX := 500
	maxX := 500
	for _, vectorSet := range rockVectors {
		for _, rockVector := range vectorSet {
			if rockVector[0] < minX {
				minX = rockVector[0]
			}
			if rockVector[0] > maxX {
				maxX = rockVector[0]
			}
		}
	}
	return minX, maxX
}

func findYRange(rockVectors [][][]int) (int, int) {
	maxY := 0
	for _, vectorSet := range rockVectors {
		for _, rockVector := range vectorSet {
			if rockVector[1] > maxY {
				maxY = rockVector[1]
			}
		}
	}
	return 0, maxY
}

func dropGrain(stage [][]string, normalizeX int) ([][]string, bool) {
	storeGrainX := 500 - normalizeX
	storeGrainY := 0
	if stage[storeGrainY][storeGrainX] == "~" {
		return stage, true
	}
	stage[storeGrainY][storeGrainX] = "+"
	for {
		if storeGrainY+1 == len(stage) {
			return stage, true
		}
		if stage[storeGrainY+1][storeGrainX] == "." {
			stage[storeGrainY][storeGrainX] = "."
			storeGrainY++
		} else if stage[storeGrainY+1][storeGrainX] == "#" || stage[storeGrainY+1][storeGrainX] == "~" {
			if storeGrainX == 0 {
				return stage, true
			}
			if stage[storeGrainY+1][storeGrainX-1] == "." {
				stage[storeGrainY][storeGrainX] = "."
				storeGrainX--
				storeGrainY++
				continue
			}

			if storeGrainX+1 == len(stage[0]) {
				return stage, true
			}
			if stage[storeGrainY+1][storeGrainX+1] == "." {
				stage[storeGrainY][storeGrainX] = "."
				storeGrainX++
				storeGrainY++
				continue
			}
			break
		}

		stage[storeGrainY][storeGrainX] = "+"
		drawOnScreen(stage)
	}
	stage[storeGrainY][storeGrainX] = "~"
	return stage, false
}

var screenInitialized = false
var drawScreen = false

func drawOnScreen(stage [][]string) {
	if !drawScreen {
		return
	}
	if screenInitialized {
		fmt.Print(fmt.Sprintf("\033[%dA\033[2K", len(stage)+1))
	}
	for ii, sceneLvl := range stage {
		fmt.Println(sceneLvl, ii)
	}
	screenInitialized = true
	fmt.Scanln()
}

func part1(file string) {
	rockVectors := parse(file)
	scene, normalizeX := drawStage(rockVectors)

	oob := false
	grainsDropped := 0
	for {
		scene, oob = dropGrain(scene, normalizeX)
		if oob {
			break
		}
		grainsDropped++
	}
	fmt.Println("dropped grains: ", grainsDropped)
}

func part2(file string) {
	rockVectors := parse(file)
	_, yMax := findYRange(rockVectors)
	xMin, xMax := findXRange(rockVectors)
	rockVectors = append(rockVectors, [][]int{{xMin - 5000, yMax + 2}, {xMax + 5000, yMax + 2}})

	scene, normalizeX := drawStage(rockVectors)
	oob := false
	grainsDropped := 0
	for {
		scene, oob = dropGrain(scene, normalizeX)
		if oob {
			break
		}
		grainsDropped++
	}
	fmt.Println("dropped grains: ", grainsDropped)
}
