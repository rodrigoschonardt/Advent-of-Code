package main

import (
	"aoc"
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Part 1: ")
	part1()
	fmt.Print("Part 2: ")
	part2()
}

func part1() {
	list := getList()

	targetValue := "XMAS"
	targetValueReverse := "SAMX"

	count := 0

	for i := 4; i < len(list)-4; i++ {
		line := list[i]

		for x := 4; x < len(line)-4; x++ {

			value := line[x] + line[x+1] + line[x+2] + line[x+3]

			if value == targetValue || value == targetValueReverse {
				count++
			}

			value = line[x] + list[i+1][x] + list[i+2][x] + list[i+3][x]

			if value == targetValue || value == targetValueReverse {
				count++
			}

			value = line[x] + list[i-1][x+1] + list[i-2][x+2] + list[i-3][x+3]

			if value == targetValue || value == targetValueReverse {
				count++
			}

			value = line[x] + list[i+1][x+1] + list[i+2][x+2] + list[i+3][x+3]

			if value == targetValue || value == targetValueReverse {
				count++
			}

		}
	}

	fmt.Println(count)
}

func part2() {
	list := getList()

	count := 0

	for i := 4; i < len(list)-4; i++ {
		line := list[i]

		for x := 4; x < len(line)-4; x++ {

			s1 := line[x] + list[i+1][x+1] + list[i+2][x+2]
			s2 := line[x+2] + list[i+1][x+1] + list[i+2][x]

			if (s1 == "MAS" || s1 == "SAM") && (s2 == "MAS" || s2 == "SAM") {
				count++
			}
		}
	}

	fmt.Println(count)
}

func getList() [][]string {
	input := aoc.GetFileContents()

	lines := strings.Split(input, "\n")

	length := len(lines)
	originalWidth := len(lines[0])

	paddedLength := length + 8
	paddedWidth := originalWidth + 8

	paddedList := make([][]string, paddedLength)
	for i := range paddedList {
		paddedList[i] = make([]string, paddedWidth)
		for j := range paddedList[i] {
			paddedList[i][j] = "."
		}
	}

	for i, line := range lines {
		for x, char := range line {
			paddedList[i+4][x+4] = string(char)
		}
	}

	return paddedList
}
