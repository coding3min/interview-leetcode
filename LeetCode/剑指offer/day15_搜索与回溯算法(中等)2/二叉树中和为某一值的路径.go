package main





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
