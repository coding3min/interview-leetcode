//题目链接：https://leetcode.cn/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof/
// day19/31
// 第 19 天主题为：搜索与回溯算法（中等）
// 包含三道题目：
// 剑指offer64.求1+2+...+n
// 剑指offer68-I.二叉搜索树的最近公共祖先
// 剑指offer68-II.二叉树的最近公共祖先
package main


//Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

//BST 的中序遍历是有序数组，root 节点的左子树中所有节点 val 都比 root.val 小，右子树中所有节点的 val 比root.val 大。
//根据这个特点，我们对 root、p、q 的节点值进行分情况讨论：
//
//1. p.val > root.val 且 q.val < root.val，说明 p 和 q 分别存在于 root 节点的右左子树中，root 节点为最近公共祖先（p.val < q.val 同理）；
//
//2. p.Val < root.Val 且 q.Val > root.Val，与第一种情况同理，p、q 分别在root 的左右子树中，root 节点为最近公共祖先；
//
//3. root.Val == p.Val 或者 root.Val == q.Val，说明另一个节点在 root 的左子树或者右子树中， root 已经是两节点的最近公共祖先，
//   因为我们是从整棵树的根节点开始往下遍历的
//
//4. p.Val < root.Val 且 q.Val < root.Val，p 和 q 均在 root 的左子树中，root 的左孩子节点会是 p 和 q 的公共祖先，
//   对 root.Left 进行递归操作寻找 p 和 q 的最近公共祖先
//
//5. p.Val > root.Val 且 q.Val > root.Val，与情况 4 同理，对 root.Right 进行递归操作寻找 p 和 q 的最近公共祖先。
//
//因为 1 和 2同理，4 和 5 同理，精简一些，只有三种情况，令 p.Val 小于 q.Val
//
//1. p.Val <= root.Val 或者 root.val < q.Val，此时 root 为 最近公共祖先
//2. p.Val < root.Val < q.Val，p 和 q 分别在 root 的左右子树中，root 为 最近公共祖先
//3. p.Val < root.Val 且 q.Val < root.Val，对 root.Left 递归操作；若均大于，对 root.Right 递归操作
//
//虽然我分析过程中写的是递归，但实际还是有递归和迭代两种写法的

//递归

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val<p.Val && root.Val<q.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}
	if root.Val>p.Val && root.Val>q.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}
	return root
}

//迭代
func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val<root.Val && q.Val<root.Val{
			root = root.Left
			continue
		}
		if p.Val>root.Val && q.Val>root.Val{
			root = root.Right
			continue
		}
		return root
	}
	// 此种情况不会发生，因为题目说明p和q均存在于给定BST中
	// 只是为了保证代码的逻辑正确
	return root
}