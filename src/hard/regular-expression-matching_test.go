package hard

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		str      string
		pattern  string
		expected bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
		{"abc", "a.c", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},

		{"mississippi", "mis*is*p*.", false},
		{"aaa", "a*a", true},
		{"aaa", "ab*a*c*a", true},
		{"aaa", "a*", true},
		{"ab", ".*c", false},
		{"ab", ".*..", true},

		{"", "c*", true},
		{"", ".", false},
		{"a", "", false},
		{"", "", true},

		{"abcd", ".*d", true},
		{"abcd", ".*e", false},
		{"aaa", "aaaa", false},
		{"abc", "abc*", true},
		{"abc", "ab.*", true},
		{"abcd", "d*", false},
	}

	for i, test := range tests {
		t.Run(test.str+"_"+test.pattern, func(t *testing.T) {
			result := isMatch(test.str, test.pattern)
			if result != test.expected {
				t.Errorf("Test %d failed: isMatch(%q, %q) = %v, expected %v",
					i+1, test.str, test.pattern, result, test.expected)
			}
		})
	}
}
