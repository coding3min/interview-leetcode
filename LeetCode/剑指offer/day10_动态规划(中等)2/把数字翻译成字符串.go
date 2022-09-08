// 题目链接：https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/?envType=study-plan&id=lcof
// day10/31
// 第 10 天主题为：动态规划（中等）
// 包含两道题目：
// 剑指offer46.把数字翻译成字符串
// 剑指offer48.最长不含重复字符的子字符串
package main

import "strconv"

//比较明显的动态规划题目，设 dp[i]代表s[:i+1]的翻译方法数目，dp[i] 明显依赖于 dp[i-1] 与 dp[i-2]，类似于斐波那契数是吧，
//本题有点不一样的地方在于 dp[i]=dp[i-1]+dp[i-2] 是有条件的
//- 当 s[i-1] 与 s[i] 组成的数字小于 25 时，s[i] 可以与 s[i-1] 组合翻译，也可以分开翻译，dp[i]=dp[i-1]+dp[i-2]，
//- 否则 s[i] 与 s[i-1] 无法组合翻译，只能单独翻译，dp[i] = dp[i-1]

//动态规划三步骤：
//- 确定dp数组大小及下标含义：dp[i] 代表 s[:i+1] 的翻译方法数目，len(dp)=len(string(num))
//- dp 数组初始化：dp[0]对应s[0]，单个字符只有一种翻译方法，dp[0]=1，当 s[:2] 小于26 且 s[i]!=0 时，dp[1]=2，否则 dp[1]=1
//- 状态转移方程：从下标 2 开始遍历，x = strconv.Atoi(s[i-1:i+1])，并且判断 s[i-1] 是否为 0
//- 	若 x < 26 且 s[i-1]!= 0，dp[i]=dp[i-1]+dp[i-2]，s[i] 可以与 s[i-1] 组合翻译，也可以单独翻译
//- 	否则，dp[i] = dp[i-1]，s[i] 只能单独翻译，s[:i+1] 的翻译方法数目依赖于 s[:i]
// 最后，返回 dp[n-1] 即可。

//有一点需要注意，在状态转移方程那里一定要判断 s[i-1] 是否为 0，因为 “01”只能翻译为 “ab”，不能翻译成 “a"，
//我第一次做这道题的时候就是忘记判断 s[i-1]是否为0，导致没有ac，刚才做的时候，又在这个地方跌到坑里了，有了再一再二，不会有再三再四了。
func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(string(s))
	if n == 0{
		return 0
	}
	dp := make([]int,n)
	dp[0] = 1
	x,_ := strconv.Atoi(s[:2])
	zero,_ := strconv.Atoi(string(s[0]))
	if x < 26 && zero != 0{
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i:=2;i<n;i++{
		x,_ = strconv.Atoi(s[i-1:i+1])
		zero,_ := strconv.Atoi(string(s[i-1]))
		if x < 26 && zero!= 0{
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}

