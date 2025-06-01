package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerfectSquare(t *testing.T) {
	var cases = []struct {
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
