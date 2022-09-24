//题目链接：https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/?envType=study-plan&id=lcof
// day26/31
// 第 26 天主题为：字符串（中等）
// 包含两道题目：
// 剑指offer20.表示数值的字符串
// 剑指offer67.把字符串转换成整数

package main

// 模拟即可
// 官方题解的解法是：确定有限状态自动机，这块我不太懂，有兴趣的同学自行了解哈
func isNumber(s string) bool {
	n := len(s)
	// 若字符串长度为0，无法表示数值
	if n == 0{
		return false
	}
	// 从下标0开始遍历
	index := 0
	// 读取字符开头处的若干空格
	for index < n && s[index] == ' '{
		index ++
	}
	// 读取整数部分,，numeric代表读取到index是否为数值
	// index 持续向前推进
	numeric := ScanInteger(s,&index)
	// 若下一个字符为'.'，说明为小数，读取小数部分
	if index < n && s[index] == '.'{
		index ++
		// 这里用逻辑或的原因：
		// 1. '.'前若无数字，则'.'后至少需要一位数字
		// 对应左表达式为 true，右表达式为 false
		// 2.'.'前有数字，则后方可以有也可以没有数字
		// 对应左true右true 和 左false右true
		numeric = ScanUnsignedInteger(s,&index) || numeric
	}
	// 若出现e，后面需要跟一个整数，原理同上
	if index < n && (s[index] == 'e' || s[index] == 'E'){
		index ++
		numeric = numeric && ScanInteger(s,&index)
	}
	// 读取字符串结尾处的若干空格
	for index < n && s[index] == ' '{
		index ++
	}
	// 若 numeric 为 true，且读取到字符串末尾，说明为数值
	return numeric && index == len(s)
}

// 读取字符串s从下标index开始的有符号整数
func ScanInteger(s string,index *int) bool{
	// 先读取正负号，若无符号，默认正数
	if *index < len(s) && (s[*index] == '+' || s[*index] == '-'){
		*index ++
	}
	// 读取符号后，读取无符号整数
	return ScanUnsignedInteger(s,index)
}

// 读取字符串s从下标index开始的无符号整数
// 若未读取到无符号整数，返回false，否则，返回true
func ScanUnsignedInteger(s string,index *int) bool {
	start_of_integer := *index
	// 根据ASCII码判断是否为数字字符
	// 若读取到非数字字符，直接跳出即可
	for *index < len(s) && s[*index] >= '0' && s[*index] <= '9'{
		*index ++
	}
	// 若*index大于读取前的值，说明有整数被读取到
	return *index > start_of_integer
}