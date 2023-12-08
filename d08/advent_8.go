package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"slices"
)

type nodes struct {
	left, right string
	
}

func readFile(fname string) (map[string]nodes, string) {
	lines := map[string]nodes{}
	var directions string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	searchNodes, _ := regexp.Compile("[A-Z]{3}")
	regexp_directions := "^[LR]+$"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		
		nodesB := searchNodes.FindAllString(line, 3)
		direction,_ := regexp.MatchString(regexp_directions,line)
		if direction {
			directions = line
		}
		if len(nodesB) > 1 {
			var node nodes

			node.left = nodesB[1]
			node.right = nodesB[2]
			lines[nodesB[0]] = node
		}
	}
	return lines, directions
}

func findNextNode(node string, nodes map[string]nodes, direction string) string {
	switch direction {
	case "L":
		return nodes[node].left
	case "R":
		return nodes[node].right
	}
	return "none"
}

func findDirection(directions string, intentMove int) string {
	
	// intentMove can be greater than the lenght of the directions string, so let's round it
	realMove := intentMove % len(directions)
	return string(directions[realMove])
	
}

func searchStartingNodes(nodes map[string]nodes) []string {

	var startingNodes []string
	var keys []string
	search, _ := regexp.Compile("[A-Z][A-Z]A")
	// keys := reflect.ValueOf(nodes).MapKeys()
	for k := range nodes {
		keys = append(keys, k)
	}

	for _, node := range keys {
		if search.FindString(node) != "" {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}
func searchEndingNodes(nodes map[string]nodes) []string {

	var Nodes []string
	var keys []string
	search, _ := regexp.Compile("[A-Z][A-Z]Z")
	// keys := reflect.ValueOf(nodes).MapKeys()
	for k := range nodes {
		keys = append(keys, k)
	}

	for _, node := range keys {
		if search.FindString(node) != "" {
			Nodes = append(Nodes, node)
		}
	}
	return Nodes
}

// greatest common divisor (GCD) via Euclidean algorithm

func GCD (a,b int) int {
	
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
	
}

func LCM (numbers []int, index int) int {
	if index == len(numbers) - 1 {
		return numbers[index]
	}
	a := numbers[index]
	b := LCM(numbers, index+1)
	return (a*b)/GCD(a,b)
}



func main() {
	INPUT := "input_8_ay.txt"
	nodes, directions := readFile(INPUT)
	node := "AAA"
	move_part1 := 0
	for node != "ZZZ" {
		direction := findDirection(directions, move_part1)
		node = findNextNode(node, nodes, direction)
		move_part1 += 1
	}
	fmt.Println("Moves: ", move_part1)
	var moves_part2 []int
	start := searchStartingNodes(nodes)
	end := searchEndingNodes(nodes)
	fmt.Println("Start: ", start)
	for _, startingNode := range start {
		fmt.Println("Searching path from ", startingNode)
		node := startingNode
		move := 0
		for ! slices.Contains(end, node) {
			direction := findDirection(directions, move)
			node = findNextNode(node, nodes, direction)
			move += 1
		}
		moves_part2 = append(moves_part2, move)
	}
	fmt.Println("Moves part2", moves_part2)
	fmt.Println("LCM: ", LCM(moves_part2, 0))
}
