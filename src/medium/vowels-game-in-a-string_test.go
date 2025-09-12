package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesAliceWin(t *testing.T) {
	type testCase struct {
		s        string
		expected bool
	}

	cases := []testCase{
		{"leetcoder", true},
		{"bbcd", false},
		{"", false},
		{"bcdfg", false},
		{"apple", true},
		{"rhythmical", true},
		{"zzza", true},
		{"leetcode", true},
		{"bcdfaeiou", true},
		{"hello", true},
	}

	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			assert.Equal(t, c.expected, doesAliceWin(c.s))
		})
	}
}
