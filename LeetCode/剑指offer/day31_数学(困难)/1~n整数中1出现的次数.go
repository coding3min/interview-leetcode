//题目链接：https://leetcode.cn/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/?envType=study-plan&id=lcof
// day31/31
// 第 31 天主题为：数学（困难）
// 包含三道题目：
// 剑指offer14.剪绳子 II
// 剑指offer43.1~n 整数中 1 出现的次数
// 剑指offer44.数字序列中某一位的数字

package main

//这道题看的官方题解，数学题么，主要还是找规律和做总结。
//
//解题思路为统计各位上1出现的次数。设输入数字 x 为 n 位数，描述为 x=xn_xn−1_xn−2...x2_x1
//若当前统计为 x3 位，高位为 xn_xn−1_xn−2...x4 ，记为 high，低位为 x2_x1 ，记为 low
//若数字 x 在 x3 位上的数字为 m，当前位上的因子 factor 为 10^(3-1) ,根据 m 的大小分情况讨论：
//
//- m=0,则该位上1出现的次数为 high * factor
//- m=1,则该位上1出现的次数为 high * factor + low + 1
//- m>1,则该位上 1 出现的次数为 (high+1) * factor
//
//问题来了，这种情况讨论是如何总结出来的呢，我觉得，就是单纯凭感觉找规律吧
//为编程方便，从个位，也就是最低位开始统计1出现的次数。

func countDigitOne(n int) int {
	res := 0
	// digits为数字n的切片表示形式，第一个元素为个位
	digits := []int{}
	temp := n
	for temp!=0{
		digits = append(digits,temp%10)
		temp /= 10
	}
	// 为编程方便，从个位开始统计1出现的次数
	// high初始化为n/10，low的位数为0，初始化为0
	high,low := n/10,0
	// 个位的factor为10的0次方
	factor := 1
	// 从个位开始统计每位上1出现的次数，分三种情况讨论
	for i:=0;i<len(digits);i++{
		if digits[i] == 0{
			res += high * factor
		} else if digits[i] == 1{
			res = res + high*factor + low+1
		} else {
			res += (high+1) * factor
		}
		// 更新high和low
		high,low = high/10,digits[i]*factor+low
		// 更新factor，为上一次因子乘以10
		factor *= 10
	}
	return res
}