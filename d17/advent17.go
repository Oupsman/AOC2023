package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

var left = []int{
	E: N,
	W: S,
	S: W,
	N: E,
}

var right = []int{
	E: S,
	S: E,
	W: N,
	N: W,
}

// GridCell represents each cell in the grid
type GridCell struct {
	X, Y   int     // Coordinates
	Weight int     // Heat loss value
}

// Node represents a state in the search
type Node struct {
	Position  GridCell
	Cost      int     // Total cost so far
	Heuristic int     // Estimated cost to destination
	Parent    *Node   // To track the path
	Index     int     // Index in the heap
	Direction int     // Current direction of movement
	Moves     int     // Count of moves in the current direction
}

// PriorityQueue implements heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].Cost+pq[i].Heuristic < pq[j].Cost+pq[j].Heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.Index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.Index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

func indexOf(arr []Dir, candidate Dir) int {
	for index, c := range arr {
		if c == candidate {
			return index
		}
	}
	return -1
}


// AStarSearch performs the A* search algorithm
func AStarSearch(grid [][]int, start, end GridCell) ([]GridCell, error) {
	// Initialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Define start and end nodes
	startNode := &Node{Position: start, Cost: 0, Heuristic: heuristic(start, end), Moves: 0, Direction: E}
	heap.Push(&pq, startNode)

	// Visited nodes map
	visited := make(map[GridCell]bool)

	// Directions (up, right, down, left)
	// directions := []GridCell{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	// currentDir := E
	// Search loop
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		// currentDir = current.Direction
		// Check if we reached the end
		if current.Position == end {
			return reconstructPath(current), nil
		}

		visited[current.Position] = true

		// availableDirections := []Dir{directions[left[currentDir]], directions[right[currentDir]], directions[currentDir]}
		fmt.Println("Position:", current.Position)
		// Explore neighbors
		for i, dir := range directions {
			nextPos := GridCell{X: current.Position.X + dir.x, Y: current.Position.Y + dir.y}

			// Check if valid move
			if isValidMove(current, nextPos, grid, i) && !visited[nextPos] {
				nextCost := current.Cost + grid[nextPos.X][nextPos.Y]
				nextHeuristic := heuristic(nextPos, end)
				fmt.Println("\tHeuristic: ", nextHeuristic, nextPos)
				fmt.Println("\tnextCost:", nextCost, nextPos)
				nextNode := &Node{
					Position:  nextPos,
					Cost:      nextCost,
					Heuristic: nextHeuristic,
					Parent:    current,
					Moves:     nextMovesCount(current, i),
	//				Moves: current.Moves+1,
					Direction: indexOf(directions, dir),
					}
					heap.Push(&pq, nextNode)
			}
		}
	}

	return nil, fmt.Errorf("path not found")
}

func heuristic (a, b GridCell) int {
	return abs(a.X - b.X) + abs(a.Y - b.Y)
}

func abs (x int ) int {
	if x < 0 {
		return -x
	}
	return x
}

func isValidMove(current *Node, nextPos GridCell, grid [][]int, nextDirectionIndex int) bool {
	// Check if next position is outside the grid
	if nextPos.X < 0 || nextPos.X >= len(grid) || nextPos.Y < 0 || nextPos.Y >= len(grid[0]) {
		return false
	}
/*
	// Check if the crucible is reversing direction (not allowed)
	if isReversingDirection(current.Direction, nextDirectionIndex) {
		return false
	}

	// Check if the crucible has exceeded the maximum moves in the current direction
	if current.Moves >= 3 && current.Direction == nextDirectionIndex {
		return false
	}
*/
	return true
}

// isReversingDirection checks if the crucible is trying to reverse its direction
func isReversingDirection(currentDirection, nextDirection int) bool {
	// Assuming directions are represented as 0: up, 1: right, 2: down, 3: left
	// A reversal would be (currentDirection + 2) % 4 == nextDirection
	return (currentDirection + 2) % 4 == nextDirection
}

func nextMovesCount (current *Node, nextDirectionIndex int) int {
	// If the crucible is just starting or if it's changing direction, reset the move count to 1.
	// We check if the direction is changing by comparing the current direction index with the next direction index.
	if current.Direction == -1 || current.Direction != nextDirectionIndex {
		return 1
	}

	// If the crucible is continuing in the same direction, increment the move count.
	// However, if the move count is already at 3, it means the crucible must turn, so we reset it to 1.
	if current.Moves < 3 {
		return current.Moves + 1
	}

	// This situation should not occur if the moves are being calculated correctly,
	// but it's a good idea to have a safe fallback.
	return 1
}

// reconstructPath backtracks from the destination to the start node
func reconstructPath(endNode *Node) []GridCell {
	var path []GridCell
	current := endNode

	// Loop back through the parent nodes
	for current != nil {
		path = append(path, current.Position)
		current = current.Parent
	}

	// Reverse the path to start from the beginning
	reversePath(path)

	return path
}

// reversePath reverses the order of the slice
func reversePath(path []GridCell) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func readFileInt(fname string) [][]int {
	var lines [][]int
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineI []int
		lineT := scanner.Text()
		for _, number := range lineT {
			num, _ := strconv.Atoi(string(number))
			lineI = append(lineI, num)
		}
		lines = append(lines, lineI)
	}
	return lines
}

func heatLoss(puzzle [][]int, points []GridCell) int {
	sum := 0
	for _, point := range points {
		sum += puzzle[point.Y][point.X]
	}
	return sum
}

func main() {

	timeStart := time.Now()
	// INPUT := "input.txt"
	INPUT := "test.txt"
	sumPart1, sumPart2 := 0, 0
	puzzle := readFileInt(INPUT)

	start := GridCell{0, 0, 0}

	end := GridCell{len(puzzle[0]), len(puzzle), puzzle[len(puzzle)-1][len(puzzle[0])-1]}

	points, err := AStarSearch(puzzle, start, end)
	if err != nil {
		fmt.Println("Error in Astar", err)
	}
	fmt.Println(points)
	sumPart1 = heatLoss(puzzle, points)

	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}
