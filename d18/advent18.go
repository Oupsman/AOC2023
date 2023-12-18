package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"regexp"
	"strconv"
	"math"
)

const (
	L = iota
	R
	U
	D
)

type Dir struct {
	x int64
	y int64
}

var directions = []Dir{
	R: {1, 0},  // E
	D: {0, 1},  // S
	L: {-1, 0}, // W
	U: {0, -1}, //N
}

type Point struct {
	x, y int64
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

func shoeLace(points []Point) int64 {
	var area int64 = 0

	j := len(points) - 1

	for i := 0; i < len(points); i++ {
		area += (points[j].x + points[i].x) * (points[j].y - points[i].y)
		j = i
	}

	return int64(math.Abs(float64(area) / 2.0))
}

func main() {
	timeStart := time.Now()
	// INPUT := "test.txt"
	INPUT := "input.txt"
	var poly []Point
	var poly2 []Point
	var perimeter int64

	decomposeLine := regexp.MustCompile(`([LRUD]) ([0-9]+) \(([#a-f0-9]{7})\)`)

	perimeter = 0
	pointer := Point{0, 0}
	poly = append(poly, pointer)
	fileContent := readFile(INPUT)
	for _, line := range fileContent {
		match := decomposeLine.FindStringSubmatch(line)
		move, _ := strconv.Atoi(match[2])
		// Substrings start at index 1
		switch match[1] {
		case "L":
			pointer.x += directions[L].x * int64(move)
			break
		case "R":
			pointer.x += directions[R].x * int64(move)
			break
		case "U":
			pointer.y += directions[U].y * int64(move)
			break
		case "D":
			pointer.y += directions[D].y * int64(move)
			break
		}
		perimeter += int64(move)
		poly = append(poly, pointer)
	}

	sumPart1 := shoeLace(poly) + perimeter/2 + 1

	pointer = Point{0, 0}
	perimeter = 0
	
	poly2 = append(poly2, pointer)
	for _, line := range fileContent {
		match := decomposeLine.FindStringSubmatch(line)
		instruction := match[3]
		
		direction,_ := strconv.Atoi(string(instruction[6]))
		move,_ := strconv.ParseInt(instruction[1:6], 16, 64)
		fmt.Println("Move:", move, direction)
		// Substrings start at index 1
		switch direction {
		case 2:
			pointer.x += directions[L].x * move
			break
		case 0:
			pointer.x += directions[R].x * move
			break
		case 3:
			pointer.y += directions[U].y * move
			break
		case 1:
			pointer.y += directions[D].y * move
			break
		}
		perimeter += move
		poly2 = append(poly2, pointer)
	}

	sumPart2 := shoeLace(poly2) + perimeter/2 + 1
	
	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
