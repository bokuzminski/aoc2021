package main

import (
	"fmt"

	"github.com/bokuzminski/aoc2021/helpers"
)

func main() {
	partOne()
	partTwo()

}

func partOne() {
	lines, _ := helpers.ReadDataToString("input.txt")
	var horizontal, depth, aim int

	for _, val := range lines {
		var (
			operation string
			value     int
		)
		fmt.Sscanf(val, "%s %d", &operation, &value)

		switch operation {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
			depth -= value
		case "down":
			aim += value
			depth += value
		}
	}

	fmt.Println("Part 1 answer : ", depth*horizontal)
}

func partTwo() {
	lines, _ := helpers.ReadDataToString("input.txt")
	var horizontal, depth, aim int

	for _, i := range lines {
		var (
			operation string
			value     int
		)
		fmt.Sscanf(i, "%s %d", &operation, &value)

		switch operation {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Println("Part 2 answer: ", depth*horizontal)
}
