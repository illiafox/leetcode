package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCDOfStrings(t *testing.T) {
	cases := [][3]string{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
	}

	for _, testCase := range cases {
		assert.Equal(t, GCDOfStrings(testCase[0], testCase[1]), testCase[2])
	}
}
