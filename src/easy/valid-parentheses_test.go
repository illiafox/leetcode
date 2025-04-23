package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	var cases = []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
	}

	for _, testCase := range cases {
		assert.Equal(t, isValidParentheses(testCase.input), testCase.expected)
	}
}
