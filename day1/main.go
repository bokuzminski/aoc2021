package main

import (
	"fmt"

	"github.com/bokuzminski/aoc2021/helpers"
)

func partOne() {
	lines, _ := helpers.ReadDataToInt("input.txt")
	var result = 0
	prevDepth := lines[0]

	for _, val := range lines {
		currentDepth := val
		if currentDepth > prevDepth {
			result++
		}
		prevDepth = currentDepth
	}

	fmt.Println(result)
}

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	lines, _ := helpers.ReadDataToInt("input.txt")
	var result = 0

	for i := 3; i < len(lines); i++ {
		a := lines[i]
		b := lines[i-3]
		if a > b {
			result++
		}
	}

	fmt.Println(result)
}
