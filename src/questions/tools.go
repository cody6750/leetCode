package questions

import (
	"fmt"
	"log"
)

//Queue .. Last in First out
type Queue struct {
	list []int
}

//Enqueue ..
func (q Queue) Enqueue(input int) {
	q.list = append(q.list, input)
}

//Deqeue ...
func (q Queue) Deqeue() int {
	if len(q.list) == 0 {
		log.Fatalf("Unable to deqeueu, queue is empty")
	}
	firstElementValue := q.list[0]
	q.list = q.list[1:]
	return firstElementValue
}

//Stack .. Last out last in
type Stack struct {
	list []int
}

//Push ..
func (s Stack) Push(input int) {
	s.list = append(s.list, input)
}

//Pop ..
func (s Stack) Pop() int {
	if len(s.list) == 0 {
		log.Fatalf("Unable to Pop, stack is empty")
	}
	lastElement := len(s.list) - 1
	lastElementValue := s.list[lastElement]
	s.list = s.list[:lastElement]
	return lastElementValue
}

func LengthOfLongestSubstrin1g(s string) int {

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
