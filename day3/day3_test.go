package main

import (
	"testing"

	"github.com/bokuzminski/aoc2021/helpers"
)

func TestDay3PartTwo(t *testing.T) {
	testData, _ := helpers.ReadDataToString("sample.txt")
	oxygen := searchValue(testData, oxygenRating)
	co2 := searchValue(testData, co2Rating)

	day3Result := co2 * oxygen

	if day3Result != 230 {
		t.Errorf("Wrong result for day 3")
	}
}
