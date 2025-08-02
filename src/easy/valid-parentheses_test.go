package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
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
