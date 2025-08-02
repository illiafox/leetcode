package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRansomNote(t *testing.T) {
	cases := []struct {
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
