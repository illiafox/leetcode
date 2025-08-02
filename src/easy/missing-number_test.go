package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	type testCase struct {
		nums   []int
		expect int
	}

	cases := []testCase{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, c := range cases {
		assert.Equal(t, missingNumber(c.nums), c.expect)
	}
}
