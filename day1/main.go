package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	prevDepth, _ := strconv.Atoi(scanner.Text())

	var result = 0

	for scanner.Scan() {
		currentDepth, _ := strconv.Atoi(scanner.Text())
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
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	var result = 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines)-3; i++ {
		a, _ := strconv.Atoi(lines[i])
		b, _ := strconv.Atoi(lines[i+1])
		c, _ := strconv.Atoi(lines[i+2])
		d, _ := strconv.Atoi(lines[i+3])

		if (a + b + c) < (b + c + d) {
			result++
		}
	}

	fmt.Println(result)
}
