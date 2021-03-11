/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (38.35%)
 * Likes:    820
 * Dislikes: 0
 * Total Accepted:    102.4K
 * Total Submissions: 267.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 * 
 * 示例:
 * 
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 * 
 * 
 * 说明:
 * 
 * 假设你总是可以到达数组的最后一个位置。
 * 
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	// 和跳跃1不同的是这里要求最少跳跃次数，要多记录一个次数
	// 使用备忘录来记录最少次数
	dp := make([]int,n)
	// 状态：当前方格步数最少
	// 选择：上一次到达dp(last) 和这一次到达 dp(now) + 1
	// 
	/* 	状态转移方程：(正因为)
		dp(new) = min(dp[now] + 1,dp[old]) , now = curr+1,curr=i+j
	*/
	// 最后一格是终点，所以只用计算到倒数第二格
	for i:=0;i<n-1;i++{
		// 每次都给新的那一格赋值最小次数
		for j:=0;j<nums[i] && i+j+1<n;j++{
			// 因为第一次给0更新后就是最小的，所以都不用比较了
			if dp[i+j+1]==0{
				// 当前位置+1就是使用当前能量跳最少步数
				dp[i+j+1] = dp[i]+1
			}
		}
	}
	return dp[n-1] 
}
// @lc code=end

