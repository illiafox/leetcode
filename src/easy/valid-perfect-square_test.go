package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerfectSquare(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{16, true},
		{14, false},
	}

	for _, testCase := range cases {
		assert.Equal(t, isPerfectSquare(testCase.input), testCase.expected)
		assert.Equal(t, isPerfectSquareMath(testCase.input), testCase.expected)
	}
}
