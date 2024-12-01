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
	list1, list2 := getLists()

	length := len(list1)

	count := 0

	for i := 0; i < length; i++ {

		smallest := list1[i]

		for x := i + 1; x < length; x++ {

			if smallest > list1[x] {
				list1[i] = list1[x]
				list1[x] = smallest
				smallest = list1[i]
			}
		}

		smallest = list2[i]

		for x := i + 1; x < length; x++ {

			if smallest > list2[x] {
				list2[i] = list2[x]
				list2[x] = smallest
				smallest = list2[i]
			}
		}

		if list1[i] > list2[i] {
			count += list1[i] - list2[i]
		} else {
			count += list2[i] - list1[i]
		}
	}

	fmt.Println(count)
}

func part2() {
	list1, list2 := getLists()

	mapNumbers := make(map[int]int)

	for _, number := range list2 {
		val, ok := mapNumbers[number]

		if ok {
			mapNumbers[number] = val + 1
		} else {
			mapNumbers[number] = 1
		}
	}

	similarity := 0

	for _, number := range list1 {
		val, ok := mapNumbers[number]

		if ok {
			similarity += number * val
		}
	}

	fmt.Println(similarity)
}

func getLists() ([]int, []int) {
	input := aoc.GetFileContents()

	lines := strings.Split(input, "\n")

	length := len(lines)

	list1 := make([]int, length)
	list2 := make([]int, length)

	for i, line := range lines {
		numbers := strings.Split(line, "   ")

		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])

		list1[i] = number1
		list2[i] = number2
	}

	return list1, list2
}
