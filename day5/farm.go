package main

import (
	"aoc/common"
	"strconv"
	"strings"
	"unicode"
)

type Range struct {
	Destination int
	Source      int
	Length      int
}

func FindMapEntry(input int, mapSet []Range) int {
	for _, entry := range mapSet {
		if input > entry.Source && input < (entry.Source+entry.Length) {
			return entry.Destination + (input - entry.Source)
		}
	}
	return input
}

func ParseFile(input []string) ([]int, [][]Range) {

	seeds := make([]int, 0)
	sentence := strings.Split(strings.Trim(input[0][6:], " "), " ")
	seeds = common.StrArrayToIntArray(sentence)
	result := make([][]Range, 0)

	currentMap := make([]Range, 0)
	for _, line := range input[3:] {
		if line == "" {
			continue
		}
		// group separator
		if !unicode.IsNumber(rune(line[0])) {
			result = append(result, currentMap)
			currentMap = make([]Range, 0)
		} else {
			var tempRng *Range = new(Range)
			tmpValue := strings.Split(line, " ")
			val, _ := strconv.Atoi(tmpValue[0])
			tempRng.Destination = val
			val, _ = strconv.Atoi(tmpValue[1])
			tempRng.Source = val
			val, _ = strconv.Atoi(tmpValue[2])
			tempRng.Length = val
			currentMap = append(currentMap, *tempRng)
		}

	}
	return seeds, result
}

func FindClosestLocation(seeds []int, farmMap [][]Range) (int, error) {
	locations := make([]int, 0)
	for _, seed := range seeds {
		currentValue := seed
		for _, mapSet := range farmMap {

			currentValue = FindMapEntry(currentValue, mapSet)
		}
		locations = append(locations, currentValue)
	}
	return common.FindMin(locations)
}
