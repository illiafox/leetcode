package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	type testCase struct {
		s      string
		shifts [][]int
		expect string
	}

	cases := []testCase{
		{
			"abc",
			[][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}},
			"ace",
		},
		{
			"iaozjb",
			[][]int{{0, 4, 0}, {0, 2, 1}, {1, 3, 1}, {0, 4, 1}, {4, 4, 1}, {2, 3, 0}, {5, 5, 0}, {3, 3, 0}, {2, 3, 0}, {5, 5, 1}, {5, 5, 1}, {5, 5, 0}},
			"jcoxkb",
		},
		{
			"ksztajnymer",
			[][]int{{4, 9, 1}, {2, 5, 1}, {7, 9, 0}, {6, 10, 0}, {1, 7, 1}, {0, 7, 1}, {0, 3, 0}, {3, 3, 1}, {7, 9, 0}, {7, 7, 0}, {6, 7, 0}, {7, 9, 0}, {3, 9, 0}, {10, 10, 1}, {8, 9, 0}, {6, 9, 1}, {1, 2, 1}, {3, 9, 0}, {8, 10, 1}, {4, 7, 1}, {9, 10, 1}, {8, 9, 0}, {5, 8, 0}, {8, 9, 1}, {0, 7, 1}, {2, 2, 1}, {8, 8, 1}, {3, 7, 0}, {1, 10, 1}, {9, 9, 1}, {4, 9, 0}, {5, 6, 1}, {7, 8, 1}, {8, 9, 0}, {7, 7, 1}, {9, 9, 1}},
			"lwfvdmowicu",
		},
	}

	for _, c := range cases {
		assert.Equal(t, shiftingLetters(c.s, c.shifts), c.expect)
	}
}
