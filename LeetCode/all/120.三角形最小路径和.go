/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (67.20%)
 * Likes:    711
 * Dislikes: 0
 * Total Accepted:    135K
 * Total Submissions: 200.8K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 * 
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：triangle = [[-10]]
 * 输出：-10
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 * 
 * 
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	n := len(triangle) 
	if n == 0{
		return 0
	}
	dp := make([]int,n+1)
	// 倒序遍历行
	for i:=n-1;i>=0;i--{
		// 正序遍历列
		for j:=0;j<len(triangle[i]);j++{
			//第一次遍历完最后一行时，dp完成了赋值，值为最后一行
			//遍历倒数第二行时，因为要求取相邻，所以比较下一行的dp[j]和dp[j+1]
			//因为相邻是指下一行的下标 i 或 i + 1，多申请一个元素就不用判断越界了，倒数第二行就收缩了
			dp[j] = min(dp[j],dp[j+1]) + triangle[i][j]
		}
	}
	// 最后只剩下顶端元素
    return dp[0]
}
func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}
// @lc code=end

