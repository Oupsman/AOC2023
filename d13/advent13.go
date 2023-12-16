package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
)


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

func findHorSymetry(puzzle []string, smudge bool) int {
	puzzleW := len(puzzle[0])
	puzzleH := len(puzzle)
	
// 	for counter, _ := range puzzle {
	for counter := 0; counter < puzzleH -1; counter++ {
		diff := 0
		for col := 0; col < puzzleW; col++ {
			for offset := 0;; offset++ {
				lowRow := counter - offset
				highRow := counter + offset + 1
				if lowRow < 0 || highRow >= puzzleH {
					break;
				}
				if puzzle[lowRow][col] != puzzle[highRow][col] {
					diff++
				}
			}
		}
		if smudge && diff == 1 {
			return counter + 1
		} else if ! smudge && diff == 0 {
			return counter + 1
		}
	}
	return 0
}

func findVertSymetry(puzzle []string, smudge bool) int {
	columns := make([]string, len(puzzle[0]))
	
	for j:=0; j<len(puzzle[0]); j++ {
		var column string
		for i:=0; i < len(puzzle); i++  {
			column += string(puzzle[i][j])
		}
		columns[j] = column
	}
	return findHorSymetry(columns, smudge)
}

func main() {

	timeStart := time.Now()

	INPUT := "input_13.txt"
	// INPUT := "test_13.txt"
	var puzzles [][]string
	lines := readFile(INPUT)
	counter := 0
	for counter < len(lines) {
		var puzzle []string
		
		for lines[counter] != "" && counter < len(lines){
			puzzle = append(puzzle, lines[counter])
			counter++
			if counter == len(lines) {
				break
			}
		}
		puzzles = append(puzzles, puzzle)
		counter++

	}

	sum_part1 := 0
	sum_part2 := 0
	for _, puzzle := range puzzles {
		sum_part1 += findVertSymetry(puzzle, false)
		sum_part1 += findHorSymetry(puzzle, false)*100
		sum_part2 += findVertSymetry(puzzle, true)
		sum_part2 += findHorSymetry(puzzle, true)*100			
		
	}
	fmt.Println("Part1:", sum_part1)
	fmt.Println("Part2:", sum_part2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}