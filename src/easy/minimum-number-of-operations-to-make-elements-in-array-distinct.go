package easy

func minimumOperations(nums []int) int {
	count := 0

	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	for len(nums) > 0 {
		if len(m) == len(nums) {
			return count
		}
		count++

		if len(nums) < 3 {
			return count
		}

		for i := 0; i <= 2; i++ {
			m[nums[i]]--
			if m[nums[i]] == 0 {
				delete(m, nums[i])
			}
		}
		nums = nums[3:]
	}

	return count
}
