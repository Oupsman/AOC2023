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

	E = 0
	S = 1
	W = 2
	N = 3
)

type Position struct {
	x, y int
}

type Move struct {
	pos   Position
	steps int
}

var directions = []Position{
	E: {1, 0},  // E
	S: {0, 1},  // S
	W: {-1, 0}, // W
	N: {0, -1}, //N
}

func contains(val Position, arr []Position) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func solve(puzzle [][]int, startingPoint Position) int {

	var visited = make(map[int][]Position)
	maxMoves := 64
	// move := 0
	visited[0] = append(visited[0], startingPoint)
	for move:=0; move < maxMoves; move++ {
		fmt.Println("Move", move)
		for _, currentPos := range visited[move] {
			for _, dir := range directions {
				targetPos := Position{currentPos.x + dir.x, currentPos.y + dir.y}
				if isMoveValid(puzzle, targetPos) {
					if !contains(targetPos, visited[move+1]) {
						visited[move+1] = append(visited[move+1], targetPos)
					}
				}
			}
		}
	}
	
	fmt.Println(len(visited))
	sum := len(visited[len(visited)-1])
	return sum
}

func isMoveValid(puzzle [][]int, targetPos Position) bool {
	width := len(puzzle[0])
	height := len(puzzle)
	if targetPos.x < 0 || targetPos.x > width-1 || targetPos.y < 0 || targetPos.y > height-1 || puzzle[targetPos.y][targetPos.x] == obstacle {
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
	x := 0
	for scanner.Scan() {
		var lineI []int
		lineT := scanner.Text()
		for y, char := range lineT {
			var num int
			switch string(char) {
			case ".":
				num = free
				break
			case "#":
				num = obstacle
				break
			case "S":
				num = start
				break
			}
			if num == start {
				startPoint = Position{x, y}
			}
			lineI = append(lineI, num)
		}
		puzzle = append(puzzle, lineI)
		x++
	}
	return puzzle, startPoint
}

func main() {

	timeStart := time.Now()
	INPUT := "input.txt"
	// NPUT := "test.txt"
	sumPart1 := 0

	puzzle, startingPoint := readFileInt(INPUT)

	sumPart1 = solve(puzzle, startingPoint)
	// 3858
	fmt.Println("Part1:", sumPart1)
	// fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
