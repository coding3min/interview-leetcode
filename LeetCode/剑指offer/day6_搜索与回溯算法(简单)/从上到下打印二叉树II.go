// 题目链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/?envType=study-plan&id=lcof


package main

// 解题思路：本质还是考察层序遍历，我们依旧使用队列来实现。
// 将每层遍历到的节点放在列表中。首先，获取当前q的长度，即当前层节点的数目 n，然后依次访问队列 q 的前 n 个节点，
// 两步操作：
//	1. 将节点 val 添加到当前层的切片中；
//	2. 若当前节点存在孩子节点，按顺序将其添加至列表 q
// 最后，q = q[n:]，删除列表中已经访问过的当前层节点，若 q 非空，下一次循环进入下一层节点的处理。

//算法流程：（算法流程来自题解：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/solution/mian-shi-ti-32-ii-cong-shang-dao-xia-da-yin-er-c-5/）
// 写的很系统，值得学习
//1. 特例处理： 当根节点为空，则返回空切片 []int{} ；
//2. 初始化： 打印结果列表 res:=[]int{} ，包含根节点的队列 queue:=[]*TreeNode{root} ；
//3. BFS 循环： 当队列 queue 为空时跳出；
//	1. 新建一个临时列表 tmp ，用于存储当前层打印结果；
//	2. 当前层打印循环： 循环次数为当前层节点数（即队列 queue 长度）；
//		1. 出队： 队首元素出队，记为 node；
//		2. 打印： 将 node.val 添加至 tmp 尾部；
//		3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
//	3. 将当前层结果 tmp 添加入 res 。
//4. 返回值： 返回打印结果列表 res 即可。

// 为避免函数命名冲突，次函数名添加后缀 “_2”
func levelOrder_2(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{root}
	for len(q) != 0{
		n := len(q)
		temp := []int{}
		for i:=0;i<n;i++{
			cur := q[i]
			temp = append(temp,cur.Val)
			if cur.Left != nil{
				q = append(q,cur.Left)
			}
			if cur.Right != nil{
				q = append(q,cur.Right)
			}
		}
		res = append(res,temp)
		q = q[n:]
	}
	return res
}