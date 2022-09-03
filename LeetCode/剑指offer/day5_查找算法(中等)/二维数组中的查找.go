// 题目链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/?plan=lcof&plan_progress=klczit3
// day5/31
// 第 5 天主题为：查找算法（中等）
// 包含三道题目：
// 剑指offer04.二维数组中的查找
// 剑指offer11.旋转数组的最小数字
// 剑指offer50.第一个只出现一次的字符

// 通俗解法就是遍历二维数组的每个元素，查看是否存在值等于target的元素，若存在，返回 true，否则，false。
// 但这完全没用到题目所给的递增条件，直接pass，就不附代码。
// 当我们需要解决一个复杂的问题时，一个很有效的办法就是从一个具体的问题入手，通过分析简单具体的例子，试图寻找普遍的规律。
// 利用下题目所给条件，每一行从左到右递增，每一列从上到下递增。这道题，找一下规律。
// 如果我们先选取数组右上角的数字，则该数字左侧元素小于其值，下方元素大于其值，类似二叉搜索树。
// 如果 target > 右上角数字的数值，则剔除第一行，第二行最后一个元素作为最右上角的数字，
// 同样的做法，每次移动，消除一行或者一列元素，直到找到目标值 target，或者 查找范围为空（即 target 不存在于二维数组中）。
package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0]) == 0{
		return false
	}
	m,n := len(matrix),len(matrix[0])
	i,j := 0,n-1
	for matrix[i][j] != target{
		if matrix[i][j] > target{
			j --
		} else {
			i ++
		}
		if i>=m || j <0{
			return false
		}
	}
	return true
}

// 另外，本题从左下角分析也可以，道理是一样的，下述代码为从左下角元素开始的“二分”
// 为避免函数命名冲突，次函数名添加后缀 “_2”
func findNumberIn2DArray_2(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0]) == 0{
		return false
	}
	m,n := len(matrix),len(matrix[0])
	i,j := m-1,0
	for matrix[i][j] != target{
		if matrix[i][j] > target{
			i --
		} else {
			j ++
		}
		if i<0 || j >= n{
			return false
		}
	}
	return true
}

// 本题考点：考察应聘者对二维数组的理解及编程能力，还考察应聘者分析问题的能力，当发现问题比较复杂时，是否能通过具体的例子找出其中的规律，是能够解决这个问题的关键所在。
// 本题只要从一个具体的二维数组的右上角开始分析，就能找到规律所在，从而找到解决问题的突破口。