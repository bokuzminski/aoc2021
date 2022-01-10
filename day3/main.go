package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileContent, _ := ioutil.ReadFile("sample.txt")

	split := strings.Split(string(fileContent), "\n")

	gamma := 0
	epsilon := 0

	for index := range split[0] {

		most, least := leastOrMostValue(split, index)
		gamma <<= 1
		epsilon <<= 1
		fmt.Print(gamma)
		if most == '1' {
			gamma |= 1
		}
		if least == '1' {
			epsilon |= 1
		}
	}

	//fmt.Println(gamma * epsilon)
}

func partOne() {}
func partTwo(input []string) {
}

func oxygenRating(highest, _ rune) rune {
	return highest
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
