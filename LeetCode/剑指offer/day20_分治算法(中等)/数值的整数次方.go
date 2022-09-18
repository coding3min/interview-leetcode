//题目链接：https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/?envType=study-plan&id=lcof


//这道题，首先要和面试官沟通关于数学上的定义，0的0次方可以是0，也可以是1，和面试官交流后再写代码，确定这种情况下返回0还是1 。
//然后处理 n<0 的情况，n小于0时，x取1/x（为防止分母为0，提前对x等于0的情况进行处理），n取其相反数，将n化为正数，可统一化算法流程。
//通俗解法就是循环 n 次，每次循环返回变量对 x 作乘积。这种做法提交后会超时！
//考虑算法：快速幂
//迭代解法的时间复杂度为 O(n)，快速幂可将时间复杂度降低至 O(logn)。

//咱从二进制角度分析快速幂算法
//设 n 的二进制长度为 m+1，则 n 的二进制表示和二进制转十进制表示如下：
//我之前使用 latex 编辑的，这里大家理解一下
//$$
//n=(0b)b_mb_{m-1}...b_1b_0=b_m*2^m+b_{m-1}*2^{m-1}+...+b_1*2^{1}+b_0*2^0
//$$
//则 x 的 n 次方可展开为：
//$$
//x^n=x^{b_m*2^m+b_{m-1}*2^{m-1}+...+b_1*2^{1}+b_0*2^0}=x^{b_m*2^m}*...*x^{b_0*2^0}
//$$
//那我们求 x 的 n 次方就可以转化为这 m+1 项的乘积，对这 m+1 项，从后向前，依次计算
package main

func myPow(x float64, n int) float64 {
	if x == 0{
		return x
	}
	var res float64 = 1
	if n < 0{
		x = 1/x
		n = -n
	}
	// 因子初始化为 x
	factor := x
	for n != 0{
		// 若当前因子的项为1，更新res
		// 否则，res不变
		if n & 1 == 1{
			res *= factor
		}
		// 之后每次 factor = factor^2
		factor *= factor
		// n右移一位
		n >>= 1
	}
	return res
}