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
	distance int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
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

func dijkstra(grid [][]int, start, end [2]int) [][2]int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // Déplacements possibles : bas, haut, droite, gauche

	// Initialisation des distances à l'infini sauf pour le point de départ
	distances := make([][]int, rows)
	for i := 0; i < rows; i++ {
		distances[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distances[i][j] = 1<<31 - 1
		}
	}
	distances[start[0]][start[1]] = 0

	// Initialisation de la file de priorité avec le point de départ
	priorityQueue := make(PriorityQueue, 0)
	heap.Push(&priorityQueue, &Node{start[0], start[1], 0})

	for len(priorityQueue) > 0 {
		node := heap.Pop(&priorityQueue).(*Node)

		if [2]int{node.x, node.y} == end {
			break // On a atteint le point d'arrêt, sortie de la boucle
		}

		if node.distance > distances[node.x][node.y] {
			continue // Ignorer les nœuds déjà visités
		}

		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			x, y := node.x+dx, node.y+dy

			if 0 <= x && x < rows && 0 <= y && y < cols {
				newDistance := node.distance + grid[x][y]

				if newDistance < distances[x][y] {
					distances[x][y] = newDistance
					heap.Push(&priorityQueue, &Node{x, y, newDistance})
				}
			}
		}
	}

	// Reconstruction du chemin
	path := make([][2]int, 0)
	x, y := end[0], end[1]

	for [2]int{x, y} != start {
		path = append(path, [2]int{x, y})
		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			nx, ny := x+dx, y+dy
			if 0 <= nx && nx < rows && 0 <= ny && ny < cols && distances[nx][ny]+grid[x][y] == distances[x][y] {
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
