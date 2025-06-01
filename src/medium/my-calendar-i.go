package medium

type interval struct {
	start int
	end   int
}

// https://leetcode.com/problems/my-calendar-i/
// TODO: binary search?
type MyCalendar struct {
	times []interval
}

func (c *MyCalendar) Book(startTime int, endTime int) bool {
	for _, val := range c.times {
		if max(startTime, val.start) < min(endTime, val.end) {
			return false
		}
	}

	c.times = append(c.times, interval{startTime, endTime})
	return true
}
