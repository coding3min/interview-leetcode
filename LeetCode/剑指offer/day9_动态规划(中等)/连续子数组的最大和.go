// 题目链接：https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/?envType=study-plan&id=lcof
package main

//简单的一维 dp 入门题目
//动态规划三步：
//1. 确定dp数组大小及下标含义：dp[i] 代表以下标 i 为最后元素的子数组的最大值，dp 数组长度与给定数组 nums 长度相同。
//2. dp 数组初始化：每个单独元素都是一个子数组，初始化 dp[0] = nums[0]
//3. 状态转移方程：从下标 1 开始，dp[i] = max(nums[i],nums[i]+dp[i-1])，若 dp[i-1]>0，说明以下标 i 截止的子数组的最大和要包含之前元素，包含多少呢？dp[i-1] 已经处理完了，我们只需要相加即可。
// 最终需返回 dp 数组的最大值
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int,n)
	dp[0] = nums[0]
	res := nums[0]
	for i:=1;i<n;i++{
		dp[i] = max(nums[i],nums[i]+dp[i-1])
		res = max(res,dp[i])
	}
	return res
}