package main

import (
	"fmt"
	"strconv"

	"github.com/bokuzminski/aoc2021/helpers"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fileContent, _ := helpers.ReadDataToString("input.txt")

	gamma := 0
	epsilon := 0

	for index := range fileContent[0] {
		most, least := leastOrMostValue(fileContent, index)
		gamma <<= 1
		epsilon <<= 1

		if most == '1' {
			gamma |= 1
		}
		if least == '1' {
			epsilon |= 1
		}
	}

	fmt.Println(gamma * epsilon)
}

func partTwo() int {
	fileContent, _ := helpers.ReadDataToString("input.txt")

	oxygen := searchValue(fileContent, oxygenRating)
	co2 := searchValue(fileContent, co2Rating)

	fmt.Println(oxygen * co2)

	return (oxygen * co2)
}

func searchValue(input []string, filter func(rune, rune) rune) int {
	remaining := make([]string, len(input))
	copy(remaining, input)
	offset := 0

	for len(remaining) > 1 {
		remainderSlice := make([]string, 0, len(remaining))
		target := filter(leastOrMostValue(remaining, offset))

		for _, line := range remaining {
			if rune(line[offset]) == target {
				remainderSlice = append(remainderSlice, line)
			}
		}
		remaining = remainderSlice
		offset++
	}

	result, _ := strconv.ParseInt(remaining[0], 2, 32)

	return int(result)
}

func oxygenRating(highest, _ rune) rune {
	return highest
}
func co2Rating(_, lowest rune) rune {
	return lowest
}

func leastOrMostValue(values []string, idx int) (rune, rune) {
	zero := 0
	one := 0
	for _, number := range values {
		if number[idx] == '0' {
			zero++
		} else {
			one++
		}
	}
	if zero > one {
		return '0', '1'
	}

	return '1', '0'
}
