/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	maxNum := math.MinInt32
	var maxLine  func(root *TreeNode) int
	maxLine = func(root *TreeNode) int{
		if root == nil{
			return 0
		}
		num,l,r := root.Val,0,0
		if root.Left != nil{
			l = maxLine(root.Left)
		}
		if root.Right != nil{
			r = maxLine(root.Right)
		}
		tmp := max(num,max(num+l,num+r))
		maxNum = max(tmp,max(maxNum,num+l+r))
		return tmp
	}
	maxLine(root)
	return maxNum
}
func max(x,y int) int{
	if x > y {
		return x
	}
	return y
}
// @lc code=end

