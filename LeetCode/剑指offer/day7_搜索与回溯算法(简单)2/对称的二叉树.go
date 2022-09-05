// 题目链接：https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/?envType=study-plan&id=lcof

package main


// 解题思路：如果二叉树只有一个根节点，那其必然是对称的，本题主要是判断节点的左右子树是否对称。
//
// 判断左右子树是否对称，若两子树根节点均为空，说明对称，若某子树为空而另一颗子树不为空，说明不对称，两子树根节点值不相等的话也必定不对称吗，否则，对称，之后递归地判断
//- 左子树的左子树 与 右子树的右子树 是否对称
//- 左子树的右子树 与 右子树的左子树 是否对称
// 两者均对称时，说明该左右子树对称；否则，不对称。

func isSymmetric(root *TreeNode) bool {
	// 根节点为空，对称
	if root == nil{
		return true
	}
	var sym func(x,y *TreeNode) bool
	// 递归判断左右子树是否对称
	sym = func(x,y *TreeNode) bool {
		// 两者均为空说明对称
		if x == nil && y == nil{
			return true
		//	其中某一子树为空，而另一子树不为空，不对称
		} else if x == nil || y == nil{
			return false
		}
		// 两者根节点值不相等时也不对称
		if x.Val != y.Val{
			return false
		}
		// 递归判断x的左右子树 与 y 的右左子树是否对称
		return sym(x.Left,y.Right) && sym(x.Right,y.Left)
	}
	return sym(root.Left,root.Right)
}