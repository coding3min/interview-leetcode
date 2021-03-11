/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.71%)
 * Likes:    965
 * Dislikes: 0
 * Total Accepted:    316.6K
 * Total Submissions: 497.1K
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
	n := len(nums)
	if n == 0{
		return
	}
	var cnt int
	i := 0
	for j :=0;j<n;j++{
		if nums[j]!=0{
			nums[i] = nums[j]
			i++
		}else{
			cnt++
		}
	}
	for ;i<n;i++{
		nums[i] = 0
	}
	return 
}
// @lc code=end

