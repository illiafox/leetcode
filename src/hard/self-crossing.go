package hard

// https://leetcode.com/problems/self-crossing/submissions/1727261324/
func isSelfCrossing(distance []int) bool {
	if len(distance) < 4 {
		return false
	}

	x, y := 0, 0

	var lines []line

	for i := 0; i < len(distance); i++ {
		l := line{
			x1: x,
			y1: y,
			x2: 0,
			y2: 0,
		}
		switch i % 4 {
		case 0:
			y += distance[i]
		case 1:
			x -= distance[i]
		case 2:
			y -= distance[i]
		case 3:
			x += distance[i]
		}
		l.x2, l.y2 = x, y
		l.normalize()

		lines = append(lines, l)

		// we have spiral pattern, self-crossings can only happen with lines from 3, 4 or 5 steps ago
		// traverse len(lines) - 4 .. len(lines)-6
		for j := len(lines) - 4; j >= 0 && j >= len(lines)-6; j-- {
			if lines[j].hasIntersection(l) {
				return true
			}
		}
	}

	return false
}

type line struct {
	x1, y1 int
	x2, y2 int
}

func (a *line) normalize() {
	a.x1, a.x2 = min(a.x1, a.x2), max(a.x1, a.x2)
	a.y1, a.y2 = min(a.y1, a.y2), max(a.y1, a.y2)
}

func (a line) hasIntersection(b line) bool {
	if a.x1 == a.x2 && b.y1 == b.y2 {
		// a is vertical, b is horizontal
		return inRange(a.x1, b.x1, b.x2) && inRange(b.y1, a.y1, a.y2)
	}

	if a.y1 == a.y2 && b.x1 == b.x2 {
		// a is horizontal, b is vertical
		return inRange(b.x1, a.x1, a.x2) && inRange(a.y1, b.y1, b.y2)
	}

	// overlapping horizontal lines
	if a.y1 == a.y2 && b.y1 == b.y2 && a.y1 == b.y1 {
		return overlap(a.x1, a.x2, b.x1, b.x2)
	}

	// overlapping vertical lines
	if a.x1 == a.x2 && b.x1 == b.x2 && a.x1 == b.x1 {
		return overlap(a.y1, a.y2, b.y1, b.y2)
	}

	return false
}

func inRange(val, bound1, bound2 int) bool {
	if bound1 > bound2 {
		bound1, bound2 = bound2, bound1
	}
	return val >= bound1 && val <= bound2
}

func overlap(a1, a2, b1, b2 int) bool {
	if a1 > a2 {
		a1, a2 = a2, a1
	}
	if b1 > b2 {
		b1, b2 = b2, b1
	}
	return max(a1, b1) <= min(a2, b2)
}
