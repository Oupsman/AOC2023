package main
import (
	"strings"
	"os"
	"bufio"
	"strconv"
	"fmt"
)
func convertLinetoInt(line string, start int) []int {
	var temp []int
	for _, number := range strings.Split(line, " ")[start:] {
		conv, err := strconv.Atoi(number)
		if err == nil {
			temp = append(temp, conv)		
		}

	}
	return temp
}

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

func sumArray(array []int) int {
	sum := 0
	for _, num := range array {
		if num == 0 {
			sum++
		}
	}
	return sum
}

func diff (array []int) []int {
	var res []int
	for i:=1; i<len(array); i++ {
		diff := array[i] - array[i-1]
		res = append(res, diff)
	}
	if len(res) == 0 {
		res = append(res, 0)

	}
	return res
}

func solvePart1(line []int) int {
//	fmt.Println("Solving line:", line)
	if len(line) == 1 {
		return line[0]
	}
	
	var result int
	deltas := [][]int{}
	
	i := 0
	deltas = append(deltas, line)
	
	for sumArray(deltas[i]) != len(deltas[i]) {
//		fmt.Println("deltas:", deltas, "iteration: ", i)
		deltas = append(deltas, diff(deltas[i]))
		i++
	}
//	fmt.Println("Deltas", deltas, len(deltas))
	
	for i:=len(deltas)-1; i>0; i-- {

		value := deltas[i-1][len(deltas[i-1])-1] + deltas[i][len(deltas[i])-1]
		deltas[i-1] = append(deltas[i-1], value)
	}
	
	result = deltas[0][len(deltas[0])-1]
	fmt.Println("result: ", result)
	return result
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	INPUT := "input_9.txt"
//	INPUT := "dumb_test.txt"
//	INPUT := "test_9.txt"
	lines := readFile(INPUT)
	sum_part1 := 0
	for lineCounter, line := range lines {
		fmt.Println("Line: ", lineCounter)
		intLine := convertLinetoInt(line, 0)
		sum_part1 += solvePart1(intLine)
	}
	fmt.Println("Part1:", sum_part1)
	sum_part2 := 0
	for lineCounter, line := range lines {
		fmt.Println("Line: ", lineCounter)
		intLine := reverse(convertLinetoInt(line, 0))
		sum_part2 += solvePart1(intLine)
	}
	fmt.Println("Part2:", sum_part2)
	
}

