/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (33.43%)
 * Likes:    920
 * Dislikes: 0
 * Total Accepted:    221K
 * Total Submissions: 659.6K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
	prev int
)
func isValidBST(root *TreeNode) bool {
	prev = math.MinInt64
	return isBST(root)
}
func isBST(root *TreeNode) bool{
	if root==nil{
		return true
	}
	if !isBST(root.Left){
		return false
	}
	if root.Val<=prev{
		return false
	}
	prev = root.Val
	if !isBST(root.Right){
		return false
	}
	return true
}




// 使用中序遍历，超时
// var(
// 	arr []int
// )
// func isValidBST(root *TreeNode) bool {
// 	isOk := rangeTree(root)
// 	arr = []int{}
// 	return isOk
// }

// func rangeTree(root *TreeNode) bool{
// 	if root == nil{
// 		return true
// 	}
// 	isLeft := rangeTree(root.Left)
// 	if len(arr)>0{
// 		fmt.Println("xxx")
// 		fmt.Println(arr)
// 		if !isLeft || root.Val<=arr[len(arr)-1]{
// 			fmt.Println("xxx")
// 			return false
// 		}
// 	}
// 	arr = append(arr,root.Val)
// 	isRight := rangeTree(root.Right)

// 	return isRight
// }
// @lc code=end

