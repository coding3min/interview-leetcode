// 题目解析在 用两个栈实现队列.go
// 本文件为 用链表实现的栈 来 实现队列
// 为解决命名冲突，本文件结构体与函数名后+1
package main

import "container/list"

type CQueue1 struct {
	stack1, stack2 *list.List
}

func Constructor1() CQueue1 {
	return CQueue1{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue1) AppendTail1(value int)  {
	this.stack1.PushBack(value)
}

func (this *CQueue1) DeleteHead1() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}