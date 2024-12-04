package main

import (
	"aoc"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Print("Part 1: ")
	part1()
	fmt.Print("Part 2: ")
	part2()
}

func part1() {
	input := aoc.GetFileContents()

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re, err := regexp.Compile(pattern)

	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	matches := re.FindAllStringSubmatch(input, -1)

	count := 0

	for _, match := range matches {
		value1, _ := strconv.Atoi(match[1])
		value2, _ := strconv.Atoi(match[2])

		count += value1 * value2
	}

	fmt.Println(count)

}

func part2() {
	input := aoc.GetFileContents()

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re, _ := regexp.Compile(pattern)

	matches := re.FindAllStringSubmatchIndex(input, -1)

	pattern = `do\(\)`

	re, _ = regexp.Compile(pattern)

	doMatches := re.FindAllStringIndex(input, -1)

	pattern = `don't\(\)`

	re, _ = regexp.Compile(pattern)

	dontMatches := re.FindAllStringIndex(input, -1)

	count := 0

	for _, match := range matches {
		index := match[0]

		doIndex := 0
		dontIndex := -1

		for _, do := range doMatches {
			if do[1] > index {
				break
			}

			doIndex = do[1]
		}

		for _, dont := range dontMatches {
			if dont[1] > index {
				break
			}

			dontIndex = dont[1]
		}

		if doIndex < dontIndex {
			continue
		}

		value1, _ := strconv.Atoi(input[match[2]:match[3]])
		value2, _ := strconv.Atoi(input[match[4]:match[5]])

		count += value1 * value2
	}

	fmt.Println(count)
}
