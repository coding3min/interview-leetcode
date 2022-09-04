// 题目链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// day6/31
// 第 6 天主题为：搜素与回溯算法（简单）
// 包含三道题目：
// 剑指offer32-I.从上到下打印二叉树
// 剑指offer32-II.从上到下打印二叉树II
// 剑指offer32-III.从上到下打印二叉树III

// 本题题解参考自：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/solution/mian-shi-ti-32-i-cong-shang-dao-xia-da-yin-er-ch-4/
// 解题思路：从上到下打印二叉树，也就是二叉树的层序遍历，考察 BFS，BFS通常借助队列的先入先出特性来实现。
//算法流程：
//1. 特例处理： 当树的根节点为空，则直接返回空切片 []int{} ；
//2. 初始化： 打印结果列表 res:=make([]int,0) ，包含根节点的队列 q:=[]*TreeNode{root}；
//3. BFS 循环： 当队列 queue 为空时跳出；
//	1. 出队： 队首元素出队，记为 node；
//	2. 打印： 将 node.val 添加至列表 res 尾部；
//	3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
//4. 返回值： 返回打印结果列表 res 即可。
package main



// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q)!=0{
		cur := q[0]
		if cur.Left != nil{
			q=append(q,cur.Left)
		}
		if cur.Right != nil{
			q = append(q,cur.Right)
		}
		res = append(res,cur.Val)
		q = q[1:]
	}
	return res
}