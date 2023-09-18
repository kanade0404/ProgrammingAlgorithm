package min_stack

type MinStack struct {
	len int
	min int
	arr []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	this.len += 1
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:this.len-1]
}

func (this *MinStack) Top() int {
	return this.arr[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
