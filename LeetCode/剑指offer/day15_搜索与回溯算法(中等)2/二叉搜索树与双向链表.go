//题目链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/?envType=study-plan&id=lcof
package main

//BST 的最重要的性质为：中序遍历为递增序列
//我们一定是需要对 BST 进行中序遍历的，因为要改变指针的指向，所以遍历的同时，需要记录当前节点的前驱结点，声明变量 pre 和 head 为 *TreeNode 类型，pre 记录当前节点的前驱结点，head 记录最终返回的头结点。
//中序遍历过程中，处理当前节点 cur 时，若 pre 为空，说明当前节点为中序遍历初始值，pre 为空，那自然也不需要调整 pre 和 cur 的指向，
//若 pre 非空，调整 cur.Right 和 cur.Left 的指向，pre.Right = cur，cur.Left = pre；
//处理完成后，pre = cur，当前节点为中序遍历中下一个节点的前驱结点。
//递归结束后，pre 指向 BST 中节点值最大的节点，即中序遍历最后一个节点。
//然后，调整 pre.Right 和 head.Left 的指向，形成循环链表。
//
//很棒的一道题目。
//这道题 LeetCode 上不支持 Go 语言，我去牛客网验证了一下，牛客网没有要求循环链表，代码倒数第三行和倒数第四行需要注释掉才能通过，LeetCode 要求循环，将其取消注释，理论上可以通过。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 *
 * @param pRootOfTree TreeNode类
 * @return TreeNode类
 */
func Convert( pRootOfTree *TreeNode ) *TreeNode {
	// write code here
	var pre,head *TreeNode
	var inorder func(cur *TreeNode)
	inorder = func(cur *TreeNode) {
		if cur == nil{
			return
		}
		inorder(cur.Left)
		if pre == nil{
			head = cur
		} else {
			pre.Right = cur
			cur.Left = pre
		}
		pre = cur
		inorder(cur.Right)
	}
	inorder(pRootOfTree)
	// head.Left = pre
	// pre.Right = head
	return head
}
