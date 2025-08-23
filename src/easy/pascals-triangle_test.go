package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePascal(t *testing.T) {
	type testCase struct {
		n      int
		expect [][]int
	}

	cases := []testCase{
		{0, [][]int{}},
		{1, [][]int{
			{1},
		}},
		{2, [][]int{
			{1},
			{1, 1},
		}},
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for _, c := range cases {
		got := generatePascalTriangle(c.n)
		assert.Equal(t, c.expect, got, "numRows=%d", c.n)
	}
}
