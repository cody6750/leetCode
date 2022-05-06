package main

import "log"

type Heap struct {
	Values []int
}

func (h *Heap) Insert(value int) {
	h.Values = append(h.Values, value)
	h.HeapifyMaxUp(len(h.Values) - 1)
	log.Print(h.Values)
}

func (h *Heap) HeapifyMaxUp(i int) {
	for {
		if h.Values[parent(i)] < h.Values[i] {
			h.swap(parent(i), i)
			i = parent(i)
		} else {
			break
		}
	}
}

func (h *Heap) HeapifyMaxDown(i int) {
	for i <= len(h.Values)-1 {
		if left(i) <= len(h.Values)-1 && right(i) <= len(h.Values)-1 {
			if max(h.Values[left(i)], h.Values[right(i)]) == h.Values[left(i)] {
				h.swap(i, left(i))
				i = left(i)
			} else {
				h.swap(i, right(i))
				i = right(i)
			}
			continue
		} else if left(i) <= len(h.Values)-1 {
			if max(h.Values[left(i)], h.Values[right(i)]) == h.Values[left(i)] {
				h.swap(i, left(i))
				i = left(i)
			}
			return
		} else {
			return
		}
		//Two Children
		//One Children
		//0 Children
	}
}

func (h *Heap) GetMax() int {

	if len(h.Values) == 0 {
		return -1
	}
	// Get top
	max := h.Values[0]
	//Replace top with last value
	h.Values[0] = h.Values[len(h.Values)-1]
	//Delete last value
	h.Values = h.Values[:len(h.Values)-1]
	h.HeapifyMaxDown(0)
	return max
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2

}

// Get parent index given the child
func parent(index int) int {
	return (index - 1) / 2
}

// Get left child index given the parent
func left(index int) int {
	return 2*index + 1
}

// Get right child index given the parent
func right(index int) int {
	return 2*index + 2
}

// Swap indexes
func (h *Heap) swap(i1, i2 int) {
	h.Values[i1], h.Values[i2] = h.Values[i2], h.Values[i1]
}

func initHeap() {
	var h Heap
	h.Insert(10)
	h.Insert(20)
	h.Insert(30)
	h.Insert(70)
	h.Insert(50)
	h.Insert(90)
	h.Insert(100)
	h.Insert(80)
	log.Print(h.GetMax(), h)
	log.Print(h.GetMax(), h)

}

//

//     			5
// 			4        3
//       7    6     4
// 5 4 3 7 6 4
