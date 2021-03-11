/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (53.37%)
 * Likes:    2928
 * Dislikes: 0
 * Total Accepted:    425.3K
 * Total Submissions: 796.5K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1]
 * 输出：1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：nums = [-1]
 * 输出：-1
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：nums = [-100000]
 * 输出：-100000
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^5 
 * 
 * 
 * 
 * 
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 * 
 */

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums)
	if n<1{
		return 0
	}
	last,resMax := nums[0],nums[0]
	for i:=1;i<n;i++{
		// 要不就是自成一派，要不就是和前面连续的放在一起
		last = max(nums[i],nums[i]+last)
		resMax = max(resMax,last)
	}
	return resMax
}
func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}

// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	if n<1{
// 		return 0
// 	}
// 	dp := make([]int,n)
// 	dp[0] = nums[0]
// 	for i:=1;i<n;i++{
// 		// 要不就是自成一派，要不就是和前面连续的放在一起
// 		dp[i] = max(nums[i],nums[i]+dp[i-1])
// 	}
// 	res := math.MinInt32
// 	for _,v := range dp{
// 		res = max(res,v)
// 	}
// 	return res
// }
// func max(x,y int)int{
// 	if x>y{
// 		return x
// 	}
// 	return y
// }
// @lc code=end

