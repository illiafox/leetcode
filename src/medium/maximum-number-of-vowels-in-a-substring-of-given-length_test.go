package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	type testCase struct {
		s        string
		k        int
		expected int
	}

	cases := []testCase{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"weallloveyou", 7, 4},
	}

	for _, c := range cases {
		assert.Equal(t, maxVowels(c.s, c.k), c.expected)
	}
}
