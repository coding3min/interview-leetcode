// 题目链接：https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/?envType=study-plan&id=lcof
// day9/31
// 第 9 天主题为：动态规划（中等）
// 包含两道题目：
// 剑指offer42.连续子数组的最大和
// 剑指offer47.礼物的最大价值

package main

//经典二维dp，在每个网格的位置有两个选择，向右 或 向下走，当前网格可获得的最大礼物价值，
//就只能从其上面的格子或左边的格子 的最大礼物价值 加上当前格子的礼物价值 得到，回溯到起始节点，容易想到用 dp 解题。
//动态规划三步骤：
//1. 确定dp数组大小及下标含义：dp[i][j] 代表 从棋盘左上角走到 (i,j) 下标位置可以获得的礼物最大价值，
//   则 len(dp) = len(grid),len(dp[0])=len(grid[0])，即 大小与给定棋盘大小相等；
//2. dp数组初始化：（一般情况下初始第一行和第一列）网格（0,0）为起始位置，dp[0][0] 没有别的选择，dp[0][0] = grid[0][0]
//   因为计算当前网格的最大礼物价值，需要知道其上方和左方网格的最大礼物价值，所以我们要初始化第一行和第一列的 dp 数组元素，防止越界情况的发生；
//3. 状态转移方程：i>1 且 j>1 时：dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + grid[i][j]
//
//最后返回 dp[m-1][n-1] 即可。


func maxValue(grid [][]int) int {
	if len(grid)==0 || len(grid[0])==0{
		return 0
	}
	m,n := len(grid),len(grid[0])
	dp := make([][]int,m)
	for i:=0;i<m;i++{
		dp[i] = make([]int,n)
	}
	dp[0][0] = grid[0][0]
	for i:=1;i<m;i++{
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i:=1;i<n;i++{
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func max(x,y int) int {
	if x > y{
		return x
	}
	return y
}