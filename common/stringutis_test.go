package common

import (
	"testing"
)

type test struct {
	input string
	want  []Tuple
}

func TestExtractNumbers(t *testing.T) {
	/*
	   	tests := []test{
	   		{input: ".....aaa2212222bbdca11", want:[]tuple{{2212222: 8}, 11: 20}},
	   		{input: "11", want: map[int]int{11: 0}},
	   		{input: "1.....aaa2bbdca11", want: map[int]int{1: 0, 2: 9, 11: 15}},
	   		{input: "1", want: map[int]int{1: 0}},
	   	}jjj

	   for _, test := range tests {

	   		result := ExtractNumbers(test.input)
	   		if !reflect.DeepEqual(result, test.want) {
	   			t.Errorf("ERROR %d not the same as %d", result, test.want)
	   		}
	   	}
	*/
}
