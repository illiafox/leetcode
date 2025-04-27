package easy

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	target := len(nums)/2 + len(nums)%2
	for k, v := range m {
		if v >= target {
			return k
		}
	}

	return -1
}
