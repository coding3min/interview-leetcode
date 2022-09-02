// 题目链接：https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
// day4/31
// 第 4 天主题为：查找算法（简单）
// 包含三道题目：
// 剑指offer03.数组中重复的数字
// 剑指offer53-I.旋转数组的最小数字
// 剑指offer50.第一个只出现一次的字符

// 解题思路：题目出现有序，优先考虑二分
// 此题的通俗解法为，遍历数组
// 当nums[i]不等于下标i时，说明从下标i开始，之后的元素值均为下标值+1，返回下标i即可。
package main

func missingNumber(nums []int) int {
	for i:=0;i<len(nums);i++{
		if i != nums[i]{
			return i
		}
	}
	return -1
}

// 解法2：二分查找
// 根据题意，可将有序数组 nums 分为两部分，第一部分其值 nums[i] = i，第二部分 nums[i]=i+1
// 并且第二部分元素数量限制均为 大于等于0 且 小于等于 n-1（此n为函数形参n，非数组长度)。
// 先考虑两种特殊情况：
//- 当第一部分长度为0，返回0
//- 第二部分长度为0，返回数组长度即可
// 在常规情况下，我们要找的是第二部分第一个元素的下标
func missingNumber_2(nums []int) int {
	n := len(nums)
	left,right := 0,n-1
	for left <= right{
		mid := left+(right-left)>>1
		//优先考虑两种极端情况
		if mid==0 && nums[mid]!=mid{
			return mid
		} else if mid==n-1 && nums[mid]==mid{
			return mid+1
		} else if nums[mid]!=mid && nums[mid-1]==mid-1{
			return mid
		} else if nums[mid] != mid{
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}