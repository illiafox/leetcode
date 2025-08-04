package hard

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"()", 2},
		{")(", 0},
		{"(()", 2},
		{")()())", 4},
		{"()(()", 2},

		{"((()))", 6},
		{"()(())", 6},
		{"()()()", 6},

		{"(()())", 6},
		{"(()(((()", 2},
		{")()())()()(", 4},
		{"()(()))))", 6},

		{"(((((((((", 0},
		{"))))))))))", 0},
		{"(()())(()())", 12},

		{"()(()))((()))", 6},
		{"()((())(()))", 12},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := longestValidParentheses(tt.input)
			if got != tt.want {
				t.Errorf("longestValidParentheses(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
