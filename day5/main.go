package main

import (
	"aoc/common"
	"fmt"
)

func main() {
	input, _ := common.ParseFile("input.txt")

	seeds, maps := ParseFile(input)

	result, err := FindClosestLocation(seeds, maps)
	fmt.Println(err)
	fmt.Println(result)
	fmt.Println(seeds)
	fmt.Println(maps)
}
