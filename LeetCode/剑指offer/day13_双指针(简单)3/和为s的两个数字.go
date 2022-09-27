// 题目链接：https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/?envType=study-plan&id=lcof
// day13/31
// 第 13 天主题为：双指针（简单）
// 包含三道题目：
// 剑指offer21.调整数组顺序使奇数位于偶数前面
// 剑指offer57.和为s的两个数字
// 剑指offer58-I.翻转单词顺序

package main

//梦回LeetCode第一题：两数之和，有点像是吧，这道题和两数之和的区别在于：数组有序。
//方法 1 ：先来暴力解法，双层遍历
//第一层遍历从头开始，第二层遍历从第一层遍历元素的下一个元素开始，若两数之和等于 target，返回即可。
func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if nums[i]+nums[j] == target{
				return []int{nums[i],nums[j]}
			}
		}
	}
	return []int{0,0}
}
//嗯...和预想的一样，超时了，时间复杂度O(n^2)，空间复杂度 O(1)

//方法 2：哈希表
//接下来用 两数之和 的解法，一次遍历，哈希表记录遍历过的数字
func twoSum_2(nums []int, target int) []int {
	record := map[int]struct{}{}
	for i:=0;i<len(nums);i++{
		if _,ok := record[target-nums[i]];ok{
			return []int{nums[i],target-nums[i]}
		}
		record[nums[i]] = struct{}{}
	}
	return []int{0,0}
}
//时间复杂度O(n)，空间复杂度O(n)


//方法三：双指针
//上面这两种解题思路没有用到数组有序的这个非常重要的条件，一定会有更优解！
//我们是如何使用双指针来进行优化的呢？这是对双层遍历的优化，在双层遍历中，第一层遍历，元素从左至右，第二层遍历，元素从第一个元素的下一个元素向右遍历。
//考虑使用数组有序的这个条件，双指针指向最小和最大的元素，若两数之和大于 target，最大元素移向第二大的元素，若小于 target，最小元素移向第二小的元素，
//这样两边同时向 target 靠拢，一次遍历就可以得到结果。

//使用双指针，left 和 right 指针初始化指向数组的第一个和最后一个元素的下标
//
//- 若两指针指向元素之和 大于 target，说明需要减小两数之和，右指针向左移
//- 若两指针指向元素之和 小于 target，说明需要增大两数之和，左指针右移
//- 若相等，将两指针指向元素组合为切片返回即可。
func twoSum_3(nums []int, target int) []int {
	left,right := 0,len(nums)-1
	for left < right{
		if nums[left]+nums[right] > target{
			right --
		} else if nums[left]+nums[right] < target{
			left ++
		} else {
			return []int{nums[left],nums[right]}
		}
	}
	return []int{0,0}
}
//时间复杂度O(n)，空间复杂度O(1)，充分利用题目所给条件，为最优解法！