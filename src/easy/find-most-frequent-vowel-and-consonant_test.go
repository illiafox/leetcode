package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFreqSum(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"successes", 6},
		{"aeiaeia", 3},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, maxFreqSum(tc.s),
			"input = %v", tc.s,
		)
	}
}
