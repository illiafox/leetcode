package easy

func nextGreatestLetter(letters []byte, target byte) byte {
	for l := range letters {
		if letters[l] > target {
			return letters[l]
		}
	}

	return letters[0]
}
