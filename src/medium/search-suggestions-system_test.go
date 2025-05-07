package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	type testCase struct {
		products   []string
		searchWord string
		expected   [][]string
	}

	var cases = []testCase{
		{
			[]string{"havana"},
			"havana",
			[][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}},
		},
	}

	for _, c := range cases {
		assert.Equal(t, suggestedProducts(c.products, c.searchWord), c.expected)
	}
}
