package easy

// https://leetcode.com/problems/water-bottles
func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles == 0 || numExchange <= 1 {
		return 0
	}
	return numBottles + (numBottles-1)/(numExchange-1)
}
