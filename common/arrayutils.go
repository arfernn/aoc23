package common

import (
	"fmt"
	"strconv"
)

func Intersect(stringsA []string, stringsB []string) []string {

	fakeSet := make(map[string]bool)

	for _, str := range stringsA {
		if str != "" {
			fakeSet[str] = false
		}
	}

	result := make([]string, 0)
	for _, str := range stringsB {
		_, ok := fakeSet[str]
		if ok {
			result = append(result, str)
		}
	}

	return result

}

func SafeSlice(arr []uint64, i, j int) []uint64 {
	if i > len(arr) {
		return nil // Return nil if i is out of bounds
	}

	if j > len(arr) || j < 0 {
		j = len(arr) // Adjust j to the length of the array if it's out of bounds
	}

	return arr[i:j]
}

func StrArrayToIntArray(input []string) []int {
	var t2 = []int{}

	for _, i := range input {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}

func FindMin(input []int) (int, error) {
	if len(input) == 0 {
		return -1, fmt.Errorf("Empty array provided")
	}

	minVal := input[0]

	for _, val := range input {
		if val < minVal {
			minVal = val
		}
	}

	return minVal, nil
}
