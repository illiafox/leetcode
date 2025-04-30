package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	type testCase struct {
		n        int
		expected [][]int
	}

	var cases = []testCase{
		{
			3,
			[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			1,
			[][]int{{1}},
		},
	}

	for _, c := range cases {
		assert.Equal(t, generateMatrix(c.n), c.expected)
	}
}
