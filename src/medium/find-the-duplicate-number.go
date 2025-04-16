package medium

// https://www.geeksforgeeks.org/floyds-cycle-finding-algorithm/
func findDuplicate(nums []int) int {
	slowPtr := nums[0]
	fastPtr := nums[0]

	for {
		slowPtr = nums[slowPtr]
		fastPtr = nums[nums[fastPtr]]

		if fastPtr == slowPtr {
			break
		}
	}

	slowPtr = nums[0]
	for slowPtr != fastPtr {
		slowPtr = nums[slowPtr]
		fastPtr = nums[fastPtr]
	}

	return slowPtr
}
