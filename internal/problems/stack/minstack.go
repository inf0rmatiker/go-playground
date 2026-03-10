package stack

const maxInt = int(^uint(0) >> 1)

type MinStack struct {
	minVals []int
	data    []int
}

func Constructor() MinStack {
	return MinStack{
		minVals: []int{},
		data:    []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minVals) == 0 || val <= this.minVals[len(this.minVals)-1] {
		this.minVals = append(this.minVals, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		val := this.Top()
		if len(this.minVals) > 0 && val == this.minVals[len(this.minVals)-1] {
			this.minVals = this.minVals[:len(this.minVals)-1]
		}
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minVals) == 0 {
		return maxInt
	}
	return this.minVals[len(this.minVals)-1]
}
