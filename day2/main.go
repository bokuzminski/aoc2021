package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileContent, _ := ioutil.ReadFile("input.txt")
	depth := 0
	horizontal := 0
	aim := 0

	split := strings.Split(string(fileContent), "\n")
	for _, i := range split {
		var (
			operation string
			value     int
		)
		fmt.Sscanf(string(i), "%s %d", &operation, &value)
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
	partTwo()

}

func partTwo() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fileContent, _ := ioutil.ReadFile("input.txt")
	depth := 0
	horizontal := 0
	aim := 0

	split := strings.Split(string(fileContent), "\n")
	for _, i := range split {
		var (
			operation string
			value     int
		)
		fmt.Sscanf(string(i), "%s %d", &operation, &value)
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
