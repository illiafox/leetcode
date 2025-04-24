package easy

func removeElement(nums []int, val int) int {
	i := 0
	for _, n := range nums {
		if n != val {
			nums[i] = n
			i++
		}
	}

	return i
}
