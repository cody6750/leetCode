//Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
package questions

import (
	"fmt"
)

func LengthOfLongestSubstring(s string) int {

	var longestLength int
	var count int
	if len(s) == 1 || len(s) == 0 {
		fmt.Print("Exit early")
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("Letter is %v\n", string(s[i]))
		count = 1
		usedLetters := make(map[string]int)
		usedLetters[string(s[i])] = 1
		for j := i + 1; j < len(s); j++ {
			fmt.Printf("2nd Letter is %v\n", string(s[j]))
			if usedLetters[string(s[j])] == 0 {
				usedLetters[string(s[j])] = 1
				count++
			} else {
				break
			}
		}
		if count > longestLength {
			longestLength = count
		}
	}
	return longestLength

}

func LengthOfLongestSubstrin1g21(s string) int {

	var longestLength int
	var count int
	if len(s) == 1 || len(s) == 0 {
		fmt.Print("Exit early")
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("Letter is %v\n", string(s[i]))
		count = 1
		usedLetters := make(map[string]int)
		usedLetters[string(s[i])] = 1
		for j := i + 1; j < len(s); j++ {
			fmt.Printf("2nd Letter is %v\n", string(s[j]))
			if usedLetters[string(s[j])] == 0 {
				usedLetters[string(s[j])] = 1
				count++
			} else {
				break
			}
		}
		if count > longestLength {
			longestLength = count
		}
	}
	return longestLength

}
