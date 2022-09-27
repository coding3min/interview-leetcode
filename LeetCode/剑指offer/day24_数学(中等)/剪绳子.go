//题目链接：https://leetcode.cn/problems/jian-sheng-zi-lcof/?envType=study-plan&id=lcof
// day24/31
// 第 24 天主题为：数学（中等）
// 包含三道题目：
// 剑指offer14-I.剪绳子
// 剑指offer57-II.和为s的连续正数序列
// 剑指offer62.圆圈中最后剩下的数字

package main

import "math"

//先上传统的动态规划解法
//
//1. 确定dp数组大小及下标含义：dp[i] 代表长度为 i 的绳子切割后的最大乘积，长度为 n+1
//2. dp 数组初始化：dp[1]=1,dp[2]=2,dp[3]=3,需要注意的是，该初始值只针对 n>=4 的情况，n<4的情况，我们会单独处理
//3. 状态转移方程：dp[i] = max(dp[i] * dp[i-j])，j 从 1 遍历到 i-1，这里我们可以使用一个小技巧，j 从 1 遍历至 n/2 即可，因为 1*6 = 6 * 1
//
//再说下 n<4 的情况，题目给定 n>=2，当 n<4 时，返回 n-1 即可。

func cuttingRope(n int) int {
	if n < 4{
		return n-1
	}
	dp := make([]int,n+1)
	dp[1],dp[2],dp[3] = 1,2,3
	for i:=4;i<=n;i++{
		for j:=1;j<=i/2;j++{
			dp[i] = max(dp[i],dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

//这道题还有一种数学的解法，时间和空间复杂度都超过dp
//
//公式我在纸上推导，拍照放在了同一文件夹中，剪绳子.jpg

func cuttingRope_2(n int) int {
	if n <= 3{
		return n-1
	}
	a,b := n/3,n%3
	if b == 0{
		return int(math.Pow(3,float64(a)))
	} else if b == 1{
		return int(math.Pow(3,float64(a-1)) * 4)
	} else {
		return int(math.Pow(3,float64(a)) * 2)
	}
}