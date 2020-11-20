package arrays

import (
	"fmt"
	"testing"
)

func TestFirstDuplicate(t *testing.T) {
	testArrays := []struct {
		input  []int
		output int
	}{
		{[]int{2, 1, 3, 5, 3, 2}, 3},
		{[]int{2, 2}, 2},
		{[]int{2, 4, 3, 5, 1}, -1},
		{[]int{1}, -1},
		{[]int{5, 5, 5, 5, 5}, 5},
		{[]int{1, 1, 2, 2, 1}, 1},
	}

	var result int
	for _, test := range testArrays {
		result = firstDuplicate(test.input)
		if result != test.output {
			fmt.Println("input: ", test.input)
			t.Errorf("Wanted %d, returned %d ", test.output, result)
		}
	}
}
