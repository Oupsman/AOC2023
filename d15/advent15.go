package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

type Lens struct {
	Label string
	Length int
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


func hash(input string) int {
	
	sum := 0
	for i := range input {
		sum += int(input[i])
		sum *= 17
		sum = sum % 256		
	}
	return sum
}

func main() {
	timeStart := time.Now()
	INPUT := "input_15.txt"
	// INPUT := "test_15.txt"	

	boxes := [256][]Lens{}
	sumPart2 := 0
	sumPart1 := 0

	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)

	for _, input := range re.FindAllStringSubmatch(strings.Join(readFile(INPUT), "\n"), -1) {
		box := &boxes[hash(input[1])]
		index := slices.IndexFunc(*box, func(l Lens) bool { return l.Label == input[1] })
		if input[2] == "-" && index != -1 {
			*box = slices.Delete(*box, index, index+1)
		} else if input[2] == "=" && index != -1 {
			(*box)[index] = Lens{input[1], int(input[3][0] - '0')}
		} else if input[2] == "=" {
			*box = append(*box, Lens{input[1], int(input[3][0] - '0')})
		}
		sumPart1 += hash(input[0])
	}

	for i, b := range boxes {
		for j, l := range b {
			sumPart2 += (i + 1) * (j + 1) * l.Length
		}
	}

	fmt.Println("Part1:", sumPart1)
	fmt.Println("Part2:", sumPart2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)

}