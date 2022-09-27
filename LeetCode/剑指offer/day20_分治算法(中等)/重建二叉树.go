//题目链接：https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/?envType=study-plan&id=lcof
package main

//解题思路：参考自题解(https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/solution/mian-shi-ti-07-zhong-jian-er-cha-shu-by-leetcode-s/
//
//对于任意一颗树而言，前序遍历的形式总是：[根节点,[左子树的前序遍历结果], [右子树的前序遍历结果]]
//即根节点总是前序遍历中的第一个节点。
//而中序遍历的形式总是：[[左子树的中序遍历结果],根节点,[右子树的中序遍历结果]]
//
//根据二叉树的前序遍历，我们可以构造出根节点，之后要考虑的是如何构造根节点的左孩子节点和右孩子节点
//只要我们在中序遍历中定位到根节点，我们就可以知道该二叉树左子树和右子树的节点数量。
//知道数量后，我们就可以在前序遍历和中序遍历中通过切片得到左子树和右子树的前序遍历和中序遍历。
//
//之后递归构造根节点的左孩子节点和右孩子节点即可
//
//在中序遍历中对根节点进行定位时，我们这里使用遍历的方式

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 递归退出条件，遍历序列长度为0，节点为空
	if len(preorder) == 0{
		return nil
	}
	rootVal := preorder[0]
	// 构造根节点
	root := &TreeNode{rootVal,nil,nil}
	// k初始化为0，寻找中序遍历中根节点的下标
	k := 0
	for ;k<len(preorder);k++{
		if inorder[k] == rootVal{
			break
		}
	}
	// 递归构造根节点的左孩子节点
	root.Left = buildTree(preorder[1:k+1],inorder[:k])
	// 递归构造根节点的右孩子节点
	root.Right = buildTree(preorder[k+1:],inorder[k+1:])
	return root
}