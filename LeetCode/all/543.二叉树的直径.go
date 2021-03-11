/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 *
 * https://leetcode-cn.com/problems/diameter-of-binary-tree/description/
 *
 * algorithms
 * Easy (52.20%)
 * Likes:    622
 * Dislikes: 0
 * Total Accepted:    94.8K
 * Total Submissions: 181.5K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 * 
 * 
 * 
 * 示例 :
 * 给定二叉树
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       2   3
 * ⁠      / \     
 * ⁠     4   5    
 * 
 * 
 * 返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
 * 
 * 
 * 
 * 注意：两结点之间的路径长度是以它们之间边的数目表示。
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
func diameterOfBinaryTree(root *TreeNode) int {
	var dps func(root *TreeNode) int
	var maxRes int
	dps = func(root *TreeNode)int{
		if root == nil{
				return 0
			}
		x := dps(root.Left)
		y := dps(root.Right)
		maxRes = max(x+y,maxRes)
		return max(x,y)+1
	}
	dps(root)
	return maxRes
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}
// @lc code=end

