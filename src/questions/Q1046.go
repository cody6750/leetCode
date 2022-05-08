type MaxHeap []int

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	mh := MaxHeap{}
	for _, s := range stones {
		mh.Add(s)
	}
	for len(mh) > 1 {
		mh.Smash()
	}
	if len(mh) == 1 {
		return (mh)[0]
	}
	fmt.Print(mh)
	return 0
}

func (m *MaxHeap) Smash() {
	if len(*m) < 1 {
		return
	}
	i := m.Pop()
	j := m.Pop()

	if i != j {
		m.Add(i - j)
	}
}

func (m *MaxHeap) Add(v int) {
	*m = append(*m, v)
	m.HeapifyUp(len(*m) - 1)
}

func (m *MaxHeap) Pop() int {
	if len(*m) == 1 {
		head := (*m)[0]
		*m = (*m)[:0]
		return head
	}
	head := (*m)[0]
	m.swap(0, len(*m)-1)
	*m = (*m)[:len(*m)-1]
	m.HeapifyDown(0)
	return head
}

func (m *MaxHeap) HeapifyDown(i int) {
	n := len(*m) - 1
	for {
		l := leftChild(i)
		r := rightChild(i)
		var compareTo int
		if l <= n && r <= n {
			if (*m)[l] > (*m)[r] {
				compareTo = l
			} else {
				compareTo = r
			}
			if (*m)[i] < (*m)[compareTo] {
				m.swap(i, compareTo)
				i = compareTo
				continue
			}
		} else if l <= n {
			if (*m)[i] < (*m)[l] {
				m.swap(i, l)
				i = l
				continue
			}
		} else if r <= n {
			if (*m)[i] < (*m)[r] {
				m.swap(i, r)
				i = r
				continue
			}
		}
		break
	}
}

func (m *MaxHeap) HeapifyUp(v int) {
	for {
		pI := parent(v)
		if (*m)[pI] < (*m)[v] {
			m.swap(v, pI)
			v = pI
			continue
		}
		break
	}
}

func (m *MaxHeap) swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func parent(i int) int {
	return (i - 1) / 2

}
func leftChild(i int) int {
	return i*2 + 1
}
func rightChild(i int) int {
	return i*2 + 2
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

// Find the two stones
// Insert smash value

// [1]

// [1 2 ]

// [2,6,3,5]

// [1,2,3,4,5]

// [9, 1,9,1,9,1]
