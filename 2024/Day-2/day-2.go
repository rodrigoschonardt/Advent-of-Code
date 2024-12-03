package main

import (
	"aoc"
	"fmt"
	"strconv"
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

	count := 0

	for i := 0; i < len(list); i++ {
		line := list[i]

		isDecreasing := true

		for x := 0; x < len(line)-1; x++ {
			if x == 0 && line[x] < line[x+1] {
				isDecreasing = false
			}

			diff := line[x+1] - line[x]

			if diff > 0 && isDecreasing {
				break
			}

			if diff < 0 && !isDecreasing {
				break
			}

			if diff > 3 || diff < -3 || diff == 0 {
				break
			}

			if x+1 == len(line)-1 {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

func part2() {
	list := getList()
	errorList := make([][]int, 0)

	count := 0

	for i := 0; i < len(list); i++ {
		line := list[i]

		isDecreasing := true

		for x := 0; x < len(line)-1; x++ {
			if x == 0 && line[x] < line[x+1] {
				isDecreasing = false
			}

			diff := line[x+1] - line[x]

			if diff > 0 && isDecreasing {
				errorList = append(errorList, line)
				break
			}

			if diff < 0 && !isDecreasing {
				errorList = append(errorList, line)
				break
			}

			if diff > 3 || diff < -3 {
				errorList = append(errorList, line)
				break
			}

			if diff == 0 {
				errorList = append(errorList, line)
				break
			}

			if x+1 == len(line)-1 {
				count++
				break
			}
		}
	}

	for i := 0; i < len(errorList); i++ {
		originalLine := errorList[i]

		for removeIndex := 0; removeIndex < len(originalLine); removeIndex++ {
			line := append([]int{}, originalLine[:removeIndex]...)
			if removeIndex+1 < len(originalLine) {
				line = append(line, originalLine[removeIndex+1:]...)
			}

			stoppedEarly := false
			isDecreasing := true

			for x := 0; x < len(line)-1; x++ {
				if x == 0 && line[x] < line[x+1] {
					isDecreasing = false
				}

				diff := line[x+1] - line[x]

				if diff > 0 && isDecreasing {
					break
				}

				if diff < 0 && !isDecreasing {
					break
				}

				if diff > 3 || diff < -3 {
					break
				}

				if diff == 0 {
					break
				}

				if x+1 == len(line)-1 {
					count++
					stoppedEarly = true
					break
				}
			}

			if stoppedEarly {
				break
			}
		}
	}

	fmt.Println(count)
}

func getList() [][]int {
	input := aoc.GetFileContents()

	lines := strings.Split(input, "\n")

	length := len(lines)

	list := make([][]int, length)

	for i, line := range lines {
		numbers := strings.Split(line, " ")
		list[i] = make([]int, len(numbers))

		for x, number := range numbers {
			list[i][x], _ = strconv.Atoi(number)
		}
	}

	return list
}
