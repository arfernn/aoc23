package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func findMinMaxIndexesRegExp(str, substr string) (int, int, error) {
	re := regexp.MustCompile(regexp.QuoteMeta(substr))
	matches := re.FindAllStringIndex(str, -1)
	if len(matches) == 0 {
		return -1, -1, fmt.Errorf("error")
	}
	return matches[0][0], matches[len(matches)-1][0], nil
}

func CountValues(lines []string) int {
	dict := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	maxCnt := 0
	for _, line := range lines {
		leftPt, rightPt := -1, -1
		var leftCh, rightCh string

		for key, val := range dict {

			min, max, err := findMinMaxIndexesRegExp(line, key)

			if err == nil {
				if min < leftPt || leftPt == -1 {
					leftCh = val
					leftPt = min
				}
				if max > rightPt || rightPt == -1 {
					rightCh = val
					rightPt = max
				}
			}
		}
		temp, _ := strconv.Atoi(leftCh + rightCh)

		maxCnt += temp
	}
	return maxCnt
}
