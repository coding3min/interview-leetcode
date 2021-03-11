/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (61.51%)
 * Likes:    426
 * Dislikes: 0
 * Total Accepted:    116.4K
 * Total Submissions: 188.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 * 
 * 
 * 返回:
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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
var(
    res [][]int
)
func pathSum(root *TreeNode, targetSum int) [][]int {
    res = [][]int{}
    dps(root,targetSum,0,[]int{})
    return res
}

func dps(root *TreeNode,targetSum,currSum int,router []int){
    if root == nil{
        return
    }
    currSum = currSum + root.Val
    router = append(router,root.Val)
    if currSum == targetSum && root.Left ==nil && root.Right ==nil {
        tmp := make([]int,len(router))
        copy(tmp,router)
        res = append(res,tmp)
        return
    }
    if root.Left !=nil{
        dps(root.Left,targetSum,currSum,router)
    }
    if root.Right!=nil{
        dps(root.Right,targetSum,currSum,router)
    }
}
// @lc code=end

