//题目链接：https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/?envType=study-plan&id=lcof
package main

//有了上一题的基础后，这一题就简简单单了
//和第一题的解析类似，这里我们也可以使用单调队列。把队列看作是第一题的窗口，只不过滑动规则有所差异，
//pop 出队时，窗口的左指针右移，并对我们实现的单调队列进行相应操作，求队列最大值时，只需返回单调队列队首元素即可。
type MaxQueue struct {
	// 存储队列
	q     []int
	// 单调队列
	monoQ []int
}

// 构造函数
func Constructor() MaxQueue {
	return MaxQueue{[]int{},[]int{}}
}

// 取最大值，若队列长度为 0，返回-1，否则返回单调队列队首元素
func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0{
		return -1
	}
	return this.monoQ[0]
}

// 入队，存储队列直接入队
// 单调队列需要判断队尾元素与新元素大小关系
// 若新元素大于队尾元素，出队
// 直至单调队列长度为 0或者 新元素值 小于等于 队尾元素
func (this *MaxQueue) Push_back(value int)  {
	this.q = append(this.q,value)
	for len(this.monoQ)>0 && value>this.monoQ[len(this.monoQ)-1]{
		this.monoQ = this.monoQ[:len(this.monoQ)-1]
	}
	this.monoQ = append(this.monoQ,value)
}

// 出队，先判断存储队列长度，若为 0，返回-1
// 获取存储队列队首元素后，存储队列队首元素出队
// 若该元素等于单调队列队首元素，则单调队列队首元素出队
func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0{
		return -1
	}
	res := this.q[0]
	this.q = this.q[1:]
	if res == this.monoQ[0]{
		this.monoQ = this.monoQ[1:]
	}
	return res
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */