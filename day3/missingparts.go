package main

import (
	"aoc/common"
	"fmt"
	"strconv"
	"unicode"
)

type coords struct {
	i int
	j int
}

func HasContiguousSymbol(i int, j int, length int, input []string, symbol rune) (bool, coords) {

	specificComparison := symbol != ' '

	//upper boundaries
	if i > 0 {
		tmp := input[i-1]
		subStr := ""
		if j == 0 {
			subStr = tmp[:(j+length)+1]
		} else if j+length > len(tmp)-1 {
			subStr = tmp[j-1:]
		} else {
			subStr = tmp[j-1 : (j + length + 1)]
		}
		for _, char := range subStr {
			if specificComparison {
				return (char == symbol,
			} else {
				if IsSymbol(char) {
					return true
				}
			}
		}
	}

	//lower boundaries
	if i < len(input)-1 {
		tmp := input[i+1]
		subStr := ""
		if j == 0 {
			subStr = tmp[:(j+length)+1]
		} else if j+length > len(tmp)-1 {
			subStr = tmp[j-1:]
		} else {
			subStr = tmp[j-1 : (j+length)+1]
		}
		for _, char := range subStr {
			if specificComparison {
				return char == symbol
			} else {
				if IsSymbol(char) {
					return true
				}
			}
		}
	}

	//midrow
	tmp := input[i]
	if j > 0 {
		if specificComparison {
			return rune(tmp[j-1]) == symbol
		} else {
			if IsSymbol(rune(tmp[j-1])) {
				return true
			}
		}

	}
	if j < len(tmp)-1-length {
		if specificComparison {
			return rune(tmp[j+length]) == symbol
		} else {
			if IsSymbol(rune(tmp[j+length])) {
				return true
			}
		}
	}

	return false
}

func CalcMissingParts(input []string) int {
	sum := 0
	for i, str := range input {
		nums := common.ExtractNumbers(str)
		for _, val := range nums {
			strVal := strconv.Itoa(val.Num)
			hasContiguous := HasContiguousSymbol(i, val.Column, len(strVal), input, ' ')
			fmt.Printf("%s,%d,%d has symbol %t \n", strVal, i, val, hasContiguous)
			if hasContiguous {
				sum += val.Num
				fmt.Println("adding " + strVal)
			}
		}
	}
	return sum
}

func CalcGearRatios(input []string) int {
	sum := 0
	for i, str := range input {
		nums := common.ExtractNumbers(str)
		for _, val := range nums {
			strVal := strconv.Itoa(val.Num)
			hasContiguous := HasContiguousSymbol(i, val.Column, len(strVal), input, '*')
			fmt.Printf("%s,%d,%d has symbol %t \n", strVal, i, val, hasContiguous)
			if hasContiguous {
				sum += val.Num
				fmt.Println("adding " + strVal)
			}
		}
	}
	return sum

}

func IsSymbol(input rune) bool {
	return (!unicode.IsNumber(input) &&
		!unicode.IsLetter(input) &&
		input != '.')
}
