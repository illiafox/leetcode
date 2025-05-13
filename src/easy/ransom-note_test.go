package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRansomNote(t *testing.T) {
	var cases = []struct {
		ransomNote string
		magazine   string

		expected bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, testCase := range cases {
		assert.Equal(t, canConstruct(testCase.ransomNote, testCase.magazine),
			testCase.expected,
		)
	}
}
