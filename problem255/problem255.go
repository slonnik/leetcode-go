package problem255

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	result := this.Top()
	this.data = this.data[:len(this.data)-1]
	return result
}

func (this *MyStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}
