/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode-cn.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (78.04%)
 * Likes:    771
 * Dislikes: 0
 * Total Accepted:    193.5K
 * Total Submissions: 247.9K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 翻转一棵二叉树。
 * 
 * 示例：
 * 
 * 输入：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 * 
 * 输出：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 * 
 * 备注:
 * 这个问题是受到 Max Howell 的 原问题 启发的 ：
 * 
 * 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
// 非递归，不推荐
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil{
// 		return root
// 	}
// 	stack := []*TreeNode{root}
// 	stackBefore := []*TreeNode{} 
// 	for len(stack)>0{
// 		n := len(stack)
// 		for i:=0;i<n;i++{
// 			node := stack[0]
// 			stack = stack[1:]
// 			if node == nil{
// 				continue
// 			}
// 			stackBefore = append([]*TreeNode{node},stackBefore...)
// 			stack = append(stack,node.Left)
// 			stack = append(stack,node.Right)
// 		}
// 		for i:=0;i<len(stackBefore);i++{
// 			stackBefore[i].Left = stack[len(stack)-1-i*2]
// 			stackBefore[i].Right = stack[len(stack)-1-i*2-1]
// 		}
// 	}
// 	return root
// }
// @lc code=end

