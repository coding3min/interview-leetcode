/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 *
 * https://leetcode-cn.com/problems/rotate-image/description/
 *
 * algorithms
 * Medium (73.63%)
 * Likes:    1106
 * Dislikes: 0
 * Total Accepted:    252.7K
 * Total Submissions: 343.2K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 * 
 * 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[7,4,1],[8,5,2],[9,6,3]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：matrix = [[1]]
 * 输出：[[1]]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：matrix = [[1,2],[3,4]]
 * 输出：[[3,1],[4,2]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * matrix.length == n
 * matrix[i].length == n
 * 1 
 * -1000 
 * 
 * 
 */

// @lc code=start
func rotate(matrix [][]int)  {
	n := len(matrix)
	// 对角线对折
	for i:=0; i<n;i++{
		for j:=0;j<n-i;j++{
			matrix[i][j],matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i],matrix[i][j]
		}
	}
	// 上下翻转
	for i:=0;i<n/2;i++{
		for j:=0;j<n;j++{
			matrix[i][j],matrix[n-1-i][j] = matrix[n-1-i][j],matrix[i][j]
		}
	}
}
// @lc code=end

