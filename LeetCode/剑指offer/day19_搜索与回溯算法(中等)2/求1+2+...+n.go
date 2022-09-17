//题目链接：https://leetcode.cn/problems/qiu-12n-lcof/?envType=study-plan&id=lcof
package main

//先不考虑限制条件，有三种解法，
//1.高斯求和公式（使用乘除法）
//2.迭代（使用for或while）
//3.递归（递归终止条件需要使用if）。
//
//考虑限制条件的话，难点在于如何实现这个 1~n 的循环。
//我们可以利用运算符的短路效应实现该循环
//
//先看下逻辑与 与 逻辑或 的短路效应
//- if a && b； 若 a 为 false，对 b 的判断不会执行（即短路），直接返回 false
//- if  a || b； 若 a 为 true，对 b的判断不会执行，直接返回 true
//
//我们可以利用这种短路机制，通过递归实现 1~n 的相加
func sumNums(n int) int {
	res := 0
	var f func(*int,int) bool
	f = func(res *int,n int) bool{
		*res += n
		return n > 0 && f(res,n-1)
	}
	f(&res,n)
	return res
}