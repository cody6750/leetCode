package questions

import "log"

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
