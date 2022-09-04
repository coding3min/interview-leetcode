//题目链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/?envType=study-plan&id=lcof

// 在第二版的基础上，加入了之字形的条件，解题思路：加入变量order，初始化为0，表示遍历方向，order%2为偶数时，从左至右，
// 否则，从右至左。这只代表了当前层节点遍历的顺序，
// 注意！添加下一层节点时始终是从左至右的顺序。第一次提交错误，就是因为添加下一层节点时，顺序搞错了。
package main

// 为避免函数命名冲突，次函数名添加后缀 “_3”
func levelOrder_3(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{root}
	order := 0
	for len(q) != 0{
		n := len(q)
		temp := []int{}
		for i:=0;i<n;i++{
			if q[i].Left != nil{
				q = append(q,q[i].Left)
			}
			if q[i].Right != nil{
				q = append(q,q[i].Right)
			}
		}
		if order % 2 == 0{
			for i:=0;i<n;i++{
				temp = append(temp,q[i].Val)
			}
		} else {
			for i:=n-1;i>=0;i--{
				temp = append(temp,q[i].Val)
			}
		}
		order ++
		q = q[n:]
		res = append(res,temp)
	}
	return res
}


// 然后看了下官方的操作，挺不错的，先按照第二版从上到下打印二叉树的思路解题，
// 最后向返回数组添加某一层遍历结果时，先将该层遍历结果进行翻转，简单易懂。
// 为避免函数命名冲突，次函数名添加后缀 “_4”
func levelOrder_4(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}
