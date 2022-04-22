package datastructures

import "log"

/* Queue
Description:
- Linear datastructure
- FIFO. First one in queue, is first to go
*/
type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return -1
	}
	d := (*q)[0]
	*q = (*q)[1:]
	return d
}

func (q *Queue) Top() int {
	if len(*q) == 0 {
		return -1
	}

	return (*q)[0]
}

func run() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	log.Print(q)
	q.Dequeue()
	log.Print(q)
	q.Dequeue()
	log.Print(q.Top())
}
