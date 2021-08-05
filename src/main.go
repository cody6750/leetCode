package main

import (
	"fmt"

	"questions"
)

func main() {
	testCases := []int{2, 7, 11, 15}
	fmt.Printf("Q1 Two sum is %v", questions.TwoSum(testCases, 9))
	fmt.Printf("Longest length is %v", questions.LengthOfLongestSubstring("abcabcbb"))

}
