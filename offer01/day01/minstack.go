package day01

type MinStack struct {
	a []int
	b []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		a: []int{},
		b: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.a = append(this.a, x)
	if len(this.b) == 0 || this.b[len(this.b)-1] >= x {
		this.b = append(this.b, x)
	}
}

func (this *MinStack) Pop() {
	v := this.a[len(this.a)-1]
	this.a = this.a[:len(this.a)-1]
	if v == this.b[len(this.b)-1] {
		this.b = this.b[:len(this.b)-1]
	}
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MinStack) Min() int {
	return this.b[len(this.b)-1]
}
