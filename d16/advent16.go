package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Coords struct {
	x int
	y int
}

const (
	E = iota
	S
	W
	N
)

type Dir struct {
	x int
	y int
}

var directions = []Dir{
	E: {1, 0},  // E
	S: {0, 1},  // S
	W: {-1, 0}, // W
	N: {0, -1}, //N
}

var MIRROR_A = []int{
	E: N,
	W: S,
	S: W,
	N: E,
}

var MIRROR_B = []int{
	E: S,
	S: E,
	W: N,
	N: W,
}

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func simulateBeam(puzzle []string, coords Coords, dir int, energized *map[Coords]int, cache *map[string]int) {
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	beamInGrid := true
	currentCoords := coords
	currentDir := dir

	key := fmt.Sprintf("%d,%d,%d,%d", currentCoords.x, currentCoords.y, currentDir)
	if _, ok := (*cache)[key]; ok {
		return
	}
	if (currentCoords.x < 0) || (currentCoords.x > puzzleW) || (currentCoords.y < 0) || (currentCoords.y > puzzleH) {
		return
	}
	for beamInGrid {

		key := fmt.Sprintf("%d,%d,%d,%d", currentCoords.x, currentCoords.y, currentDir)
		if _, ok := (*cache)[key]; ok {
			return
		} else {
			(*cache)[key] = 1
		}
		if (*energized)[currentCoords] == 0 {
			(*energized)[currentCoords] = 1
		}

		switch puzzle[currentCoords.y][currentCoords.x] {
		case byte('/'):
			currentDir = MIRROR_A[currentDir]
			break
		case byte('\\'):
			currentDir = MIRROR_B[currentDir]
			break
		case byte('-'):
			if currentDir == N || currentDir == S {

				simulateBeam(puzzle, currentCoords, E, energized, cache)
				simulateBeam(puzzle, currentCoords, W, energized, cache)
				return
			}
			break
		case byte('|'):
			if currentDir == E || currentDir == W {

				simulateBeam(puzzle, currentCoords, N, energized, cache)
				simulateBeam(puzzle, currentCoords, S, energized, cache)
				return
			}
			break
		}
		currentCoords.x += directions[currentDir].x
		currentCoords.y += directions[currentDir].y
		if (currentCoords.x < 0) || (currentCoords.x >= puzzleW) || (currentCoords.y < 0) || (currentCoords.y >= puzzleH) {
			return
		}
	}

}

func resetEnergy (energized *map[Coords]int) {
	for key, _ := range *energized {
		(*energized)[key] = 0
	}
}

func sumEnergy (energized *map[Coords]int) int {
	sum := 0
	for key, _ := range *energized {
		sum += (*energized)[key]
	}
	return sum
}

func main() {
	timeStart := time.Now()
	INPUT := "input16.txt"
	//INPUT := "test16.txt"
	sumPart1 := 0

	puzzle := readFile(INPUT)

	var energizedTiles = make(map[Coords]int)
	var cache = make(map[string]int) // cache for the beam

	var initialCoords Coords
	initialCoords.x = 0
	initialCoords.y = 0
	simulateBeam(puzzle, initialCoords, E, &energizedTiles, &cache)

	sumPart1 = sumEnergy(&energizedTiles)

	fmt.Println("Part1:", sumPart1)
	
	maxEnergy := 0
	
	for row := 0; row < len(puzzle); row ++ {
		// check for every entrypoint
		initialCoords.x = 0
		initialCoords.y = row
		cache = make(map[string]int) // reset cache
		resetEnergy(&energizedTiles)
		simulateBeam(puzzle, initialCoords, E, &energizedTiles, &cache)
		
		if sumEnergy(&energizedTiles) > maxEnergy {
			maxEnergy = sumEnergy(&energizedTiles)
		}
		
		initialCoords.x = len(puzzle[0])-1
		initialCoords.y = row

		cache = make(map[string]int) // reset cache
		resetEnergy(&energizedTiles)
		simulateBeam(puzzle, initialCoords, W, &energizedTiles, &cache)
		
		if sumEnergy(&energizedTiles) > maxEnergy {
			maxEnergy = sumEnergy(&energizedTiles)
		}
	}
	
	for col := 0; col < len(puzzle[0]); col++ {
		initialCoords.x = col
		initialCoords.y = 0
		cache = make(map[string]int) // reset cache
		resetEnergy(&energizedTiles)
		simulateBeam(puzzle, initialCoords, S, &energizedTiles, &cache)

		if sumEnergy(&energizedTiles) > maxEnergy {
			maxEnergy = sumEnergy(&energizedTiles)
		}

		initialCoords.x = col
		initialCoords.y = len(puzzle[0])-1
		cache = make(map[string]int) // reset cache
		resetEnergy(&energizedTiles)
		simulateBeam(puzzle, initialCoords, N, &energizedTiles, &cache)

		if sumEnergy(&energizedTiles) > maxEnergy {
			maxEnergy = sumEnergy(&energizedTiles)
		}
	}
	
	fmt.Println("Part2:", maxEnergy)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}
