package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimizeXOR(t *testing.T) {
	type testCase struct {
		num1, num2 int
		expected   int
	}

	var cases = []testCase{
		{
			3, 5, 3,
		},
		{
			1, 12, 3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, minimizeXor(c.num1, c.num2), c.expected)
	}
}
