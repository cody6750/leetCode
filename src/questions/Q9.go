package questions

func IsPalindrome(x int) bool {
	var isPalindromNumb bool
	var splitInput []int
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x < 10 {
		return true
	}

	for x != 0 {
		splitNumber := x % 10
		splitInput = append(splitInput, splitNumber)
		x = x / 10
	}
	isPalindromNumb = true
	for i := 0; i < len(splitInput)/2; i++ {
		if splitInput[i] != splitInput[len(splitInput)-1-i] {
			return false
		}
	}
	return isPalindromNumb
}
