package easy

func missingNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
		sum -= len(nums) - i
	}

	return -sum
}
