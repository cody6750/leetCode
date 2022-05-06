type MinStack struct {
	Values []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Values = append(this.Values, val)
}

func (this *MinStack) Pop() {
	this.Values = this.Values[:len(this.Values)-1]
}

func (this *MinStack) Top() int {
	return this.Values[len(this.Values)-1]
}

func (this *MinStack) GetMin() int {
	min := this.Values[0]
	for i := 1; i < len(this.Values); i++ {
		if this.Values[i] < min {
			min = this.Values[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */