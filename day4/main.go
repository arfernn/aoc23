package main

import (
	"aoc/common"
	"fmt"
)

func main() {
	input, _ := common.ParseFile("input.txt")
	result := GetAccumulatedScores(input)
	fmt.Println(result)
}
