package hard

import (
	"reflect"
	"sort"
	"testing"
)

func equalUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	return reflect.DeepEqual(a, b)
}

func TestRemoveInvalidParentheses(t *testing.T) {
	tests := []struct {
		in       string
		expected []string
	}{
		{"()())()", []string{"(())()", "()()()"}},
		{"(a)())()", []string{"(a())()", "(a)()()"}},
		{")(", []string{""}},
		{"", []string{""}},
		{"(())", []string{"(())"}},
		{"()()()", []string{"()()()"}},
		{"abc", []string{"abc"}},
		{"(()", []string{"()"}},
	}

	for _, tc := range tests {
		got := removeInvalidParentheses(tc.in)
		if !equalUnordered(got, tc.expected) {
			t.Errorf("removeInvalidParentheses(%q) = %v; want %v", tc.in, got, tc.expected)
		}
	}
}
