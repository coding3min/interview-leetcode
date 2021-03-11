/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
 *
 * https://leetcode-cn.com/problems/range-sum-of-bst/description/
 *
 * algorithms
 * Easy (77.98%)
 * Likes:    157
 * Dislikes: 0
 * Total Accepted:    42.4K
 * Total Submissions: 54.4K
 * Testcase Example:  '[10,5,15,3,7,null,18]\n7\n15'
 *
 * 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
 * 输出：32
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
 * 输出：23
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数目在范围 [1, 2 * 10^4] 内
 * 1 
 * 1 
 * 所有 Node.val 互不相同
 * 
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil{
		return 0
	}
	sum := 0
	if root.Val > low && root.Left !=nil{
		sum+=rangeSumBST(root.Left,low,high)
	}

	if root.Val < high && root.Right !=nil{
		sum+=rangeSumBST(root.Right,low,high)
	}
	if root.Val >= low && root.Val <= high{
		sum += root.Val
	}
	return sum 

}
// @lc code=end

