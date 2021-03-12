/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (45.51%)
 * Likes:    871
 * Dislikes: 0
 * Total Accepted:    200.2K
 * Total Submissions: 439.7K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 * 
 * 示例 1:
 * 
 * 输入: 1->2
 * 输出: false
 * 
 * 示例 2:
 * 
 * 输入: 1->2->2->1
 * 输出: true
 * 
 * 
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
 // 慢是慢了一点，但是比较清晰
func isPalindrome(head *ListNode) bool {
	quick,slow := head,head
	var beforeNode *ListNode 
	for quick!=nil && slow!=nil{
		quick = quick.Next
		if quick!=nil{
			quick = quick.Next
		}
        beforeNode = slow
		slow = slow.Next
	}
	// 奇数的情况
	if quick!=nil{
		slow = slow.Next
	}
	newHead := revertTree(slow)
	oldTree,newTree := head,newHead
	isPalindromeBool := true
	for  isPalindromeBool && newTree!=nil{
		if oldTree.Val!=newTree.Val{
			isPalindromeBool = false
		}
		oldTree = oldTree.Next
		newTree = newTree.Next
	}
	newHead = revertTree(newHead)
	beforeNode.Next = newHead
	return isPalindromeBool
}
func revertTree(head *ListNode) *ListNode{
	if head == nil || head.Next == nil{
		return head
	}
	newHead := revertTree(head.Next)
    head.Next.Next = head
	head.Next = nil
	return newHead
}
// 转换成数组再判断回文
// func isPalindrome(head *ListNode) bool {
// 	var res []int
// 	for p:=head;p!=nil;p=p.Next{
// 		res = append(res,p.Val)
// 	}
// 	for i,j:=0,len(res)-1;i<=j;{
// 		if res[i]!=res[j]{
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }
// @lc code=end

