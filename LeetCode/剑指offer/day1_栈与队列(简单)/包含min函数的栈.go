// 题目链接：https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
// 思路：建立辅助栈，与存储数据的栈大小相同，在向栈中存数据时，辅助栈同时存入一个数字
// 该辅助栈入栈元素的序列是一个非严格递增序列
// 如果该数据小于辅助栈栈顶元素，则辅助栈存入该数据，否则辅助栈还存入一个辅助栈的栈顶元素。
//（如果是存入第一个元素，辅助栈直接入栈即可，无比较操作）

package main

type MinStack struct {
	nums []int 	//储存栈
	min []int 	//辅助储存栈，存储最小值
}

/** initialize your data structure here. */
// 为解决命名冲突，这里函数名后+“2”，在LeetCode需删除
func Constructor2() MinStack {
	return MinStack{
		[]int{},
		[]int{},
	}
}

// 入栈时，存储栈直接入栈
// 对辅助栈，若栈长度为0，直接入栈
// 否则，与栈顶元素进行比较，若大于栈顶元素，入栈，否则，辅助栈再次存入栈顶元素
func (this *MinStack) Push(x int)  {
	this.nums=append(this.nums,x)
	if len(this.min)==0{
		this.min=append(this.min,x)
	}else if this.min[len(this.min)-1]<x{
		this.min=append(this.min,this.min[len(this.min)-1])
	}else{
		this.min=append(this.min,x)
	}
}

// 出栈时，两栈均常规出栈即可
func (this *MinStack) Pop()  {
	this.nums=this.nums[:len(this.nums)-1]
	this.min=this.min[:len(this.min)-1]
}

// 返回存储栈栈顶元素
func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

// 求min值时，返回辅助栈栈顶元素即可
func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */