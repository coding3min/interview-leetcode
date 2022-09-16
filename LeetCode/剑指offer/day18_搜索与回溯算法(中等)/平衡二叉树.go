// 题目链接：https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
package main

//这道题中的平衡二叉树的定义是：二叉树的每个节点的左右子树的高度差的绝对值不超过 1，则二叉树是平衡二叉树。
//根据定义，一棵二叉树是平衡二叉树，当且仅当其所有子树也都是平衡二叉树，因此可以使用递归的方式判断二叉树是不是平衡二叉树，
//递归的顺序可以是自顶向下或者自底向上。（这句话来自官方题解）
//
//AVL 的定义是递归的，因此很容易想到递归的解法。
//求当前节点左右子树的高度，然后判断高度之差是否不超过1，然后递归判断其左右子树是否为 AVL

// maxDepth 和 max 函数 在 二叉树的深度.go 文件中，这里直接调用了
func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	leftHeight := maxDepth(root.Left)
	rightHright := maxDepth(root.Right)
	return abs(leftHeight-rightHright) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0{
		return -x
	}
	return x
}

//刚才这种递归是自顶向下的，对于同一个节点，函数 maxDepth 会被重复调用，调用次数为其祖先节点的个数。
//如果使用自底向上的递归，则对每个节点，函数 maxDepth 只会被调用一次。
//
//本解法参考自官方题解：https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/solution/ping-heng-er-cha-shu-by-leetcode-solutio-6r1g/
//自底向上的做法类似后序遍历，对当前遍历到的节点，先递归地判断其左右子树是否平衡，再判断以当前节点为根节点的子树是否平衡，
//这样，在其子树不平衡时，我们直接返回 false即可。如果一棵树是平衡的，返回其高度，否则返回 -1。
//只要存在子树不平衡，则整棵树必定不平衡。

func isBalanced_2(root *TreeNode) bool {
	return depth(root) >= 0
}

func depth(node *TreeNode) int {
	if node == nil{
		return 0
	}
	leftDepth := depth(node.Left)
	rightDepth := depth(node.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if abs(leftDepth-rightDepth)>1{
		return -1
	}
	return 1+max(leftDepth,rightDepth)
}