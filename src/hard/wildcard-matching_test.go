package hard

import "testing"

func TestIsWildcardMatch(t *testing.T) {
	type testCase struct {
		str      string
		pattern  string
		expected bool
	}

	tests := []testCase{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"", "*", true},
		{"", "?", false},
		{"", "", true},
		{"a", "?", true},
		{"a", "*?", true},
		{"a", "*?*", true},
		{"abefcdgiescdfimde", "ab*cd?i*de", true},

		// Edge cases
		{"abc", "*c", true},
		{"abc", "a*", true},
		{"abc", "?*", true},
		{"abc", "*?*", true},
		{"abc", "*?c", true},
		{"abc", "*?d", false},
	}

	for i, test := range tests {
		t.Run(test.str+"_"+test.pattern, func(t *testing.T) {
			result := wildcardMatch(test.str, test.pattern)
			if result != test.expected {
				t.Errorf("Test %d failed: isMatch(%q, %q) = %v, expected %v",
					i+1, test.str, test.pattern, result, test.expected)
			}
		})
	}

	tests = append(tests, testCase{
		"bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab",
		"b*b*ab**ba*b**b***bba",
		false,
	})

	for i, test := range tests {
		t.Run(test.str+"_"+test.pattern, func(t *testing.T) {
			result := wildcardMatchMemoization(test.str, test.pattern)
			if result != test.expected {
				t.Errorf("Test %d failed: wildcardMatchMemoization(%q, %q) = %v, expected %v",
					i+1, test.str, test.pattern, result, test.expected)
			}
		})
	}
}
