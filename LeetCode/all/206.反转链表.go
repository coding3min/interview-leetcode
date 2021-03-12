/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (71.46%)
 * Likes:    1570
 * Dislikes: 0
 * Total Accepted:    463.3K
 * Total Submissions: 648.3K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 * 
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var prev,curr *ListNode
	prev,curr = nil,head
	for curr!=nil{
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
 // 递归方式
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil{
// 		return head
// 	}
// 	newHead := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return newHead
// }
// @lc code=end

