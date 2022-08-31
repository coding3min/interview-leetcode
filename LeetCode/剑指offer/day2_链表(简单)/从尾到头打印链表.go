// 题目链接：https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/?plan=lcof&plan_progress=klczit3
// day2/31
// 第二天主题为：链表（简单）
// 包含三道题目：
// 剑指offer06.从尾到头打印链表
// 剑指offer30.反转链表
// 剑指offer35.复杂链表的复制

// 解题思路：从尾到头打印链表，最后返回整数切片，也就是说，链表头节点的值为切片最后的元素
// 而链表尾结点的值为切片第一个元素
// 与链表的遍历结合，很容易想到先进后出的解题策略
// 说到”先进后出“，那必然会用到栈，先用栈来解题
// 从头到尾遍历链表，将节点依次入栈
// 遍历结束后，在从栈顶逐个输出节点的值至整数切片即可

package main


//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}


func reversePrint(head *ListNode) []int {
	stack := []*ListNode{}
	res := []int{}
	if head == nil{
		return []int{}
	}
	for head != nil{
		stack = append(stack,head)
		head = head.Next
	}
	for len(stack) > 0{
		res = append(res,stack[len(stack)-1].Val)
		stack = stack[:len(stack)-1]
	}
	return res
}

// 法2：递归在本质上是一个栈结构，针对本题，我们也可以使用递归来实现
// 容易想到，要实现反过来输出链表，每当我们访问到一个节点时，要将
// 当前节点的值放在返回切片的末尾，先递归输出其之后的节点，
// 当访问到当前的节点为空节点时，返回空切片即可
// 利用系统栈，实现了从尾到头打印链表
// 为避免命名冲突，此函数名后添加了后缀 ”_recursion“
func reversePrint_recursion(head *ListNode) []int {
	if head == nil{
		return []int{}
	}
	return append(reversePrint(head.Next),head.Val)
}