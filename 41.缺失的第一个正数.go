/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (40.49%)
 * Likes:    945
 * Dislikes: 0
 * Total Accepted:    109.4K
 * Total Submissions: 270K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
 * 
 * 
 * 
 * 进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,0]
 * 输出：3
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,4,-1,1]
 * 输出：2
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [7,8,9,11,12]
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * -2^31 
 * 
 * 
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
		// 首先既然是连续正整数，那必然是符合数组长度，且不为0，那就是在[1-300]之间
		// 当前下标存储的数就应该是在当前数-1的位置，因为数组是从0开始的，最后存储 [1,2,3] 这样
		// 所以比较应该是 num[i] 也就是当前值，和num[i]-1作为下标的正确位置，不相等就是没放对，就交换
		// 至于为什么是循环，是因为交换结束后nums[i]的值更新了，要再找一次正确位置
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    return n + 1
}

// 能想到用超大的数组下标当标记位，或者用字典存储完了以后遍历一遍
// @lc code=end

