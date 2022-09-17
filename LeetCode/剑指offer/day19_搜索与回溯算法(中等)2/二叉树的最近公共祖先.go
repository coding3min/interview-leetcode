//题目链接：https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/?envType=study-plan&id=lcof
package main

//取消了第一版题目BST的条件，退化为普通的二叉树，我们就只能相对暴力地去解决这个问题。
//
//解法参考自题解：https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/solution/mian-shi-ti-68-ii-er-cha-shu-de-zui-jin-gong-gon-7/
//
//考虑通过递归对二叉树进行先序遍历，遇到 p 或 q 时返回，从底至顶回溯，当节点 p、q 分别在 root 的两侧 或 p、q 两者之一就是 root 节点时，
//root 为 p 和 q 的LCA，返回 root。
//
//然后具体分析如何实现：
//
//1. 终止条件：
//	1. 当越过叶节点时，直接返回 nil；
//	2. root 等于 p 或 q 时，返回 root
//2. 递推工作：
//	1. 递归 root.Left，记返回值为 left；
//	2. 递归 root.Right，记返回值为 right；
//3. 返回值：根据 left 和 right，分为以下四种情况：
//	1. left 和 right 同时为空，说明 root 的左右子树均不包含 p 和 q，以 root 为根节点的子树不可能包含 p 和 q，返回 nil
//	2. left 为 空，right 不为空，说明 p 和 q 都包含在 root.Right 中，返回 right（分两种情况，
//		第一种：p 和 q 其中一个节点在 root 的右子树中，此时 right 为所包含的节点；
//		第二种：p 和 q 均在 root 的右子树中，此时 right 指向 p 和 q 的 LCA）
//	3. right 为 空，left 不为空，与情况2同理
//	4. left 和 right 均不为空，说明 p、q 分别在 root 的左右子树中，root 为 LCA，返回 root 即可。
func lowestCommonAncestor_3(root, p, q *TreeNode) *TreeNode {
	// 递归结束条件
	if root == nil{
		return nil
	}
	if root.Val==p.Val || root.Val==q.Val{
		return root
	}
	// 开始递归
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	// 情况1
	if left==nil && right==nil{
		return nil
	}
	// 情况2
	if left == nil{
		return right
	}
	// 情况3
	if right == nil{
		return left
	}
	// 情况4
	return root
}

//这里想看过的一段话：
// 学习一个算法，最要的是弄清楚这个算法要解决什么样的问题，它的已知量（input）是什么？
// 待求的未知量（output）是什么
// 如果这三个问题没找到答案就去学习算法，就会浪费大量时间在逻辑猜测与记忆上

//针对本题，对 lowestCommonAncestor 函数，本人分析如下
//函数 lowestCommonAncestor 输入为 root、p 和 q
//
//输出分几种情况：
//1. 若 root 为空节点，直接返回 root 即可
//2. 以 root 为根节点的树包含 p 和 q，返回 p 和 q 的 LCA；
//3. 以 root 为根节点的树只包含 p 或 q 其中一个，则返回 包含的节点；
//4. 以 root 为根节点的树不包含 p 且 不包含 q，返回 nil

//解决问题：在以 root 为根节点的子树中，寻找 p 和 q 的 LCA
