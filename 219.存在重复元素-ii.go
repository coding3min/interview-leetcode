/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (41.09%)
 * Likes:    243
 * Dislikes: 0
 * Total Accepted:    78.6K
 * Total Submissions: 191.2K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的
 * 绝对值 至多为 k。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 * 
 * 示例 3:
 * 
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 * 
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	queryMap := make(map[int]int)
	for i,v := range nums{
		if _,ok := queryMap[v];!ok{
			queryMap[v] = i 
		}else{
			lastIndex := queryMap[v]
			if i > lastIndex + k{
				queryMap[v] = i
			}else{
				return true
			}
		}
	}
	return false
}
// @lc code=end

