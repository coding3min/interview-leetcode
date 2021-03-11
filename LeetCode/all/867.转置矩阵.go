/*
 * @lc app=leetcode.cn id=867 lang=golang
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0{
		return matrix
	}
	m := len(matrix[0])
	res := make([][]int,m)
	for i:=0;i<m;i++{
		res[i] = make([]int,n)
		for j:=0;j<n;j++{
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
// @lc code=end

