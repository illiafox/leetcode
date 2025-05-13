package easy

func canConstruct(ransomNote string, magazine string) bool {
	var buf [26]int

	for i := range magazine {
		buf[magazine[i]-'a'] += 1
	}
	for i := range ransomNote {
		buf[ransomNote[i]-'a'] -= 1
	}

	for i := range buf {
		if buf[i] < 0 {
			return false
		}
	}

	return true
}
