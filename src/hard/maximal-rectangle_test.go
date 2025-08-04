package hard

import (
	"fmt"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{
			matrix: [][]byte{},
			want:   0,
		},
		{
			matrix: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			matrix: [][]byte{
				{'0'},
			},
			want: 0,
		},
		{
			matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 4,
		},
		{
			matrix: [][]byte{
				{'1', '1', '1', '1'},
			},
			want: 4,
		},
		{
			matrix: [][]byte{
				{'1'},
				{'1'},
				{'1'},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.matrix), func(t *testing.T) {
			got := maximalRectangle(tt.matrix)
			if got != tt.want {
				t.Errorf("maximalRectangle() = %d, want %d", got, tt.want)
			}
		})
	}
}
