package main

import (
	"aoc/common"
	"fmt"
)

func main() {

	entries, err := common.ParseFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := CountValues(entries)

	fmt.Println(result)

}
