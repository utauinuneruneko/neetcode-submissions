type MinStack struct {
	stack []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		minstack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minstack) == 0 || this.minstack[len(this.minstack)-1] >= val {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.minstack[len(this.minstack)-1] == val {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
