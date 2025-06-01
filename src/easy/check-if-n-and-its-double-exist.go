package easy

func checkIfExist(arr []int) bool {
	doubles := make(map[int]struct{})

	for _, v := range arr {
		if _, ok := doubles[v*2]; ok {
			return true
		}
		if v%2 == 0 {
			if _, ok := doubles[v/2]; ok {
				return true
			}
		}

		doubles[v] = struct{}{}
	}

	return false
}
