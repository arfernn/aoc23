package main

import (
	"fmt"
	"testing"
)

func TestCountValues(t *testing.T) {
	type test struct {
		input []string
		want  int
	}

	tests := []test{
		{input: []string{"twll1"}, want: 11},
		{input: []string{"eightwothree"}, want: 83},
		{input: []string{"zoni2"}, want: 22},
		{input: []string{"2"}, want: 22},
		{input: []string{""}, want: 0},
		{input: []string{"onetwo"}, want: 12},
		{input: []string{"one"}, want: 11},
		{input: []string{"one", "two", "apsod2nsodv8aaa", "asdf1"}, want: 72},
	}

	for _, test := range tests {
		result := CountValues(test.input)
		fmt.Printf("TEST %d == %d?\n", test.want, result)
		if test.want != result {
			t.Errorf("%d != %d", test.want, result)
		}
	}
}
