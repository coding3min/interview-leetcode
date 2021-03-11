/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Medium (45.56%)
 * Likes:    872
 * Dislikes: 0
 * Total Accepted:    225.4K
 * Total Submissions: 494.7K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 * 
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: nums = [1,2,3,4,5,6,7], k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入：nums = [-1,-100,3,99], k = 2
 * 输出：[3,99,-1,-100]
 * 解释: 
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -2^31 
 * 0 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
func rotate(nums []int, k int)  {
	// 取余数，存在k超过len的情况
	splitIndex := k%len(nums)
	revert(nums)
	revert(nums[:splitIndex])
	revert(nums[splitIndex:])
	
}

func revert(nums []int){
	i,j := 0,len(nums)-1
	for i<j{
		nums[i],nums[j]= nums[j],nums[i]
		i++
		j--
	}
}
// 暴力解法1
// func rotate(nums []int, k int)  {
// 	splitIndex := k%len(nums)
// 	newNums := []int{}
// 	for i:= splitIndex;i < len(nums);i++{
// 		newNums = append(newNums,nums[i])
// 	}
// 	for i:=0;i<splitIndex;i++{
// 		newNums = append(newNums,nums[i])
// 	}

// 	for i:=0;i<len(newNums);i++{
// 		nums[i] = newNums[i]
// 	}
// }

// 暴力2
// func rotate(nums []int, k int)  {
// 	splitIndex := k%len(nums)
// 	newNums := []int{}
// 	newNums = append(nums[splitIndex:],nums[:splitIndex]...)

// 	for i:=0;i<len(newNums);i++{
// 		nums[i] = newNums[i]
// 	}
// }

// 全部翻转，然后再翻转2次即可


// @lc code=end

