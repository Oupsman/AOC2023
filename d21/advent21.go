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

func solve(puzzle [][]int, startingPoint Position, maxMoves int, part2 bool, a *[3]int64) int64 {
	var visited = make(map[int][]Position)

	// move := 0
	visited[0] = append(visited[0], startingPoint)
	found := 0
	prevLen := int64(0)
	for move := 0; move < maxMoves; move++ {
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
		if part2 {
			if (move % len(puzzle)) == (maxMoves % len(puzzle)) {
				fmt.Println("Move", move, len(visited[move]), "prevLen", int64(len(visited[move])) - prevLen)
				prevLen = int64(len(visited[move]))
				(*a)[found] = prevLen
				found++
			}
			if found == 3 {
				break
			}
		}
	}

	sum := int64(len(visited[len(visited)-1]))
	return sum
}

func isMoveValid(puzzle [][]int, targetPos Position) bool {
	realX := ((targetPos.x % len(puzzle))+ len(puzzle)) % len(puzzle)
	realY := ((targetPos.y % len(puzzle))+ len(puzzle)) % len(puzzle)

	if realX < 0 || realX > len(puzzle[0])-1 || realX < 0 || realY < 0 || realY > len(puzzle)-1 || puzzle[realY][realX] == obstacle {
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

func f(x int64, a [3]int64) int64 {
	b0 := a[0]
	b1 := a[1]-a[0]
	b2 := a[2]-a[1]
	return b0 + b1*x + (x*(x-1)/2)*(b2-b1)

}

func main() {

	timeStart := time.Now()
	INPUT := "input.txt"
	// INPUT := "test.txt"
	sumPart1 := int64(0)
	sumPart2 := int64(0)

	puzzle, startingPoint := readFileInt(INPUT)

	sumPart1 = solve(puzzle, startingPoint, 64, false, nil)
	var a [3]int64
	solve(puzzle, startingPoint, 26501365, true, &a)
	sumPart2 = f(int64(26501365/len(puzzle)), a)
	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
