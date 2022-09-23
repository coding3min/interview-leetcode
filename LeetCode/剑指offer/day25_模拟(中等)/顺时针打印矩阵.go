//题目链接：https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/?envType=study-plan&id=lcof
package main

//简单模拟，这种题目，个人感觉比较看经验，做过类似题目的话，很快可以模拟出来，没做过或者没思路的话，就有点无从下手或者不知道如何用代码实现自己的想法。
//需要考虑三个问题：移动方向、边界 和 结束条件。
//
//1. 移动方向：很明显为 右、下、左、上 这样的循环，我们可以用一个二维数组代表四个方向，每移动到边界，更改方向；
//2. 边界：边界问题是本题的重点，因为边界是随着遍历在变化的，打印矩阵的过程中，边界逐渐变小。规则是：如果当前行（列）遍历结束后，将次行（列）的边界向内移动一格；
//3. 结束条件：当矩阵所有位置都被打印过，即遍历数组的长度等于矩阵元素个数时，结束遍历。
//
//代码中，x和y代表当前遍历到的元素所在位置索引，dirs代表移动的四个方向（上、右、下、左是一个循环），
//dir[cur_d]代表下一次移动的方向，结束条件是 res 数组元素个数等于矩阵元素个数。
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return []int{}
	}
	m,n := len(matrix),len(matrix[0])
	directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
	cur_dir := 0
	top,right,bottom,left := 0,n-1,m-1,0
	res := make([]int,0)
	x,y := 0,0
	for len(res) < m*n{
		res = append(res,matrix[x][y])
		if cur_dir == 0 && y == right{
			cur_dir ++
			top ++
		} else if cur_dir == 1 && x == bottom{
			cur_dir ++
			right --
		} else if cur_dir == 2 && y == left{
			cur_dir ++
			bottom --
		} else if cur_dir == 3 && x == top{
			cur_dir ++
			left ++
		}
		cur_dir %= 4
		x += directions[cur_dir][0]
		y += directions[cur_dir][1]
	}
	return res
}

//也可以按层模拟，将矩阵分为若干层，先打印最外层元素，再打印次外层的元素，直到输出完最内层的元素。
//
//对于每层，从左上方开始以顺时针的顺序遍历当前层所有元素。还是要着重考虑边界的问题，假设当前层左上角元素索引为(top,left)，右下角元素索引为(bottom,right).
//
//- 从左到右遍历上侧元素，依次为 (top,left) 到 (top,right)。
//- 从上到下遍历右侧元素，依次为 (top+1,right) 到 (bottom,right)。
//- 如果 left<right 且 top<bottom，则从右到左遍历下侧元素，依次为 (bottom,right−1) 到 (bottom,left+1)，以及从下到上遍历左侧元素，依次为 (bottom,left) 到 (top+1,left)。
//
//遍历完当前层的元素之后，将 left 和 top 分别增加 1，将 right 和 bottom 分别减少 1，进入下一层继续遍历，直到遍历完所有元素为止。

func spiralOrder_2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return []int{}
	}
	m,n := len(matrix),len(matrix[0])
	top,right,bottom,left := 0,n-1,m-1,0
	res := []int{}
	for left <= right && top <= bottom{
		for i := left;i<=right;i++{
			res = append(res,matrix[top][i])
		}
		for i := top+1;i<=bottom;i++{
			res = append(res,matrix[i][right])
		}
		// 这里一定要进行判断，否则会重复收录元素
		if left < right && top < bottom{
			for i := right-1;i>=left;i--{
				res = append(res,matrix[bottom][i])
			}
			for i := bottom-1;i>top;i--{
				res = append(res,matrix[i][left])
			}
		}
		top ++
		right --
		bottom --
		left ++
	}
	return res
}