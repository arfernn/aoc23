package main

import "testing"

func TestCalculatePrize(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "1 2 3 4 5 6 7 | 2 11 20 48 3 3", want: 4},
		{input: "13 32 20 16 61 | 61 30 68 82 17 32 24 19", want: 2},
		{input: " 1 21 53 59 44 | 69 82 63 72 16 21 14  1", want: 2},
	}

	for _, test := range tests {
		result, err := CalculatePrize(test.input)
		if err != nil {
			t.Error("Function failed")
		}

		if result != test.want {
			t.Errorf("Unexpected result, got %d, expected %d", result, test.want)
		}
	}
}
