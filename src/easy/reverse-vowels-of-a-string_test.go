package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	var cases = [][2]string{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede1"},
	}

	for _, testCase := range cases {
		assert.Equal(t, reverseVowels(testCase[0]), testCase[1])
	}
}
