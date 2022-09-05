// 题目链接：https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/

package main


// 此题解参考自：https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/solution/mian-shi-ti-26-shu-de-zi-jie-gou-xian-xu-bian-li-p/
// 若树 B 是 A 的子结构，则 A 子结构的根节点可能是 A 中任意一个节点，我们判断 B 是否为 A 的子结构，就需要遍历 A 的所有节点，
// 然后判断是否有某节点 以 该节点为 根节点的子树包含（有点绕，其实就是，若 A 与 B 根节点相同， B 是否为 A 的子结构，
// 但当 B 为空的时候，B 必为 A 的子结构）。
// 名词规定：树 A 的根节点记作 节点 A，树 B 的根节点称为 节点 B 。
// recur(A, B)函数，用于判断树 A 中以 A为根节点的子结构是否包含树 B
//终止条件：
//- 当节点 B 为空：说明树 B 已匹配完成（越过叶子节点），因此返回 true ；
//- 当节点 A 为空：说明已经越过树 A 叶子节点（而 B 节点非空），即匹配失败，返回 false ；
//- 当节点 A 和 B 的值不同：说明匹配失败，返回 false ；
//返回值：
//- 判断 A 和 B 的左子节点是否相等，即 recur(A.left, B.left) ；
//- 判断 A 和 B 的右子节点是否相等，即 recur(A.right, B.right) ；
//两者取逻辑与后返回。
// 特例处理：当 树 A 为空 或 树 B 为空时，直接返回 false（对应题目中的约定：空树不是任意一个树的子结构） 。
// 之后我们对 A 作遍历（前中后序都可以），对其中每个节点与 B 进行 recur 判断，若存在 true 结果，返回最终的 true；若全为 false，说明 B 不是 A 的子结构，返回 false。
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil{
		return false
	}
	nodes := []*TreeNode{}
	// 对树的遍历，这里我采用前序遍历
	// 使用 中序遍历、后序遍历 or 层序遍历都是可以的
	var perorder func(root *TreeNode)
	perorder = func(root *TreeNode){
		if root == nil{
			return
		}
		nodes = append(nodes,root)
		if root.Left != nil{
			perorder(root.Left)
		}
		if root.Right != nil{
			perorder(root.Right)
		}
	}
	perorder(A)
	// 判断B是否为以为根节点的子树的在子结构
	var sub func(A,B *TreeNode) bool
	sub = func(A,B *TreeNode) bool{
		if B == nil{
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return sub(A.Left,B.Left) && sub(A.Right,B.Right)
	}
	for _,node := range nodes{
		if sub(node,B){
			return true
		}
	}
	return false
}