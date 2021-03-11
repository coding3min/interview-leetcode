/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (38.18%)
 * Likes:    5596
 * Dislikes: 0
 * Total Accepted:    701.9K
 * Total Submissions: 1.8M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 * 
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 * 
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0 
 * 题目数据保证列表表示的数字不含前导零
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    tmp,a,b := 0,0,0
    resArray := []int{}
    for l1!=nil || l2!=nil{
        if l1 != nil{
            a = l1.Val
            l1 = l1.Next
        }else{ 
            a = 0
        }
        if l2 != nil{
            b=l2.Val
            l2=l2.Next
        }else{
            b=0
        }
        add := a+b+tmp
        resArray = append(resArray,add%10)
        tmp = add/10
    }
    
    for tmp!=0{
        resArray = append(resArray,tmp%10)
        tmp = tmp/10
    }

    var head *ListNode
    var beforeNode *ListNode
    for _,v := range resArray{
       curr := &ListNode{
            v,
            nil,
       }
       if head == nil{
           head = curr
           beforeNode = curr
       }else{
           beforeNode.Next = curr
           beforeNode = curr
       }
    }
    return head
}
// @lc code=end

