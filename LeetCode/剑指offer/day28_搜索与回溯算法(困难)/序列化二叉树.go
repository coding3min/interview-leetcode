//题目链接：https://leetcode.cn/problems/xu-lie-hua-er-cha-shu-lcof/?envType=study-plan&id=lcof
package main

import (
	"strconv"
	"strings"
)

//本题接参考自官方题解：https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/solution/er-cha-shu-de-xu-lie-hua-yu-fan-xu-lie-hua-by-le-2/
//
//二叉树的序列化本质上是对其值进行编码，更重要的是对其结构进行编码。可以遍历树来完成上述任务。
//众所周知，我们一般有两个策略：广度优先搜索和深度优先搜索。
//广度优先搜索可以按照层次的顺序从上到下遍历所有的节点
//
//深度优先搜索可以从一个根开始，一直延伸到某个叶，然后回到根，到达另一个分支。
//根据根节点、左节点和右节点之间的相对顺序，可以进一步将深度优先搜索策略区分为：先序遍历、中序遍历 以及 后序遍历。
//
//对题目所给例子进行举例，最终序列化字符串是 1,2,3,None,None,4,None,None,5,None,None,。
//其中，None，None 用来标记缺少左右子节点，这是我们在序列化期间保存树结构的方式。
//
//即我们可以先序遍历这颗二叉树，遇到空子树的时候序列化成 None，否则继续递归序列化。那么我们如何反序列化呢？
//首先我们需要根据 , 把原先的序列分割开来得到先序遍历的元素列表，然后从左向右遍历这个序列：
//- 如果当前的元素为 None，则当前为空树
//- 否则先解析这棵树的左子树，再解析它的右子树

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil{
			sb.WriteString("nil,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString(",")
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data,",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if s[0] == "nil"{
			s = s[1:]
			return nil
		}
		val,_ := strconv.Atoi(s[0])
		s = s[1:]
		node := &TreeNode{val,build(),build()}
		return node
	}
	return build()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */