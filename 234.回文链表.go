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
func isPalindrome(head *ListNode) bool {
	var res []int
	for p:=head;p!=nil;p=p.Next{
		res = append(res,p.Val)
	}
	for i,j:=0,len(res)-1;i<=j;{
		if res[i]!=res[j]{
			return false
		}
		i++
		j--
	}
	return true
}
// @lc code=end

