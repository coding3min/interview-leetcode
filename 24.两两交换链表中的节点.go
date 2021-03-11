/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (69.02%)
 * Likes:    803
 * Dislikes: 0
 * Total Accepted:    218.8K
 * Total Submissions: 316.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 * 
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = [1]
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [0, 100] 内
 * 0 
 * 
 * 
 * 
 * 
 * 进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	temp := head.Next
	head.Next = swapPairs(temp.Next)
	temp.Next = head
	return temp
}

// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil{
// 		return head
// 	}
// 	prev,curr := head,head.Next
// 	head = curr
// 	for {
// 		three := curr.Next
// 		// p c 对调，p指向下一个，c指向上一个
// 		prev.Next = three
// 		curr.Next = prev
// 		// 更新p c的值
// 		if three == nil || three.Next == nil{
// 			break
// 		}
// 		// 更新前一个指针
// 		prev.Next = three.Next
// 		prev = three
// 		curr = three.Next
// 	}
// 	return head
// }
// @lc code=end

