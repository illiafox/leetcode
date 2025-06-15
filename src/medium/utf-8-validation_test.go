package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidUTF8(t *testing.T) {
	type testCase struct {
		data     []int
		expected bool
	}

	var cases = []testCase{
		{
			[]int{197, 130, 1},
			true,
		},
		{
			[]int{235, 140, 4},
			false,
		},
		{
			[]int{85, 115, 101, 115, 32, 49, 45, 52, 32, 98, 121, 116, 101, 115, 32, 116, 111, 32, 114, 101, 112, 114, 101, 115, 101, 110, 116, 32, 97, 32, 99, 104, 97, 114, 97, 99, 116, 101, 114, 44, 32, 115, 117, 98, 106, 101, 99, 116, 101, 100, 32, 116, 111, 32, 116, 104, 101, 32, 102, 111, 108, 108, 111, 119, 105, 110, 103, 32, 114, 117, 108, 101, 115, 58, 10},
			true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, validUtf8(c.data), c.expected)
	}
}
