package easy

func containsDuplicate(nums []int) bool {
	distinct := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := distinct[n]; ok {
			return true
		}
		distinct[n] = struct{}{}
	}

	return false
}
