//题目链接：https://leetcode.cn/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

package main

//这道题比较通俗的解法是两次遍历
//先一次遍历链表，得到链表的长度 n，倒数第 k 个节点，即正数第 n-k+1 个节点，第二次遍历输出该节点即可。
func getKthFromEnd(head *ListNode, k int) *ListNode {
	n := 0
	cur := head
	// 第一次遍历得到链表长度n
	for cur != nil{
		n ++
		cur = cur.Next
	}
	cur = head
	x := n-k+1
	//第二次遍历找到整数第n-k+1个节点并返回
	for i:=0;i<x-1;i++{
		cur = cur.Next
	}
	return cur
}

//方法2：快慢指针
//使用快慢指针后一次遍历即可得到倒数第 k 个节点
//快慢指针 fast 和 slow 初始化指向 head 节点，fast 指针先走 k 步，此时 fast 和 slow 之间距离为 k，然后两指针同时向前走，
//当 fast 指向链表尾结点的Next域（即空节点）时，slow 指向倒数第 k 个节点，返回 slow 即可。
func getKthFromEnd_2(head *ListNode, k int) *ListNode {
	fast,slow := head,head
	for i:=0;i<k;i++{
		fast = fast.Next
	}
	for fast != nil{
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}