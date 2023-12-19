package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Rating map[string]int64

type Instruction struct {
	part           string
	quantity       int64
	superior       bool
	targetWorkflow string
	follow         bool
	accept         bool
	reject         bool
}

type Workflow map[string][]Instruction

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

func executeWorkflow(ratings Rating, workflows Workflow, workflow string) bool {
	// sum := int64(0)
	fmt.Println("Executing workflow:", workflow)
	instructions := workflows[workflow]
	for _, instruction := range instructions {
		valid := false
		_, ok := ratings[instruction.part]
		if ok {
			// If there is a instruction for this variable
			if instruction.quantity < ratings[instruction.part] && instruction.superior {
				valid = true
			} else if instruction.quantity > ratings[instruction.part] && !instruction.superior {
				valid = true
			}
			if instruction.quantity > 0 && valid && instruction.follow {
				return executeWorkflow(ratings, workflows, instruction.targetWorkflow)
			}
			if valid && instruction.accept {
				return true
			}
		} else if instruction.follow && instruction.part == "" {
			return executeWorkflow(ratings, workflows, instruction.targetWorkflow)
			break
		} else if instruction.quantity == 0 && instruction.accept {
			return true
		} else if instruction.quantity == 0 && instruction.reject {
			fmt.Println("REJECT !")
			return false
		}

	}
	// Shouldn't be here, just being safe
	return false
}

func executeWorkflows(ratings Rating, workflows Workflow) int64 {

	// First execute the in workflow
	val := executeWorkflow(ratings, workflows, "in")
	fmt.Println("Val:", val)
	sum := int64(0)
	if val {
		for _, rating := range ratings {
			sum += rating		
		}
	}
	return sum
}

func solve(input []string, part2 bool) int64 {
	decomposeLine := regexp.MustCompile(`^([a-z]+){(.*)}$`)
	decomposeInstructions := regexp.MustCompile(`([ARa-z]+)([<>]?)([0-9]+?):?([ARa-z]+?)$`)
	var workflows = make(Workflow)
	var sum int64
	inWorkflows := true
	for index, line := range input {
		if input[index] == "" {
			inWorkflows = false
			continue
		}
		if inWorkflows {
			var instructions []Instruction
			match := decomposeLine.FindStringSubmatch(line)
			// workflows.name = match[1]
			for _, instruction := range strings.Split(match[2], ",") {
				var targetInstruction Instruction
				matchI := decomposeInstructions.FindStringSubmatch(instruction)
				if len(matchI) > 0 {
					targetInstruction.part = matchI[1]
					targetInstruction.superior = matchI[2] == ">"
					targetInstruction.accept = matchI[4] == "A"
					targetInstruction.reject = matchI[4] == "R"
					if matchI[4] != "A" && matchI[4] != "R" {
						targetInstruction.targetWorkflow = matchI[4]
						targetInstruction.follow = true

					}
					num, _ := strconv.Atoi(matchI[3])
					targetInstruction.quantity = int64(num)
				} else if instruction == "A" || instruction == "R" {
					targetInstruction.accept = instruction == "A"
					targetInstruction.reject = instruction == "R"
				} else {
					targetInstruction.targetWorkflow = instruction
					targetInstruction.follow = true
				}

				instructions = append(instructions, targetInstruction)
				workflows[match[1]] = instructions
			}
			continue
		} else {
			regexpParts := regexp.MustCompile(`{(.+)}`)
			rawParts := regexpParts.FindStringSubmatch(line)
			// fmt.Println(line, rawParts)
			parts := strings.Split(rawParts[1], ",")
			var ratings = Rating{}
			for _, part := range parts {
				//var rating Rating
				rawPart := strings.Split(part, "=")
				num, _ := strconv.Atoi(rawPart[1])
				ratings[rawPart[0]] = int64(num)
			}
			//			fmt.Println("Ratings:", ratings)
			val := executeWorkflows(ratings, workflows)
			//			fmt.Println(" Value:", val)

			sum += val
		}

	}
	return sum
}

func main() {
	timeStart := time.Now()
	// INPUT := "test.txt"
	INPUT := "input.txt"

	fileContent := readFile(INPUT)

	sumPart1 := solve(fileContent, false)

	// sumPart2 := solve(fileContent, true)

	fmt.Println("Part1:", sumPart1)
	// fmt.Println("Part2:", sumPart2)

	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}
