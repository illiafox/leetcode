package easy

// https://leetcode.com/problems/power-of-three
func isPowerOfThree(n int) bool {
	// 3^19 = 1162261467, biggest power of 3 for int32
	// any true power of 3 will divide 3^19 evenly
	return n > 0 && 1162261467%n == 0
}
