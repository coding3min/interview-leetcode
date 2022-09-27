//题目链接：https://leetcode.cn/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/?envType=study-plan&id=lcof

package main

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//很经典的链表题目了，计算机考研算法必做题
//我的思路是，新建一个最终返回链表的头结点 head，Val域为0（随便设），Next域为空，之后的链表合并操作在其 Next域进行，最终返回 head.Next 即可。
//再声明一个变量 cur = head，用于合并链表，具体分为两步：
//1. 循环合并：当 l1 与 l2 链表头结点均非空时，判断两者头结点的Val域，谁更小，cur.Next 指向该节点，
//   然后该链表头节点指向其下一个节点，cur也指向其Next域，如此循环，直至 l1 或 l2 为空；
//2. 合并剩余尾部：此时若 l1 链表非空，则 head 指向的链表所有节点值均小于等于 l1 链表剩余节点值，cur.Next 指向 l1 即可；
//   若是 l2 链表非空，同理。

//最终返回 head.Next 即可。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0,nil}
	cur := head
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l1 != nil{
		cur.Next = l1
	}
	if l2 != nil{
		cur.Next = l2
	}
	return head.Next
}

//这道题还可以递归解决：
//当 l1 或 l2 链表为空时，无需合并，直接返回另一个非空链表即可。
//否则，判断 l1 和 l2 链表的头结点谁更小，改变其 Next 域，递归地指向 l1 和 l2.Next 合并链表的结果，最后返回该链表头结点。递归结束条件为某一链表为空。
func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val <= l2.Val{
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}