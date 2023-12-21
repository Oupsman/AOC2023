package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
	"sync"
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

func solve(puzzle [][]int, startingPoint Position, maxMoves int, part2 bool) int {
	var visited = make(map[int][]Position)

	// move := 0
	visited[0] = append(visited[0], startingPoint)
	for move:=0; move < maxMoves; move++ {
		for _, currentPos := range visited[move] {
			for _, dir := range directions {
				targetPos := Position{currentPos.x + dir.x, currentPos.y + dir.y}
				/* if part2 {
				if targetPos.x < 0 {
				targetPos.x = width-1
				}
				if targetPos.y < 0 {
				targetPos.y = height-1
				}
				if targetPos.x >= width {
				targetPos.x = 0
				}
				if targetPos.y >= height {
				targetPos.y = 0
				}
				} */
				if isMoveValid(puzzle, targetPos) {
					if !contains(targetPos, visited[move+1]) {
						visited[move+1] = append(visited[move+1], targetPos)
					}
				}
			}
		}
	}

	sum := len(visited[len(visited)-1])
	return sum
}

func isMoveValid(puzzle [][]int, targetPos Position) bool {
	realX := targetPos.x % 131
	realY := targetPos.y % 131
	if realX < 0 || realX > len(puzzle[0])-1 || realX <0 || realY < 0 || realY > len(puzzle)-1 || puzzle[realY][realX] == obstacle {
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

func f(x int64) int64 {
	fmt.Println("Target steps:", 65+x*131)
	return 3943 + 13136*x + 3405*int64(math.Pow(float64(x), 2.0))
}

func main() {

	timeStart := time.Now()
	INPUT := "input.txt"
	// NPUT := "test.txt"
	sumPart1 := 0
	sumPart2 := int64(0)

	puzzle, startingPoint := readFileInt(INPUT)

	sumPart1 = solve(puzzle, startingPoint, 64, false)
	// 3858
	wg := sync.WaitGroup{}
	for i := 0; i<10; i++ {
		wg.Add(1)
		go func (i int) {
			value := solve(puzzle, startingPoint, 65+i*131, true)
			fmt.Println("i", i, i*131+65, value)
		}(i)
	}
	wg.Wait()
	sumPart2 = f(202300)
	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
