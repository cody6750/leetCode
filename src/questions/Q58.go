package questions

func lengthOfLastWord(s string) int {
	startWord := false
	var startIndex int
	var endIndex int
	var listOfWords []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " && !startWord {
			startWord = true
			startIndex = i
		}

		if len(s)-1 == i && startWord {
			endIndex = i
			startWord = false
			listOfWords = append(listOfWords, s[startIndex:endIndex+1])
			break
		} else if len(s)-1 > i {
			if string(s[i+1]) == " " && startWord {
				endIndex = i
				startWord = false
				listOfWords = append(listOfWords, s[startIndex:endIndex+1])
			}
		}
	}

	return len(listOfWords[len(listOfWords)-1])
}
