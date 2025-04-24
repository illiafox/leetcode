package easy

// XOR
// a | b | a ^ b
// --|---|------
// 0 | 0 | 0
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 0
func singleNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum ^= nums[i]
	}

	return sum
}
