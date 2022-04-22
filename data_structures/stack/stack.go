package datastructures

import "log"

/* Stack
Description:
- Linear datastructure
- LIFO. Last one in, is first to go
*/
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	off := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]
	return off
}

func Run() {
	s := Stack{}
	s.Push(2)
	s.Push(1)
	s.Push(0)
	log.Print(s)
	s.Pop()
	log.Print(s)

}
