/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (50.93%)
 * Likes:    675
 * Dislikes: 0
 * Total Accepted:    91.1K
 * Total Submissions: 178.8K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
 * 
 * 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
 * 
 * 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入：prices = [3,3,5,0,0,3,1,4]
 * 输出：6
 * 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
 * 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：prices = [1,2,3,4,5]
 * 输出：4
 * 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4
 * 。   
 * 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
 * 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：prices = [7,6,4,3,1] 
 * 输出：0 
 * 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
 * 
 * 示例 4：
 * 
 * 
 * 输入：prices = [1]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	/*  
		状态： 买1、卖1，买2，卖2
		子问题：每次买卖的最大收益
		选择：
		当前节点买 when 不存在买
		当前节点卖 when 存在已买
		状态转移方程:
		0 if len<2
		dp 用来存储收益
		dp[0] = -prices[0]
		dp[0] = max{dp[0],-prices[i]} 最少钱买入，收益为负
		dp[1] = max{dp[1],dp[0]+prices[i]} 减去成本，比较收益
		dp[2] = max{dp[2],dp[1]-prices[i]} 第二次买入，减去成本最大
		dp[3] = max{dp[3],dp[2]+prices[i]} 第二次卖出，总收益最大
		存在重复买入卖出，其实相当于同一个买入卖出就是没有任何操作的情况
	*/
	if n < 2{
		return 0
	}
	dp := make([]int,4)
	dp[0] = -prices[0] 
	dp[2] = -prices[0]
	for _,v := range prices{
		dp[0] = max(dp[0],-v)
		dp[1] = max(dp[1],dp[0]+v)
		dp[2] = max(dp[2],dp[1]-v)
		dp[3] = max(dp[3],dp[2]+v)
	}
	return dp[3]
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
// @lc code=end

