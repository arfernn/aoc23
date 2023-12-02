package main

import (
	"aoc/common"
	"fmt"
)

func main() {
	input, _ := common.ParseFile("input.txt")
	fmt.Println(GetValidGamesCount(input))
	fmt.Println(GetPowerCount(input))
}
