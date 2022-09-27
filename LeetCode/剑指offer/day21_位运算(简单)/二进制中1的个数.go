//题目链接：https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
package main

//方法1：逐位判断，循环检查每一个二进制位是否为 1
//对 Golang，题目输入为一个 uint32 类型的数字，但它实际代表的是一个数的二进制表示形式。
//我们可以进行 32 次的循环，判断每一位是否为 1 。
func hammingWeight(num uint32) int {
	res := 0
	for i:=0;i<32;i++{
		if num & 1 == 1{
			res ++
		}
		num >>= 1
	}
	return res
}

//方法2：位运算优化，消除二进制末尾的 1
//非常巧妙的做法，当 n 非零时，n=n&(n-1)可消除n的二进制中最后一个出现的 1.
//因此，执行 n=n&(n-1)使得n变成 0 的操作次数，就是 n 的二进制中 1 的个数。
func hammingWeight_2(num uint32) int {
	res := 0
	for num != 0{
		res ++
		num &= num-1
	}
	return res
}

//第一种解法的时间复杂度为 O(logn)，第二种解法为 O(m)，m 为二进制串中 1 的个数；
//两种解法的空间复杂度均为 O(1)。
//
//个人感觉啊，位运算优化只是一种技巧，在效率方面对比逐位检查提升并不大。
//因为逐位检查已经是对数级别了，对数级别的时间复杂度已经是相当高效。