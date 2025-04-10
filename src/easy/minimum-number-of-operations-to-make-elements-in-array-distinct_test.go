package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumOperations(t *testing.T) {
	var cases = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 2, 3, 3, 5, 7}, 2},
		{[]int{4, 5, 6, 4, 4}, 2},
		{[]int{6, 7, 8, 9}, 0},
	}

	for _, testCase := range cases {
		assert.Equal(t, minimumOperations(testCase.input), testCase.expected)
	}
}
