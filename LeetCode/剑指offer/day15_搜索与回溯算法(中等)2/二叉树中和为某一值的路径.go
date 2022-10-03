//题目链接：https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
package main


//方法一：DFS，枚举每一条从根节点到叶子节点的路径，遍历到叶子节点时，如果路径和恰好为 target，说明找到一个满足条件的路径。
//注意代码中 copy(x,path) 操作，将 path 进行拷贝后再加入 res，若直接 res = append(res,path)，之后的路径将当前路径覆盖后，res 中的切片也会发生变化，这里是个小坑，需要注意一下。


//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//
func pathSum(root *TreeNode, target int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrace func(node *TreeNode,path []int,sum int)
	backtrace = func(node *TreeNode,path []int,sum int){
		if node == nil{
			return
		}
		path = append(path,node.Val)
		sum += node.Val
		if node.Left==nil && node.Right==nil{
			if sum == target{
				x := make([]int,len(path))
				copy(x,path)
				res = append(res,x)
				return
			} else {
				return
			}
		}
		backtrace(node.Left,path,sum)
		backtrace(node.Right,path,sum)
	}
	backtrace(root,path,0)
	return res
}

// 本题还有BFS解法，我改了很久总是出问题，ac后再附上

// 回来了，已 ac
//方法二：BFS
//pair 结构体保存节点以及从根节点到该节点的路径值。
//BFS 必定要用到队列，队列初始化时，加入根节点的 pair，sum 属性为 root.Val，同时，为了获取到叶子节点到根节点的路径，
//我们声明一个 parent 字典，保存节点的父节点，在进行 BFS 的同时更新 parent。
type pair struct {
	node  *TreeNode
	sum  int
}

func pathSum_2(root *TreeNode, target int) [][]int {
	res := [][]int{}
	// 判断 root 是否为空，对 root 为空的情况提前处理
	if root == nil{
		return res
	}
	// 队列初始化只保存根节点，利用队列实现 BFS
	q := []pair{{root,root.Val}}
	// parent字典帮助寻找叶子节点到根节点的路径
	parent := map[*TreeNode]*TreeNode{}
	// getPath利用parent寻找叶子节点到根节点的路径
	getPath := func(node *TreeNode) []int {
		path := []int{}
		// 不断向根节点回退
		for ;node!=nil;node=parent[node]{
			path = append(path,node.Val)
		}
		// path倒序，变为从根节点到叶子节点的路径
		left,right := 0,len(path)-1
		for left < right{
			path[left],path[right] = path[right],path[left]
			left ++
			right --
		}
		return path
	}
	// BFS
	for len(q) != 0{
		// 取队首元素
		cur := q[0]
		node := cur.node
		// 先更新 parent，将孩子节点指向该节点
		if node.Left != nil{
			parent[node.Left] = node
		}
		if node.Right != nil{
			parent[node.Right] = node
		}
		// 走到叶节点
		if node.Left==nil && node.Right==nil{
			// 且路径和等于target，更新res
			if cur.sum == target{
				res = append(res,getPath(node))
			}
		} else {
			// 若未到根节点
			// 将孩子节点加入队列
			if node.Left != nil{
				q = append(q,pair{node.Left,cur.sum+node.Left.Val})
			}
			if node.Right != nil{
				q = append(q,pair{node.Right,cur.sum+node.Right.Val})
			}
		}
		// 队首元素出队
		q = q[1:]
	}
	return res
}