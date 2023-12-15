package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
	"strings"
	
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


func hash(input string) int {
	
	sum := 0
	for i,_ := range input {
		sum += int(input[i])
		sum *= 17
		sum = sum % 256		
	}
	return sum
}

func main() {
	// var cache = make(map[string]int)
	timeStart := time.Now()
	INPUT := "input_15.txt"
	// INPUT := "test_15.txt"	
	fileContent := strings.Split(strings.Join(readFile(INPUT), "\n"), ",")
	sum_part1 := 0
	for _, value := range fileContent {
		sum_part1 += hash(value)
	}
	fmt.Println("Part1:", sum_part1)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}