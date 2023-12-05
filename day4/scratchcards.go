package main

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

func CalculatePrize(input string) (int, int, error) {

	temp := strings.Split(strings.SplitAfter(input, ":")[1], "|")

	if len(temp) != 2 {
		return -1, -1, fmt.Errorf("%s is malformed, should have only two sets divided by |", input)
	}

	inputA := strings.Split(strings.Trim(temp[0], " "), " ")
	inputB := strings.Split(strings.Trim(temp[1], " "), " ")
	matches := common.Intersect(inputA, inputB)
	numMatches := len(matches)
	result := int(math.Pow(2, float64(numMatches-1)))
	return result, numMatches, nil
}

func GetAccumulatedScores(input []string) uint64 {
	var sum uint64 = 0
	multipliers := make([]uint64, len(input))

	for i := range multipliers {
		multipliers[i] = 1
	}
	scores := make([]uint64, len(input))

	for i, line := range input {
		_, numMatches, _ := CalculatePrize(line)
		scores[i] = uint64(numMatches)
		currentMultiplier := multipliers[i]
		// increase multipliers
		for j, _ := range common.SafeSlice(multipliers, i+1, i+1+numMatches) {
			multipliers[i+1+j] += currentMultiplier
		}

	}

	fmt.Println("--------SCORES--------")
	fmt.Println(scores)
	fmt.Println("--------MULTIPLIER----")
	fmt.Println(multipliers)
	fmt.Println("----------------------")

	for _, multi := range multipliers {
		sum += uint64(multi)
	}

	return sum
}
