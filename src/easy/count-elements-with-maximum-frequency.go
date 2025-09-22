package easy

// https://leetcode.com/problems/count-elements-with-maximum-frequency
func maxFrequencyElements(nums []int) int {
	c := make(map[int]int, len(nums))
	maxFreq, count := 0, 0
	for _, x := range nums {
		c[x]++
		if c[x] > maxFreq {
			maxFreq = c[x]
			count = 1
		} else if c[x] == maxFreq {
			count++
		}
	}
	return maxFreq * count
}
