package medium

import "testing"

type testCase struct {
	initWords []string
	cases     []struct {
		word     string
		expected bool
	}
}

func TestMagicDictionary_Table(t *testing.T) {
	tests := []testCase{
		{
			initWords: []string{"hello", "leetcode"},
			cases: []struct {
				word     string
				expected bool
			}{
				{"hello", false},
				{"hhllo", true},
				{"hell", false},
				{"leetcoded", false},
			},
		},
		{
			initWords: []string{"hello", "hallo", "leetcode"},
			cases: []struct {
				word     string
				expected bool
			}{
				{"hello", true},
				{"hhllo", true},
				{"hell", false},
				{"leetcoded", false},
			},
		},
		{
			initWords: []string{"a", "b"},
			cases: []struct {
				word     string
				expected bool
			}{
				{"a", true},
				{"b", true},
				{"c", true},
				{"d", true},
				{"e", true},
				{"f", true},
				{"ab", false},
				{"ba", false},
			},
		},
	}

	for ti, tc := range tests {
		d := MagicDictionaryConstructor()
		d.BuildDict(tc.initWords)
		for ci, c := range tc.cases {
			got := d.Search(c.word)
			if got != c.expected {
				t.Fatalf("test %d case %d: Search(%q)=%v, want %v", ti+1, ci+1, c.word, got, c.expected)
			}
		}
	}
}
