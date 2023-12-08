package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func solve(t int, d int) int {
	delta := math.Sqrt(math.Pow(float64(t),2) - 4*float64(d))
	lo := (float64(t) - delta) / 2
	hi := (float64(t) + delta) / 2
	return int(math.Ceil(hi) - math.Floor(lo) - 1)
}

func ConvertLinetoInt(line string, start int) []int {
	var temp []int
	for _, number := range strings.Split(line, " ")[start:] {
		conv, err := strconv.Atoi(number)
		if err == nil && conv > 0 {
			temp = append(temp, conv)		
		}
	
	}
	return temp
}

func ConcatenateArray(input []int) string {
	temp := ""
	for _, element := range input {
		temp += strconv.Itoa(element)
	}
	return temp
}

func main() {
	var lines []string
	INPUT := "input_6.txt"
	
	file, err := os.Open(INPUT)

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	
	times := ConvertLinetoInt(lines[0], 1)
	distances := ConvertLinetoInt(lines[1], 1)
	part1 := 1
	for i := 0; i<len(times); i++ {
		part1 = part1 * solve(times[i], distances[i])
	}
	fmt.Println("Part1: ", part1)
	// now concatenate all the numbers in the array to get only one int
	time,_ := strconv.Atoi(ConcatenateArray(times))
	distance,_ := strconv.Atoi(ConcatenateArray(distances))
	
	fmt.Println("Part2: ", solve(time, distance))
	
}