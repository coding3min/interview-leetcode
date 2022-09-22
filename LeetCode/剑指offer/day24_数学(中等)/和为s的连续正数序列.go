//题目链接：https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/?envType=study-plan&id=lcof

package main

//解题思路：滑动窗口
//
//设连续正整数序列的左边界与右边界分别为 left 和 right，构建滑动窗口向右滑动。
//循环中，每轮判断滑动窗口内元素和与 target 的大小关系
//- 若相等则记录结果，然后移动左指针
//- 若 大于 target 则移动左指针
//- 若小于 target 则移动右指针
//
//每次移动指针的同时，更新滑动窗口内元素和 sum
//因为序列至少由两个数字组成，所以左指针边界为 [1,target/2]
func findContinuousSequence(target int) [][]int {
	// left,right 为滑动窗口左右指针
	// sum 动态记录窗口元素和
	// 窗口至少含有两个数
	left, right, sum := 1, 2,3
	res := make([][]int, 0)
	// 序列至少右两个元素组成，所以左边界只需遍历到 target/2
	for left <= target>>1 {
		if sum < target {
			// move right cursor and increase sum
			right++
			sum += right
		} else if sum > target {
			// move left cursor and reduce sum
			sum -= left
			left++
		} else {
			nums := make([]int, 0)
			for i := left; i <= right; i++ {
				nums = append(nums, i)
			}
			res = append(res,nums)
			sum -= left
			left++
		}
	}
	return res
}