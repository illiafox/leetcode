package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAnagrams(t *testing.T) {
	nilSlice := []string(nil)

	cases := []struct {
		name string
		in   []string
		want []string
	}{
		{
			name: "consecutive anagrams grouped",
			in:   []string{"abba", "baba", "bbaa", "cd", "cd", "dccc"},
			want: []string{"abba", "cd", "dccc"},
		},
		{
			name: "identical words are skipped (treated as same)",
			in:   []string{"a", "a", "b"},
			want: []string{"a", "b"},
		},
		{
			name: "non-adjacent anagrams not removed",
			in:   []string{"ab", "cd", "ba"},
			want: []string{"ab", "cd", "ba"},
		},
		{
			name: "single element",
			in:   []string{"hello"},
			want: []string{"hello"},
		},
		{
			name: "empty input -> nil output",
			in:   []string{},
			want: nilSlice,
		},
		{
			name: "no anagrams",
			in:   []string{"one", "two", "three"},
			want: []string{"one", "two", "three"},
		},
		{
			name: "all anagrams in a run -> keep first only",
			in:   []string{"ab", "ba", "ab", "ba"},
			want: []string{"ab"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeAnagrams(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
