package main

import (
	"aoc/common"
	"fmt"
)

func main() {
	input, _ := common.ParseFile("input.txt")

	fmt.Println(CalcMissingParts(input))
}
