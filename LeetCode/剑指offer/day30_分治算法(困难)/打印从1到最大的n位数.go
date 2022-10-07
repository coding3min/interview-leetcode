//题目链接：https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
// day30/31
// 第 30 天主题为：分治算法（困难）
// 包含两道题目：
// 剑指offer17.打印从1到最大的n位数
// 剑指offer51.数组中的逆序对

package main

import (
	"math"
	"strconv"
	"strings"
)

//题目描述为返回整数列表，那通过不断迭代，从1，不断加至最大的n位数，依次添加至结果列表，返回即可。
func printNumbers(n int) []int {
	res := []int{}
	limit := math.Pow(10,float64(n))
	x := 1
	for x < int(limit){
		res = append(res,x)
		x += 1
	}
	return res
}


//在《剑指offer》中，需要打印字符串数组，因为题目没有规定 n 的范围，当 n 非常大的时候，最大的 n 位数会超出整型的范围，
//所以说，一定要考虑到大数问题，这才是本题的考点。大数问题的解决办法就是用字符串表达大数。
//
//最直观的解法是用字符串模拟数字加法，另外，也容易想到，返回的字符串数组为 n 个 0-9 的全排列。全排列就要用到递归。
//基于分治的思想，先确定高位，依次向低位移动。例如 n=2 时，先确定十位，再确定个位。
//最后需要处理每个字符串的前置零，数组的第一个元素为全零字符串，也需要删除，代码注释很详细了，如下：
func printNumbers_2(n int) []int {
	// 先按照返回字符串数组的思路进行解题
	// strs 存储回溯得到的字符串数组（未处理前置零）
	strs := []string{}
	var backtrace func(path string)
	backtrace = func(path string){
		// 字符串长度达到n时，将字符串添加至strs
		// 结束递归
		if len(path) == n {
			strs = append(strs,path)
			return
		}
		// 从 0-9 迭代
		for i:=0;i<10;i++{
			// 将当前数字的字符添加至 path
			path += strconv.Itoa(i)
			backtrace(path)
			// 回溯，删除最新添加的字符
			path = path[:len(path)-1]
		}
	}
	backtrace("")
	// 删除数组第一个全零字符串
	strs = strs[1:]
	k := len(strs)
	// delZeroStrs 用来保存去除前置零的字符串数组
	delZeroStrs := make([]string,k)
	for i:=0;i<k;i++{
		delZeroStrs[i] = delZero(strs[i])
	}
	// 最后，为符合本题要求，转化为整数数组
	res := make([]int,k)
	for i:=0;i<k;i++{
		num,_ := strconv.Atoi(delZeroStrs[i])
		res[i] = num
	}
	return res

}

// delZero 用于删除输入字符串 s 前置零
func delZero (s string) string{
	for i:=0;i<len(s);i++{
		if s[i] != '0'{
			return s[i:]
		}
	}
	return s
}

//下述内容是关于回溯写法相关的
//对 printNumbers 函数中的 backtrace 函数，我们输入参数为 path，调用时传入空字符串，然后回看代码的时候发现，
//提前声明一个 path 变量，为空字符串，然后回溯函数好像就不需要参数列表了

//做了尝试，发现完全没问题
func printNumbers_3(n int) []int {
	strs := []string{}
	path := ""
	var backtrace func()
	backtrace = func(){
		if len(path) == n {
			strs = append(strs,path)
			return
		}
		for i:=0;i<10;i++{
			path += strconv.Itoa(i)
			backtrace()
			path = path[:len(path)-1]
		}
	}
	backtrace()
	strs = strs[1:]
	k := len(strs)
	delZeroStrs := make([]string,k)
	for i:=0;i<k;i++{
		delZeroStrs[i] = delZero(strs[i])
	}
	res := make([]int,k)
	for i:=0;i<k;i++{
		num,_ := strconv.Atoi(delZeroStrs[i])
		res[i] = num
	}
	return res
}

//另外，如果我们的 path 是一个 string 切片的话，我们就不再需要回溯操作，每次迭代操作会将上一次迭代的字符覆盖掉，变成一个 dfs 。
//这段代码是我挺久之前写的，注释就不修改了，可以当参考看一下
func printNumbers_4(n int) []int {
	// 先按照返回字符串数字思路进行解题
	ans := []string{}
	var dfs func(x int)
	// temp存储当前字符串的字节切片（代码用字符串切片存储）
	temp := make([]string,n)
	dfs = func(x int) {
		if x == n{
			// 达到n位时，递归结束，
			// 组合temp添加至ans切片末尾
			ans = append(ans,strings.Join(temp,""))
			return
		}
		// 每一位从0-9进行递归
		for i:=0;i<10;i++{
			temp[x] = strconv.Itoa(i)
			// 当前位处理完成后，进行下一位递归处理
			dfs(x+1)
		}
	}
	// 从第0位开始递归
	dfs(0)
	// 去除每个字符串的前置零
	for i:=0;i<len(ans);i++{
		ans[i] = delzero(ans[i])
	}
	// 去除第一个字符串，该字符串全零
	ans = ans[1:]
	// 最后，为符合本题要求，转化为整数切片
	res := make([]int,len(ans))
	for i:=0;i<len(ans);i++{
		res[i],_ = strconv.Atoi(ans[i])
	}
	return res
}
// 去除字符串高位 0
func delzero(s string) string {
	start := 0
	for i:=0;i<len(s);i++{
		if s[i]=='0'{
			start++
		} else {
			break
		}
	}
	return s[start:]
}