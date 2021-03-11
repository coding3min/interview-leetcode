/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (52.21%)
 * Likes:    980
 * Dislikes: 0
 * Total Accepted:    206.6K
 * Total Submissions: 394.2K
 * Testcase Example:  '[[byte('1'),byte('1'),byte('1'),byte('1'),"0"],[byte('1'),byte('1'),"0",byte('1'),"0"],[byte('1'),byte('1'),"0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 
 * 此外，你可以假设该网格的四条边均被水包围。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [
 * ⁠ [byte('1'),byte('1'),byte('1'),byte('1'),"0"],
 * ⁠ [byte('1'),byte('1'),"0",byte('1'),"0"],
 * ⁠ [byte('1'),byte('1'),"0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [
 * ⁠ [byte('1'),byte('1'),"0","0","0"],
 * ⁠ [byte('1'),byte('1'),"0","0","0"],
 * ⁠ ["0","0",byte('1'),"0","0"],
 * ⁠ ["0","0","0",byte('1'),byte('1')]
 * ]
 * 输出：3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 
 * grid[i][j] 的值为 '0' 或 '1'
 * 
 * 
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j] == byte('1'){
				dfs(i,j,grid)
				count++
			}
		}
	}
	return count
}
// 深度优先的思路
func dfs(i,j int,grid [][]byte){
	grid[i][j] = 0
	if i-1>=0 && grid[i-1][j] == byte('1'){
		dfs(i-1,j,grid)
	}
	if i+1<len(grid) && grid[i+1][j] == byte('1'){
		dfs(i+1,j,grid)
	}
	if j-1>=0 && grid[i][j-1] == byte('1'){
		dfs(i,j-1,grid)
	}
	if j+1<len(grid[i]) && grid[i][j+1] == byte('1'){
		dfs(i,j+1,grid)
	}
}
// @lc code=end

