//题目链接：https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/
package main

import "strconv"

//这是非常考验边界问题处理的一道题目，解题思路不难想到，但要想ac，非常有难度。
//题解参考自：https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/
//
//解题分为三步：
//1. 确定 第n位 所在数字的数字长度digit，比如，个位数 3 的数字长度为1，百位数 202 的数字长度为3
//2. 确定 第n位 所在的数字 num，比如三位数的第二个数字为 101
//3. 确定 第n位 在数字 num 的第几位，比如三位数 101 的第二位数字为 0
//
//一步一步来
//对每个数字长度范围的数字规律总结
//| 数字范围  | 单个数字长度 | 数字数量  | 该范围所有数字长度之和 |
//| --------- | ------------ | --------- | ---------------------- |
//| 1~9       | 1            | 9         | 9                      |
//| 10~99     | 2            | 90        | 180                    |
//| 100~999   | 3            | 900       | 2700                   |
//|           |              |           |                        |
//| start~end | digit        | 9 * start | 9 * start * digit      |

//由上表，可得
//数字长度递归公式：digit = digit + 1
//起始数字递归公式：start = start * 10
//该范围内所有数字长度之和：count = 9 * start * digit

//一、确定 第n位 所在的数字长度
//	循环执行 n 减去一位数、二位数、...、的数字长度之和，直至 n<=count 跳出

//二、确定 第n位 所在的数字
// 	这里建议看一下题解开头附上的链接，有讲解的图片，我这里直接敷上公式
//  所求数位 在从数字 start 开始的第 [(n - 1) / digit] 个 数字 中（ start 为第 0 个数字）。

//三、确定 第n位 在该数字的哪一位，很明显这里是用取余操作
//  第n位在 num 的第 (n-1)%2 位

func findNthDigit(n int) int{
	// digit为数字长度，start为该长度数字的第一个数字
	// count为该长度数字占据的位数
	digit,start,count := 1,1,9
	// 1.确定n所在数字的数字长度
	for n > count{
		n -= count
		start *= 10
		digit ++
		count = 9 * start * digit
	}
	// 确定n所在的数字num
	num := start + (n-1)/digit
	str := strconv.Itoa(num)
	// 确定n所在数字num的哪一位
	res,_ := strconv.Atoi(string(str[(n-1)%digit]))
	return res
}