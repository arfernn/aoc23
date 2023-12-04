package common

import (
	"strconv"
	"unicode"
)

type Tuple struct {
	Num    int
	Column int
}

func ExtractNumbers(input string) []Tuple {
	isNumber := false
	result := make([]Tuple, 0)
	rPtr, lPtr := 0, 0
	for i, char := range input {

		//		fmt.Println(input[lPtr:rPtr])

		if unicode.IsNumber(char) {
			// number continues
			if isNumber {
				rPtr++
			} else {
				// number begins
				lPtr = i
				rPtr = i
			}

			isNumber = true
		} else {
			if isNumber {
				//number finished
				isNumber = false
				var num int
				if rPtr != lPtr {
					num, _ = strconv.Atoi(input[lPtr : rPtr+1])
				} else {
					num, _ = strconv.Atoi(string(input[lPtr]))
				}
				result = append(result, Tuple{Num: num, Column: lPtr})
			}
		}
	}
	//one last number
	if isNumber {
		var num int
		num, _ = strconv.Atoi(input[lPtr:])
		result = append(result, Tuple{Num: num, Column: lPtr})
	}
	return result
}
