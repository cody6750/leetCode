package questions

import (
	"fmt"
	"log"
)

//Queue .. Last in First out
type Queue struct {
	List []int
}

//Enqueue ..
func (q *Queue) Enqueue(input int) {
	q.List = append(q.List, input)
	fmt.Println(q.List)
}

//Deqeue ...
func (q *Queue) Deqeue() int {
	if len(q.List) == 0 {
		log.Fatalf("Unable to deqeueu, queue is empty")
	}
	firstElementValue := q.List[0]
	q.List = q.List[1:]

	return firstElementValue
}

//Stack .. Last out last in
type Stack struct {
	List []int
}

//Push ..
func (s *Stack) Push(input int) {
	s.List = append(s.List, input)
	fmt.Println(s.List)

}

//Pop ..
func (s *Stack) Pop() int {
	if len(s.List) == 0 {
		log.Fatalf("Unable to Pop, stack is empty")
	}
	fmt.Println(s.List)
	lastElement := len(s.List) - 1
	lastElementValue := s.List[lastElement]
	s.List = s.List[:lastElement]
	return lastElementValue
}
