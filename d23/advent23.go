package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	free = iota
	obstacle
	start
	slope

	E = 0
	S = 1
	W = 2
	N = 3
)

type Position struct {
	x, y int
}

var directions = []Position{
	E: {1, 0},  // E
	S: {0, 1},  // S
	W: {-1, 0}, // W
	N: {0, -1}, //N
}

func isMoveValid(puzzle [][]int, targetPos Position, visited map[Position]bool) bool {
	realX := ((targetPos.x % len(puzzle)) + len(puzzle)) % len(puzzle)
	realY := ((targetPos.y % len(puzzle)) + len(puzzle)) % len(puzzle)

	if realX < 0 || realX > len(puzzle[0])-1 || realX < 0 || realY < 0 || realY > len(puzzle)-1 || puzzle[realY][realX] == obstacle || visited[targetPos] {
		return false
	}
	return true
}

func readFileInt(fname string) ([][]int, Position) {
	var puzzle [][]int
	var startPoint Position
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		var lineI []int
		lineT := scanner.Text()
		for x, char := range lineT {
			var num int
			switch string(char) {
			case ".":
				if y == 0 {
					num = start
					break
				}

				num = free
				break
			case "#":
				num = obstacle
				break
			case ">", "<", "v", "^":
				num = slope
				break
			}
			if num == start {
				startPoint = Position{x, y}
			}
			lineI = append(lineI, num)
		}
		puzzle = append(puzzle, lineI)
		y++
	}
	return puzzle, startPoint
}

func solve(puzzle [][]int, startingPoint Position, part2 bool) int {
	var visited = make(map[Position]bool)
	var queue = []Position{}
	queue = append(queue, startingPoint)
	visited[startingPoint] = true
	height := len(puzzle) - 1
	for len(queue) > 0 {
		currentPos := queue[0]
		visited[currentPos] = true
		queue = queue[1:]
		if currentPos.y == height {
			if puzzle[currentPos.y][currentPos.x] == free {
				// We found the exit, pathfinding is done !
				break
			}
		}
		for _, dir := range directions {
			var nextPos Position
			nextPos.x = currentPos.x + dir.x
			nextPos.y = currentPos.y + dir.y
			if isMoveValid(puzzle, nextPos, visited) {
				queue = append(queue, nextPos)
			}

		}
	}
	return len(visited)
}

func main() {
	timeStart := time.Now()
	// INPUT := "test.txt"
	INPUT := "input.txt"
	fileContent, startingPoint := readFileInt(INPUT)
	fmt.Println(fileContent, startingPoint)
	sumPart1 := solve(fileContent, startingPoint, false)

	// sumPart2 := solve(fileContent, true)

	fmt.Println("Part1:", sumPart1)
	// fmt.Println("Part2:", sumPart2)

	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
