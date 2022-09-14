// 题目链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
// day16/31
// 第 16 天主题为：排序（简单）
// 包含两道题目：
// 剑指offer45.把数组排成最小的数
// 剑指offer61.扑克牌中的顺子
package main

import (
	"sort"
)


//解题思路：排序

//具体细节如下：
//给定数组长度等于 5，数据量很小，我们先对其排序
//然后一次遍历，检查是否存在重复元素（0除外），若存在，一定不连续
//此时，数组中除 0 外，不存在重复元素，得到其最大值和最小值（0除外）,若最大值与最小值之差小于4，说明连续，否则，不连续。
func isStraight(nums []int) bool {
	sort.Ints(nums)
	for i:=1;i<5;i++{
		if nums[i] == 0 {
			continue
		}
		if nums[i]==nums[i-1] {
			return false
		}
	}
	max := nums[len(nums)-1]
	min := nums[0]
	for i:=0;nums[i]==0;i++{
		min = nums[i+1]
	}
	if max - min > 4{
		return false
	}
	return true
}


//另外，我们也可以不用排序，而是使用一次遍历的方式求得最大最小值（0除外）
//思路和排序是一致的，代码如下
//这种解法注意 minNum 和 maxNum 的初始化，minNum 和 maxNum 分别初始化为最大和最小的数字
func isStraight_2(nums []int) bool {
	minNum,maxNum := 14,0
	record := map[int]struct{}{}
	for _,num := range nums{
		if num == 0{
			continue
		}
		if _,ok:=record[num];ok{
			return false
		}
		record[num] = struct{}{}
		minNum = min(minNum,num)
		maxNum = max(maxNum,num)
	}
	if maxNum - minNum > 4{
		return false
	}
	return true
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

func min(x,y int) int{
	if x < y{
		return x
	}
	return y
}