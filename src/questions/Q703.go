type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	K := KthLargest{k: k}
	for _, num := range nums {
		K.Add(num)
	}
	return K
}

func (this *KthLargest) Add(val int) int {
	this.HeapifyUp(val)
	if len(this.nums) > this.k {
		this.Pop()
	}
	return this.nums[0]
}

func (this *KthLargest) Pop() {
	if len(this.nums) == 0 {
		return
	}
	if len(this.nums) == 1 {
		this.nums = this.nums[:0]
		return
	}
	last := len(this.nums) - 1
	this.nums[0], this.nums[last] = this.nums[last], this.nums[0]
	this.nums = this.nums[:last]
	this.HeapifyDown(0)
}

func (this *KthLargest) HeapifyDown(i int) {
	numsLength := len(this.nums) - 1
	for i <= numsLength {
		leftIndex := leftChild(i)
		rightIndex := rightChild(i)
		if leftIndex <= numsLength && rightIndex <= numsLength {
			var indexToCompare int
			if this.nums[leftIndex] <= this.nums[rightIndex] {
				indexToCompare = leftIndex
			} else {
				indexToCompare = rightIndex
			}
			if this.nums[indexToCompare] <= this.nums[i] {
				this.nums[indexToCompare], this.nums[i] = this.nums[i], this.nums[indexToCompare]
				i = indexToCompare
				continue
			}

		} else if leftIndex <= numsLength {
			if this.nums[leftIndex] <= this.nums[i] {
				this.nums[leftIndex], this.nums[i] = this.nums[i], this.nums[leftIndex]
				i = leftIndex
			}
		} else if rightIndex <= numsLength {
			if this.nums[rightIndex] < this.nums[i] {
				i = rightIndex
			}
		}

		break
	}
}

func leftChild(parent int) int {
	return parent*2 + 1
}

func rightChild(parent int) int {
	return parent*2 + 2
}

func (this *KthLargest) HeapifyUp(val int) {
	this.nums = append(this.nums, val)
	valIndex := len(this.nums) - 1
	if len(this.nums) == 1 {
		return
	}
	for valIndex != 0 {
		if val < this.nums[parent(valIndex)] {
			this.nums[parent(valIndex)], this.nums[valIndex] = this.nums[valIndex], this.nums[parent(valIndex)]
			valIndex = parent(valIndex)
		} else {
			break
		}
	}
}

func parent(index int) int {
	return index / 2
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */