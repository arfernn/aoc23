package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var colorMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func GetValidGamesCount(games []string) int {

	sum := 0
	for _, game := range games {
		fmt.Println(game)
		valid, id := IsValidGame(game)
		if valid {
			sum += id
		}
		fmt.Println(valid)
		fmt.Println(id)
		fmt.Println(sum)
	}
	return sum
}

func GetPowerCount(games []string) int {
	sum := 0
	for _, game := range games {
		sum += GetViableNumbersPower(game)
	}
	return sum
}

// Checks if a game string is valid and returns the game id
func IsValidGame(game string) (bool, int) {
	re := regexp.MustCompile("[0-9]+")
	id, _ := strconv.Atoi(re.FindString(game))

	gameArray := strings.Split(game[strings.Index(game, ":")+1:], ";")
	for _, game := range gameArray {
		colorArray := strings.Split(string(game), ",")
		for _, color := range colorArray {
			tuple := strings.Split(color[1:], " ")
			amount, _ := strconv.Atoi(tuple[0])
			maxAmount := colorMap[tuple[1]]
			if amount > maxAmount {
				return false, id
			}
		}
	}

	return true, id
}

func GetViableNumbersPower(game string) int {
	minRed, minGreen, minBlue := 0, 0, 0
	gameArray := strings.Split(game[strings.Index(game, ":")+1:], ";")
	for _, game := range gameArray {
		colorArray := strings.Split(string(game), ",")
		for _, color := range colorArray {
			tuple := strings.Split(color[1:], " ")
			amount, _ := strconv.Atoi(tuple[0])

			if tuple[1] == "red" && amount > minRed {
				minRed = amount
			} else if tuple[1] == "green" && amount > minGreen {
				minGreen = amount
			} else if tuple[1] == "blue" && amount > minBlue {
				minBlue = amount
			}
		}
	}
	return minRed * minGreen * minBlue
}
