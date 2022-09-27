//题目链接：https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/?envType=study-plan&id=lcof
// day25/31
// 第 25 天主题为：模拟（中等）
// 包含两道题目：
// 剑指offer29.顺时针打印矩阵
// 剑指offer31.栈的压入、弹出序列
package main

//简单题，模拟该过程，声明一个栈 stack 和 变量 i（i 符合条件的出栈元素数量），遍历压入序列 pushed，依次入栈，每次入栈结束后
//判断栈顶元素是否和出栈序列 popped[i] 相等
//若相等，出栈，i++，再次判断栈顶元素和 popped[i]是否相等...直至栈长度为0或者栈顶元素和popped[i]不相等。
//
//遍历完成后，若 i = len(popped)，说明所有元素出栈成功，返回 true。判断条件改为 len(stack)=0 也可以，道理相同
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _,num := range(pushed){
		stack = append(stack,num)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i]{
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}