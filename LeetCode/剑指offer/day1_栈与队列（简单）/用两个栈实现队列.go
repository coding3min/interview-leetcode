// 题目链接：https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/?envType=study-plan&id=lcof
// day1/31
// 第一天主题为：栈与队列（简单）
// 包含两道题目：
// 剑指offer09：用两个栈实现队列
// 剑指offer30.包含min函数的栈

//这题太经典了，两个栈实现队列。思路：两个栈，一个用来入栈（支持插入操作），一个用来出栈（支持删除操作），
//栈的特点是先进后出，队列是先进先出，不考虑中间的出入，那一批数据通过栈和队列后的顺序的刚好相反的。我
//们需要将栈的先进后出再倒一次顺序，就和队列的顺序相同了。
//
//官方描述：根据栈先进后出的特性，我们每次往第一个栈里插入元素后，栈顶是最后插入的元素，栈底是下一个待删除的元素。
//为了维护队列先进先出的特性，我们引入第二个栈，用第二个栈维护待删除的元素，在执行删除操作的时候我们首先看下第二个栈是否为空。
//如果为空，我们将第一个栈里的元素一个个弹出插入到第二个栈里，
//这样第二个栈里元素的顺序就是待删除的元素的顺序，要执行删除操作的时候我们直接弹出第二个栈的元素返回即可。

//具体实现：维护两个栈stack1和stack2，其中，stack1用来入队，stack2用于出队，初始时，两个栈均为空，插入元素时，stack1 插入元素即可
//删除元素时，若stack2为空，将stack1所有元素出栈至stack2，若stack2仍然为空，返回-1，否则，从stack2出栈一个元素并返回。
package main

// stack1与stack2分别用于入队和出队
type CQueue struct {
	stack1 []int
	stack2 []int
}

// 构造函数，初始时，两个栈均为空
func Constructor() CQueue {
	return CQueue{[]int{},[]int{}}
}

// stack1用于入队
func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1,value)
}

// stack2用于出队
func (this *CQueue) DeleteHead() int {
	// 若stack2长度为0，将stack1所有元素出栈至stack2
	if len(this.stack2) == 0{
		for len(this.stack1) > 0{
			x := this.stack1[len(this.stack1)-1]
			this.stack2 = append(this.stack2,x)
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	// stack2出栈的元素即为队首元素
	if len(this.stack2) > 0{
		res := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return res
	}
	// 若stack2长度仍为0，说明队列为空，返回-1
	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */