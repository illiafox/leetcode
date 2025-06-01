package easy

import "slices"

func targetIndicesSort(nums []int, target int) (out []int) {
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			out = append(out, i)
		}
	}

	return out
}
