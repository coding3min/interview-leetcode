/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}
	for len(queue)!=0{
		n := len(queue)
		for i:=0;i<n;i++{
			p := queue[0]
			queue = queue[1:]
			if i==0{
				res = append(res,p.Val)
			}
			if p.Right != nil{
				queue = append(queue,p.Right)
			}
			if p.Left != nil{
				queue = append(queue,p.Left)
			}
		}
	}
	return res
}
// @lc code=end

