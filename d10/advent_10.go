package main

import (
	"bufio"
	"os"
	"strings"

	"fmt"
)

func contains(val int, arr []int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
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

func getTransformation(currentDir int, char string) int {
	// fmt.Println("currentDir:", currentDir, "char:", char)
	switch currentDir {
	case 0:
		switch char {
		case "-":
			return 0
		case "7":
			return 1
		case "J":
			return 3

		}
	case 2:
		switch char {
		case "-":
			return 2
		case "F":
			return 1
		case "L":
			return 3

		}
	case 1:
		switch char {
		case "|":
			return 1
		case "L":
			return 0
		case "J":
			return 2
		}
	case 3:
		switch char {
		case "|":
			return 3
		case "F":
			return 0

		case "7":
			return 2
		}
	}
	return -1
}

func main() {
	// INPUT := "test_10.txt"
	INPUT := "input_10.txt"

	fileContent := readFile(INPUT)

	mazeH := len(fileContent)
	mazeW := len(fileContent[0])

	startX := -1
	startY := -1
	// Scan the maze to get the exit is
	for y := 0; y < mazeH; y++ {
		for x := 0; x < mazeW; x++ {

			if string(fileContent[x][y]) == string("S") {
				startX = x
				startY = y
			}
		}
	}
	fmt.Println("Start position: (", startX, ",", startY, ")")

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	happy := []string{"-7J", "|LJ", "-FL", "|F7"}
	Sdirs := []int{}
	for i := 0; i < 4; i++ {
		pos := dirs[i]
		bx := startX + pos[0]
		by := startY + pos[1]
		if bx >= 0 && bx <= mazeH && by >= 0 && by <= mazeW && strings.Contains(happy[i], string(fileContent[bx][by])) {
			Sdirs = append(Sdirs, i)
		}
	}
	Svalid := contains(3, Sdirs)

	fmt.Println("Sdirs:", Sdirs)
	curdir := Sdirs[0]
	cx := startX + dirs[curdir][0]
	cy := startY + dirs[curdir][1]
	ln := 1
	var O [1000][1000]int
	O[cx][cy] = 1
	for !(cx == startX && cy == startY) {
		O[cx][cy] = 1
		ln += 1
		curdir = getTransformation(curdir, string(fileContent[cx][cy]))

		cx = cx + dirs[curdir][0]
		cy = cy + dirs[curdir][1]
	}

	fmt.Println("Length:", ln)
	fmt.Println("length / 2:", ln/2)
	// Part 2
	ct := 0
	for i := 0; i < mazeH; i++ {
		inn := false
		for j := 0; j < mazeW; j++ {
			if O[i][j] == 1 {
				if string(fileContent[i][j]) == "|" ||
					string(fileContent[i][j]) == "J" ||
					string(fileContent[i][j]) == "L" ||
					(string(fileContent[i][j]) == "S" && Svalid) {
					inn = !inn
				}
			} else {
				if inn {
					ct += 1
				}
			}
		}
	}
	fmt.Println(ct)

}
