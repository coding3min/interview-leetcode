//题目链接：https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/?envType=study-plan&id=lcof

package main

import (
	"strings"
)

//我的想法是使用栈，每遇到一个完整的单词，添加到栈中，最后出栈。
//
//具体过程：一次遍历，用 string.Builder 构建字符串，若当前字符不是空格，则写入将当前字符写入 sb，
//若是，则将 sb 的内容写入栈 并 清空内容，准备下一个单词的填充。
func reverseWords(s string) string {
	var sb strings.Builder
	n := len(s)
	if n == 0{
		return ""
	}
	stack := []string{}
	for i:=0;i<n;i++{
		if s[i]==32 && i>0 && s[i-1]!=32{
			stack = append(stack,sb.String())
			sb.Reset()
		} else if s[i]!=32{
			sb.WriteByte(s[i])
		}
	}
	if s[n-1] != 32{
		stack = append(stack,sb.String())
	}
	sb.Reset()
	for i:=len(stack)-1;i>=0;i--{
		sb.WriteString(stack[i])
		sb.WriteByte(' ')
	}
	res := sb.String()
	// 去掉末尾空格
	// 同时注意特殊情况：s全为空格，res长度为0，对其进行去除末尾空格操作会索引越界
	if len(res)==0{
		return ""
	}
	res = res[:len(res)-1]
	return res
}


//还可以用双指针：
//
//声明变量 res 为空字符串，left 和 right 两指针初始化为 0，左右指针分别用于寻找每个单词的左右边界。
//一次遍历字符串，左指针向右移动寻找第一个非空格字符，为要寻找字符串的左边界，然后开始寻找右边界，先将 left 的值赋给 right，然后 right 向右移动，
//直到到达字符串末尾或者寻找到空格字符，此时 left 和 right 分别指向一个单词的左右边界，将其加入 res，然后将 right 赋给 left，开始寻找下一个单词。
//
//具体实现时，注意每加入一个单词的同时，还要添加一个空格，最后返回 res 前，处理掉最后添加的空格。
//
//注意特殊情况的处理，当输入字符串全为空格时，res 长度为空，此时如果对其进行去除空格的操作，会索引越界。
// 方法2：
func reverseWords_2(s string) string {
	n := len(s)
	left,right := 0,0
	res := ""
	for left < n{
		if s[left] == 32{
			left ++
		} else {
			right = left
			for right<n && s[right]!=32{
				right ++
			}
			res = s[left:right] + " " + res
			left = right
		}
	}
	if len(res) == 0{
		return ""
	} else {
		return res[:len(res)-1]
	}
}