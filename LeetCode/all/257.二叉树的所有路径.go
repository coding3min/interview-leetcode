/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (66.50%)
 * Likes:    440
 * Dislikes: 0
 * Total Accepted:    95.8K
 * Total Submissions: 143.9K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
var(
	ans []string
)
func binaryTreePaths(root *TreeNode) []string {
	ans = ans[:0]
	dfs(root,"")
	return ans
}
func dfs(root *TreeNode,path string){
	if root == nil{
		return
	}
	newStr := path + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil{
		ans = append(ans,newStr)
		return
	}
	newStr = newStr + "->"
	if root.Left !=nil{
		dfs(root.Left,newStr)
	}
	if root.Right != nil{
		dfs(root.Right,newStr)
	}
	return
}
// @lc code=end

