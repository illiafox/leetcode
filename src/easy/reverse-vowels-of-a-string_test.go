package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	cases := [][2]string{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}

	for _, testCase := range cases {
		assert.Equal(t, reverseVowels(testCase[0]), testCase[1])
	}
}
