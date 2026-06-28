package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name: "Standard maze with path around obstacles",
			grid: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0},
			},
			k:        1,
			expected: 6,
		},
		{
			name: "Blocked maze due to insufficient k",
			grid: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0},
			},
			k:        1,
			expected: -1,
		},
		{
			name: "Large complex maze snake route",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 1, 1, 1, 1, 1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			},
			k:        1,
			expected: 20,
		},
		{
			name: "High k triggering Manhattan shortcut shortcut",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 1},
				{0, 0, 0},
			},
			k:        6,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := shortestPath(tt.grid, tt.k)
			assert.Equal(t, tt.expected, actual,
				"shortestPath(%v, %d) should match expected steps", tt.name, tt.k,
			)
		})
	}
}
