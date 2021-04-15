/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (80.14%)
 * Likes:    408
 * Dislikes: 0
 * Total Accepted:    97.6K
 * Total Submissions: 122K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：[[1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	var res [][]int
	if n == 0{
		return res
	}
	res = make([][]int,n)
	for i:=0;i<n;i++{
		res[i] = make([]int,n)
	}
	t,l,r,b := 0,0,n-1,n-1
	num,end := 1,n*n
	for num<=end{
		for i:=l;i<=r;i++{
			res[t][i] = num
			num++
		}
		t++
		for i:=t;i<=b;i++{
			res[i][r] = num
			num++
		}
		r--
		for i:=r;i>=l;i--{
			res[b][i] = num
			num++
		}
		b--
		for i:=b;i>=t;i--{
			res[i][l] = num
			num++
		}
		l++
	}
	return res
}
// @lc code=end

