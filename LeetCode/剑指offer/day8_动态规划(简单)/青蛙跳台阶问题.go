//题目链接：https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

package main

import "math"


//本质上和斐波那契数列一样，区别在于 dp 数组初始化不同。
//跳到第 n 级台阶，我们可以从第 n-1 个台阶跳上去，也可以从 n-2 个台阶跳上去，跳法，即为跳到 n-1 级台阶 和 跳到 n-2 级台阶的跳法之和。
//动态规划三步：
//1. 确定dp数组及下表含义：dp[i] 代表跳到第 i 级台阶的跳法数量，求跳到第 n 级台阶的跳法数量，dp[0]情况特殊，则 dp 数组长度设为 n+1；
//2. dp 数组初始化：由题意可知，dp[0]=1,dp[1]=1,dp[2]=2
//3. 状态转移方程：当 i > 2 时，dp[i] = dp[i-1] + dp[i-2]
//最终，返回 dp[n] 即可。
func numWays(n int) int {
	if n <= 1{
		return 1
	}
	x := int(math.Pow(10,9)+7)
	dp := make([]int,n+1)
	dp[0] = 1
	dp[1] = 1
	for i:=2;i<=n;i++{
		dp[i] = (dp[i-1]+dp[i-2]) % x
	}
	return dp[n]
}

//由于每个阶段的状态只与前两个状态有关，所以我们可以用滚动数组代替 dp 数组解题，
//将空间复杂度从 O(n) 降低至 O(1)。
func numWays_2(n int) int {
	if n <= 1{
		return 1
	}
	x := int(math.Pow(10,9)+7)
	a,b := 1,1
	var res int
	for i:=2;i<=n;i++{
		res = (a+b) % x
		a,b = b,res
	}
	return res
}