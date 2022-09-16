// 题目链接：https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
// day18/31
// 第 18 天主题为：搜索与回溯算法（中等）
// 包含两道题目：
// 剑指offer55-I.二叉树的深度
// 剑指offer55-II.平衡二叉树

package main

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}



func maxDepth_2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
	for len(q) != 0{
		res ++
		n := len(q)
		for i:=0;i<n;i++{
			cur := q[i]
			if cur.Left != nil{
				q = append(q,cur.Left)
			}
			if cur.Right != nil{
				q = append(q,cur.Right)
			}
		}
		q = q[n:]
	}
	return res
}