/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 *
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (56.18%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    24.8K
 * Total Submissions: 44.1K
 * Testcase Example:  '[4,2,6,1,3]'
 *
 * 给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入: root = [4,2,6,1,3,null,null]
 * 输出: 1
 * 解释:
 * 注意，root是树节点对象(TreeNode object)，而不是数组。
 * 
 * 给定的树 [4,2,6,1,3,null,null] 可表示为下图:
 * 
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \    
 * ⁠   1   3  
 * 
 * 最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 二叉树的大小范围在 2 到 100。
 * 二叉树总是有效的，每个节点的值都是整数，且不重复。
 * 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
 * 相同
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
var(
	prev,ans int
)
func minDiffInBST(root *TreeNode) int {
	ans,prev = math.MaxInt32,math.MinInt32
	dfs(root)
	return ans
}
func dfs(root *TreeNode){
	if root == nil{
		return
	}
	dfs(root.Left)
	tmp := root.Val -  prev
	if tmp<ans{
		ans = tmp
	}
	prev = root.Val
	dfs(root.Right)
	return
}
// @lc code=end

