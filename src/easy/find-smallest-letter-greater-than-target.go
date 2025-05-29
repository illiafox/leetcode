package easy

func nextGreatestLetter(letters []byte, target byte) byte {
	for l := 0; l < len(letters); l++ {
		if letters[l] > target {
			return letters[l]
		}
	}

	return letters[0]
}
