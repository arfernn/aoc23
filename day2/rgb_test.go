package main

import (
	"fmt"
	"testing"
)

func TestIsValidGame(t *testing.T) {
	type test struct {
		input     string
		wantedVal bool
		wantedId  int
	}

	tests := []test{
		{input: "Game 1: 1 red, 10 blue, 5 green; 11 blue", wantedVal: true, wantedId: 1},
		{input: "Game 1: 1 red", wantedVal: true, wantedId: 1},
		{input: "Game 4: 1 blue", wantedVal: true, wantedId: 4},
		{input: "Game 55: 5 green; 11 blue; 20 blue", wantedVal: false, wantedId: 55},
		{input: "Game 2: 55 red", wantedVal: false, wantedId: 2},
		{input: "Game 10: 23 red, 10 blue, 5 green; 11 blue", wantedVal: false, wantedId: 10},
	}

	for _, test := range tests {
		val, id := IsValidGame(test.input)
		fmt.Printf("is this valid->%s\n", test.input)
		fmt.Println(val)
		fmt.Println(id)
		if test.wantedVal != val || test.wantedId != id {
			t.Errorf("%s FAILED", test.input)
		}
	}
}

func TestGetValidGamesCount(t *testing.T) {
	test := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := 8

	result := GetValidGamesCount(test)

	if expected != result {
		t.Errorf("%d!=%d", expected, result)
	}
}
