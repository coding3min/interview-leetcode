/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (41.01%)
 * Likes:    426
 * Dislikes: 0
 * Total Accepted:    132.3K
 * Total Submissions: 297.3K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 * 
 * 
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -10^4 
 * 
 * 
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	n1 := len(matrix)
	if n1==0{
		return false
	}
	n2 := len(matrix[0])
	n := n1*n2
	l,r := 0,n-1
	for l<=r{
		mid := (r-l+1)/2 + l
		fmt.Println(mid,mid/n2,mid%n2)
		num := matrix[mid/n2][mid%n2]
		if num<target{
			l = mid+1
		}else if num >target{
			r = mid-1
		}else{
			return true
		}
	}
	return false
}

//简单搜索
// func searchMatrix(matrix [][]int, target int) bool {
// 	n := len(matrix)
// 	if n == 0{
// 		return false
// 	}
// 	var raw,col int
// 	raw = n-1
// 	for raw>=0 && col < len(matrix[raw]){
// 		num := matrix[raw][col]
// 		if num == target{
// 			return true
// 		}
// 		if num > target{
// 			raw--
// 		}
// 		if num < target{
// 			col++
// 		}
// 	}
// 	return false
// }
// @lc code=end

