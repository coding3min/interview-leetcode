// 题目链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/?envType=study-plan&id=lcof
// day15/31
// 第 145 天主题为：搜索与回溯算法（中等）2
// 包含三道题目：
// 剑指offer34.二叉树中和为某一值的路径
// 剑指offer36.二叉搜索树与双向链表
// 剑指offer54.二叉搜索树的第k大节点

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//通俗解法：提到 BST，我们都知道其中序遍历是一个递增有序序列，先通过中序遍历得到该序列，然后返回倒数第 k 个值即可。
func kthLargest(root *TreeNode, k int) int {
	nums := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode){
		if node == nil{
			return
		}
		inorder(node.Left)
		nums = append(nums,node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return nums[len(nums)-k]
}


//我们要找第k大的元素，而我们中序序列得到的是一个递增有序序列，导致我们需要保存下来中序遍历的元素才能找到第 k 大的元素，如果我们能够得到递减的有序序列，
//那我们就不再需要保存所有节点值，而是遍历到底 k 个有效节点时，保存该节点值，最后返回即可。
//
//如何做到呢？中序遍历顺序是 [左子树的中序遍历，根节点，右子树的中序遍历]，其实只要把左右子树的遍历顺序颠倒一下，就可以得到从大到小的序列。
//我们用一个参数保存当前是遍历的第几个节点，不再需要一个数组保存中序遍历的序列，可以将空间复杂度从 O(n) 降低到 O(1)。
func kthLargest_2(root *TreeNode, k int) int {
	i := 0
	res := 0
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil{
			return
		}
		inorder(node.Right)
		i ++
		if i == k{
			res = node.Val
		}
		inorder(node.Left)
	}
	inorder(root)
	return res
}