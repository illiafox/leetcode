package medium

import "testing"

// goos: darwin
// goarch: arm64
// pkg: solutions/src/medium
// cpu: Apple M1 Pro
// BenchmarkMaximumLengthSuffix
// BenchmarkMaximumLengthSuffix-10    	    2530	    466290 ns/op
// BenchmarkMaximumLength
// BenchmarkMaximumLength-10          	    3967	    301453 ns/op

func cloneSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func TestMaximumLength(t *testing.T) {
	functions := map[string]func([]int) int{
		"Final":  maximumLength,
		"Suffix": maximumLengthSuffix,
	}

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1", []int{5, 4, 1, 2, 2}, 3},
		{"example 2", []int{1, 3, 2, 4}, 1},
		{"all ones even", []int{1, 1, 1, 1}, 3},
		{"all ones odd", []int{1, 1, 1, 1, 1}, 5},
		{"single element", []int{10}, 1},
		{"long chain [2,4,16]", []int{2, 2, 4, 4, 16, 16, 16}, 5},
		{"root with 1 count broken", []int{2, 4, 1, 1, 4, 16, 14, 20}, 3},
		{"no valid chains bigger than 1", []int{2, 3, 5, 7}, 1},
		{"large numbers no chain", []int{1000000000, 999999999}, 1},
	}

	for name, fn := range functions {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					inputCopy := cloneSlice(tc.input)

					actual := fn(inputCopy)
					if actual != tc.expected {
						t.Errorf("Expected %d, got %d for input %v", tc.expected, actual, tc.input)
					}
				})
			}
		})
	}
}

func generateLargeInput() []int {
	const size = 5000

	input := make([]int, 0, size)

	for range size {
		// mix of chains, ones, and random large elements
		input = append(input, 1, 2, 2, 4, 4, 16, 16, 256, 3, 3, 9)
	}

	return input
}

func BenchmarkMaximumLengthSuffix(b *testing.B) {
	input := generateLargeInput()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inputCopy := cloneSlice(input)
		b.StartTimer()

		_ = maximumLengthSuffix(inputCopy)
	}
}

func BenchmarkMaximumLength(b *testing.B) {
	input := generateLargeInput()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = maximumLength(input)
	}
}
