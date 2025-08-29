package hard

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s        string
		words    []string
		expected []int // order doesn't matter
	}{
		{
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
		},
		{
			s:        "barfoo",
			words:    []string{"baz", "qux"},
			expected: nil,
		},
		{
			s:        "foo",
			words:    []string{"foo", "bar"},
			expected: nil,
		},
		{
			s:        "barfoofoobarthefoobarman",
			words:    []string{"bar", "foo", "the"},
			expected: []int{6, 9, 12},
		},
		{
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "word"},
			expected: nil,
		},
		{
			s:        "aaaaaaaaaaaaaa",
			words:    []string{"aa", "aa"},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			s:        "abxabab",
			words:    []string{"ab"},
			expected: []int{0, 3, 5},
		},
		{
			s:        "fooXbarfoo",
			words:    []string{"foo", "bar"},
			expected: []int{4},
		},
		{
			s:        "aaaaaa",
			words:    []string{"aa", "aa", "aa", "aa"},
			expected: nil,
		},
		{
			s:        "foobarbarfoo",
			words:    []string{"foo", "bar"},
			expected: []int{0, 6},
		},
	}

	for _, tc := range tests {
		sort.Ints(tc.expected)

		got := findSubstring(tc.s, tc.words)
		sort.Ints(got)
		assert.Equal(t, tc.expected, got)

		got = findSubstringOld(tc.s, tc.words)
		sort.Ints(got)
		assert.Equal(t, tc.expected, got)
	}
}
