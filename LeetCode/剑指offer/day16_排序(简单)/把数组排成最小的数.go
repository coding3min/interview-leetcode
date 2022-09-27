// 题目链接：https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
package main

import (
	"sort"
	"strconv"
)

//这道题我刚开始以为数组的元素都是个位数，想着用哈希表统计每个数字出现的频率，然后从 0 开始遍历，将制
//定数目的数字添加至返回的在字符串中， 看了下示例，才知道数组的元素可以不是个位数。
//
//但其实原理也是类似的，小数在前，大数在后，本质上是一个排序问题，设 nums 中两整数分别为 x 和 y，
//先将其转化为字符串，然后拼接位 x+y 和 y+x 作比较
//- 若拼接字符串 x+y > y+x，则 x > y
//- 否则（包括相等的情况），x < y
//根据此规则，套用任何排序算法对 nums 执行即可。
func minNumber(nums []int) string {
	n := len(nums)
	strNums := make([]string,len(nums))
	for i:=0;i<n;i++{
		strNums[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] < strNums[j]+strNums[i]
	})
	res := ""
	for i:=0;i<n;i++{
		res += strNums[i]
	}
	return res
}


//点过去看了下sort.Slice() 源码，在数组长度>12 的情况下，使用的是快速排序，经过递归，长度降至 12以下时，使用插入排序。
//下面咱来不用标准库，自己复习下冒泡，来解这道题
func minNumber_2(nums []int) string {
	n := len(nums)
	strNums := make([]string,len(nums))
	for i:=0;i<n;i++{
		strNums[i] = strconv.Itoa(nums[i])
	}
	for i:=0;i<n-1;i++{
		for j:=0;j<n-1-i;j++{
			if strNums[j]+strNums[j+1] > strNums[j+1]+strNums[j]{
				strNums[j],strNums[j+1] = strNums[j+1],strNums[j]
			}
		}
	}
	res := ""
	for i:=0;i<n;i++{
		res += strNums[i]
	}
	return res
}
//也是完全没有问题的