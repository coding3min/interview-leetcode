// 题目链接：https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/?envType=study-plan&id=lcof
// day7/31
// 第 7 天主题为：搜索与回溯算法（简单），与前一天主题相同
// 包含三道题目：
// 剑指offer10-I.斐波那契数列
// 剑指offer10-II.青蛙跳台阶问题
// 剑指offer63.股票的最大利润

//关于函数命名冲突的问题，从今天的题解开始，就不再额外标注了
//我一般通过函数名添加后缀的方式解决

package main

import "math"

//解题思路：最基础的动态规划问题，题目已经给出状态转移方程，或者可以直接理解为 递归，一个阶段只有一个状态
//动态规划三步：
//1. 确定dp数组及下标含义：dp[i] 代表 F(i)，要求第 n 个斐波那契数，则 dp 数组长度为 n；
//1. 数组初始化：题目已给出，dp[0]=0,dp[1]=1；
//2. 状态转移方程：题目已给出，i>=2时，dp[i] = dp[i-1]+dp[i-2]。
//最终返回 dp[n-1] 即可

func fib(n int) int {
	if n <= 1{
		return n
	}
	x := int(math.Pow(10,9)) + 7
	dp := make([]int,n+1)
	dp[0],dp[1] = 0,1
	for i:=2;i<n+1;i++{
		dp[i] = (dp[i-1] + dp[i-2]) % x
	}
	return dp[n]
}


//由于每个阶段的状态只与前两个状态有关，所以我们可以用滚动数组代替 dp 数组解题，
//将空间复杂度从 O(n) 降低至 O(1)。
func fib_2(n int) int {
	if n <= 1{
		return n
	}
	x := int(math.Pow(10,9)) + 7
	a,b := 0,1
	var res int
	for i:=2;i<n+1;i++{
		res = (a+b) % x
		a,b = b,res
	}
	return res
}