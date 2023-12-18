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

type Node struct {
	x, y     int
	heatLoss int
	Dir      Dir
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].heatLoss < pq[j].heatLoss }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func dijkstra(puzzle [][]int, start, end [2]int) [][2]int {

	rows, cols := len(puzzle), len(puzzle[0])

	// Initialisation des heatLosses à l'infini sauf pour le point de départ
	heatLosses := make([][]int, rows)
	for i := 0; i < rows; i++ {
		heatLosses[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			heatLosses[i][j] = 1<<31 - 1
		}
	}

	fmt.Println("Starting dijkstra search")
	heatLosses[start[0]][start[1]] = 0

	// Initialisation de la file de priorité avec le point de départ
	priorityQueue := make(PriorityQueue, 0)
	heap.Push(&priorityQueue, &Node{start[0], start[1], puzzle[0][0], directions[E]})
	currentDir := E
	for len(priorityQueue) > 0 {
		fmt.Println("For loop")
		node := heap.Pop(&priorityQueue).(*Node)

		if [2]int{node.x, node.y} == end {
			break // On a atteint le point d'arrêt, sortie de la boucle
		}
		if node.heatLoss < heatLosses[node.y][node.x] {
			continue // Ignorer les nœuds déjà visités
		}
		leftDir := directions[left[currentDir]]
		rightDir := directions[right[currentDir]]
		straight := directions[currentDir]
		availableDirs := []Dir{leftDir, straight, rightDir}
		minHeatLoss := 1<<31 -1
		followDir := directions[currentDir]
		followX := 0
		followY := 0
		for _, direction := range availableDirs {
			dx, dy := direction.x, direction.y
			x, y := node.x+dx, node.y+dy
			
			if 0 <= x && x < rows && 0 <= y && y < cols {
				newHeatLoss := puzzle[y][x]
				fmt.Println("Heatloss: ", newHeatLoss)
				if newHeatLoss < minHeatLoss {
					fmt.Println("Following to :", x, y)
					minHeatLoss = newHeatLoss
					followDir = direction
					followX = x
					followY = y
				}

			}
		}
		heatLosses[followX][followY] = minHeatLoss
		fmt.Println("Current Coords: ", followX, followY)
		heap.Push(&priorityQueue, &Node{followX, followY, minHeatLoss, followDir})
	}

	// Reconstruction du chemin
	path := make([][2]int, 0)
	x, y := end[0], end[1]

	for [2]int{x, y} != start {

		path = append(path, [2]int{x, y})
		for _, direction := range directions {
			dx, dy := direction.x, direction.y
			nx, ny := x+dx, y+dy
			if 0 <= nx && nx < rows && 0 <= ny && ny < cols && heatLosses[nx][ny]+puzzle[x][y] == heatLosses[x][y] {
				x, y = nx, ny
				break
			}
		}
	}

	path = append(path, [2]int{start[0], start[1]})

	// Inverser le chemin pour qu'il soit du point de départ au point d'arrêt
	reversePath := make([][2]int, len(path))
	for i, j := 0, len(path)-1; i < len(path); i, j = i+1, j-1 {
		reversePath[i] = path[j]
	}
	return reversePath
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

func heatLoss(puzzle [][]int, points [][2]int) int {
	sum := 0
	for _, point := range points {
		sum += puzzle[point[1]][point[0]]
	}
	return sum
}

func main() {

	timeStart := time.Now()
	// INPUT := "input.txt"
	INPUT := "test.txt"
	sumPart1, sumPart2 := 0, 0
	puzzle := readFileInt(INPUT)

	start := [2]int{0, 0}
	end := [2]int{len(puzzle[0]), len(puzzle)}

	points := dijkstra(puzzle, start, end)
	fmt.Println(points)
	sumPart1 = heatLoss(puzzle, points)

	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}
