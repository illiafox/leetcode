package easy

// old solution
func findDifference(nums1 []int, nums2 []int) [][]int {
	diff := make(map[int][2]bool)

	traverse := func(arr []int, idx int) {
		for _, n := range arr {
			d := diff[n]
			d[idx] = true
			diff[n] = d
		}
	}

	traverse(nums1, 0)
	traverse(nums2, 1)

	var answer0 []int
	var answer1 []int

	for k, v := range diff {
		if v[0] && v[1] {
			continue
		}
		if v[1] {
			answer0 = append(answer0, k)
		} else {
			answer1 = append(answer1, k)
		}
	}

	return [][]int{answer1, answer0}
}
