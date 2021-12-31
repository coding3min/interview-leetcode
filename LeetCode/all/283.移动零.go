/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (64.04%)
 * Likes:    1361
 * Dislikes: 0
 * Total Accepted:    578.2K
 * Total Submissions: 902.9K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 * 
 * 说明:
 * 
 * 
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 * 
 * 
 */

// @lc code=start
func moveZeroes(nums []int)  {
	if len(nums) == 0{
		return
	}
	i := 0
	for j := 0 ; j<len(nums);j++{
		if nums[j] == 0{
			continue
		}
		nums[i] = nums[j]
		i++
	}
	for ;i<len(nums);i++{
		nums[i] = 0
	}
}
// @lc code=end

