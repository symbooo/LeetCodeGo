package stack

type MinStack struct {
	stack []int
	mini  [][2]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if l := len(this.mini); l == 0 {
		a := [...]int{x, 0}
		this.mini = append(this.mini, a)
	} else {
		if x < this.mini[l-1][0] {
			a := [...]int{x, (len(this.stack) - 1)}
			this.mini = append(this.mini, a)
		}
	}
}

func (this *MinStack) Pop() {
	if l := len(this.stack); l != 0 {
		this.stack = this.stack[:l-1]
		if this.mini[len(this.mini)-1][1] >= l-1 {
			this.mini = this.mini[:len(this.mini)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mini[len(this.mini)-1][0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
