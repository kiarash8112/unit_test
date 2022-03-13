package main

import (
	"testing"
)

func Test_Ints(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		result  int
	}{
		{"number 6", []int{1, 2, 3}, 6},
		{"result 6", []int{4}, 4},
		{"nil result", nil, 2},
	}

	for counter, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if tc.result != s {
				t.Errorf("it is not working %v", counter)
			}
		})
	}
}
