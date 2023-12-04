package main

import (
	"fmt"
	"testing"
)

func TestHasContigousSymbol(t *testing.T) {
	type test struct {
		input  []string
		i      int
		j      int
		length int
		want   bool
	}

	inputTest := []string{"....*......97..839.......308.........*....115....512..........................497.392.....563........=...............448...........462*884.."}
	tests := []test{
		{input: inputTest, i: 0, j: 131, length: 3, want: true},
	}

	for _, test := range tests {

		result := HasContiguousSymbol(test.i, test.j, test.length, test.input)
		if result != test.want {
			t.Errorf("ERROR %d,%d, got %t, expected %t", test.i, test.j, result, test.want)
		}
	}
}

func TestSymbol(t *testing.T) {
	type test struct {
		input rune
		want  bool
	}

	tests := []test{
		{input: '.', want: false},
		{input: '*', want: true},
		{input: '*', want: true},
		{input: '&', want: true},
		{input: '-', want: true},
		{input: '=', want: true},
		{input: '+', want: true},
		{input: '3', want: false},
		{input: 'a', want: false},
	}

	for _, test := range tests {
		if IsSymbol(test.input) != test.want {
			fmt.Printf("%c RESULT %t\n", test.input, IsSymbol(test.input))
			t.Errorf("unexpected result for %c", test.input)
		}
	}
}

func TestCalcMissingParts(t *testing.T) {
	type test struct {
		input []string
		want  int
	}
	inputTest := []string{
		".......................*......*",
		"...910*...............233..189.",
		"2......391.....789*............",
		"...................983.........",
		"0........106-...............226",
		".%............................$",
		"...*......$812......812..851...",
		".99.711.............+.....*....",
		"...........................113.",
		"28*.....411....%...............",
	}
	tests := []test{{input: inputTest, want: 7253}}
	for _, test := range tests {
		output := CalcMissingParts(inputTest)
		if output != test.want {
			t.Errorf("ERROR, EXPECTED %d GOT %d", test.want, output)
		}
	}
}
