// 题目链接：https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/?envType=study-plan&id=lcof
// day7/31
// 第 7 天主题为：搜索与回溯算法（简单），与前一天主题相同
// 包含三道题目：
// 剑指offer26.树的子结构
// 剑指offer27.二叉树的镜像
// 剑指offer28.对称的二叉树


package main


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 解题思路：观察例子可以发现，镜像 就是 每个节点的左右子树进行翻转。所以，两步完成本题：
//1. 遍历二叉树保存所有节点（前中后序遍历都可以），用节点的指针类型切片保存；
//2. 遍历切片，交换其左右子节点 node.Left,node.Right = node.Right,node.Left
func mirrorTree(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{}
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode){
		if root == nil{
			return
		}
		nodes = append(nodes,root)
		if root.Left != nil{
			preorder(root.Left)
		}
		if root.Right != nil{
			preorder(root.Right)
		}
	}
	preorder(root)
	for _,node := range nodes{
		node.Left,node.Right = node.Right,node.Left
	}
	return root
}

// 上面这种解法的时间空间复杂度均为 O(n)
// 我们也可以在遍历的过程中交换节点的左右子节点，省去 nodes 切片占用的内存空间，但这本质上并不会降低时间和空间复杂度，
// 因为我们在对二叉树进行遍历的过程中，调用了系统栈，系统需要使用 O(n) 大小的栈空间。
// 前序遍历的过程中交换左右子节点
// 为避免函数命名冲突，次函数名添加后缀 “_2”
func mirrorTree_2(root *TreeNode) *TreeNode {
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode){
		if root == nil{
			return
		}
		root.Left,root.Right = root.Right,root.Left
		if root.Left != nil{
			preorder(root.Left)
		}
		if root.Right != nil{
			preorder(root.Right)
		}
	}
	preorder(root)
	return root
}


// 另外，在LeetCode题解区我还看到了有人用层序遍历的方式解题，本质上和上面两种思路是一样的，就是选择了另一种遍历二叉树的方式，
// 交换左右子节点的核心操作无任何变化，代码开头的判断 root 节点是否为空不可省略，因为 q 初始化默认加入root节点，若该节点为空，后续代码会出错。
// 为避免函数命名冲突，次函数名添加后缀 “_3”
func mirrorTree_3(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	q := []*TreeNode{root}
	for len(q) != 0{
		node := q[0]
		node.Left,node.Right = node.Right,node.Left
		if node.Left != nil{
			q = append(q,node.Left)
		}
		if node.Right != nil{
			q = append(q,node.Right)
		}
		q = q[1:]
	}
	return root
}


// 二叉树的定义就是递归的，所以做二叉树的题目一定要有递归的想法，有思路后开始着手写代码，不要在脑海中模拟太多层的递归，否则很容易绕进去。