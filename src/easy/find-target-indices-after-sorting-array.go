package easy

import "slices"

func targetIndicesSort(nums []int, target int) (out []int) {
	slices.Sort(nums)
	for i := range nums {
		if nums[i] == target {
			out = append(out, i)
		}
	}

	return out
}
