package questions

func plusOne(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		var i int
		for len(digits)-1-i >= 0 {
			if len(digits)-1-i == 0 && digits[0] == 9 {
				digits[0] = 1
				digits = append(digits, 0)
				break
			} else if digits[len(digits)-1-i] == 9 {
				digits[len(digits)-1-i] = 0
				i++
			} else {
				digits[len(digits)-1-i]++
				break
			}
		}
	} else {
		digits[len(digits)-1]++
	}
	return digits
}
