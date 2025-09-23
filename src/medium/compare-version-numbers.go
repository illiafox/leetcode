package medium

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/compare-version-numbers
func compareVersion(version1 string, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")

	i, j := 0, 0

	for i < len(ver1) || j < len(ver2) {
		var (
			n1, n2 int
			err    error
		)

		if i < len(ver1) {
			n1, err = strconv.Atoi(ver1[i])
			if err != nil {
				panic(err)
			}
			i++
		}
		if j < len(ver2) {
			n2, err = strconv.Atoi(ver2[j])
			if err != nil {
				panic(err)
			}
			j++
		}

		if n1 > n2 {
			return 1 // greater
		} else if n1 < n2 {
			return -1 // less
		}
	}

	return 0 // equal
}
