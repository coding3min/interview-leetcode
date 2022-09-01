// 题目链接：https://leetcode.cn/problems/ti-huan-kong-ge-lcof/?envType=study-plan&id=lcof

// 先说点题外话：为什么要替换空格，在网络编程中，如果URL参数含有特殊字符，如空格、'#'等，
// 可能导致服务端无法获得正确的参数值。我们需要将这些特殊符号转换成服务器可以识别的字符。
// 转换规则是在百分号后跟上ASCII码的两位十六进制的表示。比如空格的ASCII码是32，即十六进制的0x20，因此空格被替换成“%20”
//
// 回到本题：字符串的替换，设返回字符串为res，输入字符串s的长度为n，首先我们要统计s中空格的数量，
// 用来计算res的长度，每出现一次空格，字符串的长度+2，设 len(res)=m
// 然后准备一次遍历，使用两个指针 i 和 j 分别指向 s 和 res 当前遍历到的位置下标
// 指针从前往后或从后往前都可以，这里先从后往前，那么 i 和 j 分别初始化为 n-1 和 m-1，
// 当指针 i 指向的位置对应字符为空格时，我们需要在 j 指向的位置及前两个位置插入“%20”，之后，
// i向前移动1格，j向前移动3格；若i对应位置不是空格，res[j]=s[i]即可，i和j各自向前移动1格。

package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	num_Space := 0
	for _,x := range s{
		if x == ' '{
			num_Space ++
		}
	}
	m := len(s) + 2 * num_Space
	fmt.Println(len(s),m)
	res := make([]rune,m)
	j := m - 1
	for i:=len(s)-1;i>=0;i--{
		if s[i]==' '{
			res[j-2] = '%'
			res[j-1] = '2'
			res[j] = '0'
			j -= 3
		} else {
			res[j] = rune(s[i])
			j --
		}
	}
	return string(res)
}

// 若遍历方向改为从前往后，同理
// 为避免命名冲突，此函数名添加后缀 “_2"
func replaceSpace_2(s string) string {
	nums := 0
	for _,x := range s{
		if x == ' '{
			nums ++
		}
	}
	m := len(s) + 2 * nums
	res := make([]rune,m)
	j := 0
	for i:=0;i<len(s);i++{
		if s[i]==' '{
			res[j] = '%'
			res[j+1] = '2'
			res[j+2] = '0'
			j += 3
		} else {
			res[j] = rune(s[i])
			j ++
		}
	}
	return string(res)
}

// 另外，还可以皮一下，直接调用库函数api
func replaceSpace_3(s string) string {
	res := strings.Replace(s," ","%20",-1)
	return res
}

// 最后，我们还应该思考一下从前往后遍历和从后往前遍历这两种解法的区别
// 设想这样一种场景，题目要求原地修改输入的参数，并且保证改参数有充足的空间保存修改后的元素，并且仅允许使用 O(1) 的空间复杂度。
// 这时从前向后遍历解法就不再生效了，而从后往前依然可以使用。这种从后往前遍历的解法是一类题目的解决方案，最好能牢记于心。
// 另外，由于字符串本身是不可修改的，所以这里谈到的输入参数不会再是字符串，可以为 rune切片，int切片等。