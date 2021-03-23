/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (64.42%)
 * Likes:    987
 * Dislikes: 0
 * Total Accepted:    150.7K
 * Total Submissions: 234K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 进阶：
 * 
 * 
 * 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = [1,2,3,4,5], k = 1
 * 输出：[1,2,3,4,5]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：head = [1], k = 1
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 列表中节点的数量在范围 sz 内
 * 1 
 * 0 
 * 1 
 * 
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
 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 这一题用递归做是最简单的，拆分成翻转链表的题目，只是多一个k的状态维护，遍历到k个节点的时候翻转这部分就好了
    end := head
	// 遍历k个节点，从1开始是因为初始就有一个节点
    for i := 1; i < k && end!=nil;i++{
        end = end.Next
    }
	// end 正常应该指向最后一个节点，如果提前指向nil说明没有跑到k个，不翻转
    if end==nil{
        return head
    }
	// 保存下一跳的状态，截断链表
    prev := end.Next
    end.Next = nil
	// 翻转链表
    newHead := reverseList(head)
	// 继续翻转后面的
    head.Next = reverseKGroup(prev,k)
    return newHead
}
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
// @lc code=end

