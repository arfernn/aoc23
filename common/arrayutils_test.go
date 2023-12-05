package common

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type test struct {
		inputA []string
		inputB []string
		want   []string
	}

	tests := []test{{
		inputA: []string{"1", "2", "3", "4", "5", "6", "7"},
		inputB: []string{"2", "11", "20", "48", "3", "3"},
		want:   []string{"2", "3", "3"}}}
	for _, test := range tests {
		output := Intersect(test.inputA, test.inputB)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("ERROR, EXPECTED %s GOT %s", test.want, output)
		}
	}
}
