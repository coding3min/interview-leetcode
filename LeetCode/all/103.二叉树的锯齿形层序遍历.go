/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (57.03%)
 * Likes:    387
 * Dislikes: 0
 * Total Accepted:    114K
 * Total Submissions: 199.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层序遍历如下：
 * 
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil{
		return ans
	}
	queue := []*TreeNode{root}
	isDesc := false

	for len(queue)!=0{
		n := len(queue)
		var levelQueue []int
		for i:=0;i<n;i++{
			node := queue[i]
			if isDesc{
				tmp := []int{node.Val}
				levelQueue = append(tmp,levelQueue...)
			}else{
				levelQueue = append(levelQueue,node.Val)
			}
			if node.Left != nil {queue = append(queue,node.Left)}
			if node.Right != nil {queue = append(queue,node.Right)}
		}
		// 一次性出队或者一个一个出队都可以
		queue = queue[n:]
		ans = append(ans, levelQueue)
		isDesc = !isDesc
	}
	return ans
}
// @lc code=end

