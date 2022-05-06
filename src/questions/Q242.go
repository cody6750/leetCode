func isAnagram(s string, t string) bool {
	counter := make(map[string]int)
	if len(s) != len(t) {
		return false
	}
	for index, _ := range s {
		counter[string(s[index])] += 1
		counter[string(t[index])] -= 1
	}
	for _, letter := range s {
		if counter[string(letter)] != 0 {
			return false
		}
	}

	return true
}