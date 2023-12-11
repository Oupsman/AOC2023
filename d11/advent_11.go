package main

import (
	"bufio"
	"math"
	"os"
	"strings"

	"fmt"
)

type COORDS []int


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

func weight (universe []string) ([]int, []int) {
	rows := make([]int, len(universe))
	columns := make([]int, len(universe[0]))
	
	for i:=0; i<len(universe); i++ {
		if strings.Contains(universe[i], "#") {
			rows[i] = 1
		} else {
			rows[i] = 2
		}
	}
	
	for j:=0; j<len(universe[0]); j++ {
		columns[j] = 2
		for i:=0; i<len(universe); i++ {
			if string(universe[i][j]) == "#" {
				columns[j] = 1				
				break
			}
		}
	}
	
	return rows, columns
}

func getAllGalaxies (universe []string) []COORDS {
	var galaxies []COORDS

	for i:=0; i<len(universe); i++ {
		for j:=0; j<len(universe[i]); j++ {
			var galaxy COORDS
			if string(universe[i][j]) == "#" {
				galaxy = append(galaxy, i)
				galaxy = append(galaxy, j)
				galaxies = append(galaxies, galaxy)
			}
		}
	}
	return galaxies
}

func sumSubpart(array []int, start int, end int, part2 bool) int {
	var startLoop, endLoop int
	res := 0
	// fmt.Println("Start, end:", start, end)
	if start >= end {
		startLoop = end
		endLoop = start
	} else {
		startLoop = start
		endLoop = end
	}
	for i := startLoop; i < endLoop; i++ {
	//	fmt.Println(array, start, end, array[i], res)
		if part2 && array[i] == 2 {
			res += 1000000
		} else {
			res += array[i]			
		}

	}
	return res
}

func main() {
	INPUT := "input_11.txt"
	// INPUT := "test_11.txt"
	var galaxies []COORDS
	
	rawUniverse := readFile(INPUT)
	
	// universeH := len(rawUniverse)
	// universeW := len(rawUniverse[0])
	
	rows, columns := weight(rawUniverse)
	// fmt.Println("Rows:", rows, "Columns: ", columns)
	galaxies = getAllGalaxies(rawUniverse)
	// fmt.Println("Galaxies :", galaxies)
	sum_part1 := 0
	pairCounter := 0
	for i, galaxy := range galaxies {
		for j, targetGalaxy := range galaxies {
			if i != j && j < len(galaxies){
				pairCounter++	
	//			fmt.Println("Galaxies involved: ", galaxy, targetGalaxy)
				numCols := math.Abs(float64(sumSubpart(columns, targetGalaxy[1], galaxy[1], false)))
				numRows := math.Abs(float64(sumSubpart(rows, targetGalaxy[0], galaxy[0], false)))
	//			fmt.Println("Differences:", numRows, numCols)
				shortestPath := numRows + numCols
	//			fmt.Println("ShortestPath: ", shortestPath)
				sum_part1 += int(shortestPath)
			}
		}
	}
	fmt.Println("Part 1:", sum_part1/2, pairCounter)
	sum_part2 := 0
	for i, galaxy := range galaxies {
		for j, targetGalaxy := range galaxies {
			if i != j && j < len(galaxies){
				pairCounter++	
				//			fmt.Println("Galaxies involved: ", galaxy, targetGalaxy)
				numCols := math.Abs(float64(sumSubpart(columns, targetGalaxy[1], galaxy[1], true)))
				numRows := math.Abs(float64(sumSubpart(rows, targetGalaxy[0], galaxy[0], true)))
				//			fmt.Println("Differences:", numRows, numCols)
				shortestPath := numRows + numCols
				//			fmt.Println("ShortestPath: ", shortestPath)
				sum_part2 += int(shortestPath)
			}
		}
	}
	fmt.Println("Part 2:", sum_part2/2)
	
}
