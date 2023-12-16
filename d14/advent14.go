package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
	"bytes"
)

const (
	BOLDER = 79
	NOTHING = 46
)

func readFile(fname string) [][]byte {
	var lines [][]byte
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []byte(line))
	}
	return lines
}

func sumUp(puzzle [][]byte) int {
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])

	sum := 0
	for line := 0; line < puzzleH; line++ {
		for col := 0; col < puzzleW; col++ {
			if puzzle[line][col] == BOLDER {
				sum += puzzleH - line
			}
		} 
	}
	return sum
}

func tiltNorth(inputPuzzle [][]byte) ([][]byte, int) {
	puzzle := inputPuzzle
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	move := 0
	for line := 1; line < puzzleH; line++ {
		for col := 0; col < puzzleW; col++ {
			if string(puzzle[line][col]) == "O" {
				// Rolling bolder can go up
				switch puzzle[line-1][col] {
				case NOTHING:
					puzzle[line-1][col] = BOLDER
					puzzle[line][col] = NOTHING
					move++
				}
			}
		}
	}
	if move == 0 {
		return puzzle, 0
	} else {
		return tiltNorth(puzzle)
	}
}

func tiltSouth(inputPuzzle [][]byte) ([][]byte, int) {
	puzzle := inputPuzzle
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	move := 0
	for line := 0; line < puzzleH-1; line++ {
		for col := 0; col < puzzleW; col++ {
			if string(puzzle[line][col]) == "O" {
				// Rolling bolder can go up
				switch puzzle[line+1][col] {
				case NOTHING:
					puzzle[line+1][col] = BOLDER
					puzzle[line][col] = NOTHING
					move++
				}
			}
		}
	}
	if move == 0 {
		return puzzle, 0
	} else {
		return tiltSouth(puzzle)
	}
}

func tiltWest(inputPuzzle [][]byte) ([][]byte, int) {
	puzzle := inputPuzzle
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	move := 0
	for line := 0; line < puzzleH; line++ {
		for col := 1; col < puzzleW; col++ {
			if string(puzzle[line][col]) == "O" {
				// Rolling bolder can go up
				switch puzzle[line][col-1] {
				case NOTHING:
					puzzle[line][col-1] = BOLDER
					puzzle[line][col] = NOTHING
					move++
				}
			}
		}
	}
	if move == 0 {
		return puzzle, 0
	} else {
		return tiltWest(puzzle)
	}
}

func tiltEast(inputPuzzle [][]byte) ([][]byte, int) {
	puzzle := inputPuzzle
	puzzleH := len(puzzle)
	puzzleW := len(puzzle[0])
	move := 0
	for line := 0; line < puzzleH; line++ {
		for col := 0; col < puzzleW-1; col++ {
			if string(puzzle[line][col]) == "O" {
				// Rolling bolder can go up
				switch puzzle[line][col+1] {
				case NOTHING:
					puzzle[line][col+1] = BOLDER
					puzzle[line][col] = NOTHING
					move++
				}
			}
		}
	}
	if move == 0 {
		return puzzle, 0
	} else {
		return tiltEast(puzzle)
	}
}

func printPuzzle(puzzle [][]byte) {
	//puzzleH := len(puzzle)
	for i, _ := range puzzle {
		fmt.Println(string(puzzle[i]))	
	}
}

func main() {
	var cache = make(map[string]int)
	timeStart := time.Now()
	INPUT := "input_14.txt"
	// INPUT := "test_14.txt"
	
	fileContent := readFile(INPUT)
	
	puzzle,_ := tiltNorth(fileContent)
	
	load_part1 := sumUp(puzzle)
	fmt.Println("Load part1:", load_part1)
	
	puzzle_part2 := readFile(INPUT)
	breakP := 1000000000	
	for cycle:=0; cycle<1000000000; cycle++ {
		puzzle_part2, _ = tiltNorth(puzzle_part2)
		puzzle_part2, _ = tiltWest(puzzle_part2)
		puzzle_part2, _ = tiltSouth(puzzle_part2)
		puzzle_part2, _ = tiltEast(puzzle_part2)
		key := string(bytes.Join(puzzle_part2, []byte{}))
		if i, ok := cache[key]; ok {
			breakP = cycle + (1000000000-i)%(cycle-i)-1
		} else {
			cache[key] = cycle
		}
		
		if cycle == breakP {
			fmt.Println("Cycle: ", cycle)
			break
		}
		
	}
	
	load_part2 := sumUp(puzzle_part2)
	fmt.Println("Load part2:", load_part2)
	
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}