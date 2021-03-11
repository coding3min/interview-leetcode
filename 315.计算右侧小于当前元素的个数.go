/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (42.16%)
 * Likes:    528
 * Dislikes: 0
 * Total Accepted:    41K
 * Total Submissions: 97.2K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于
 * nums[i] 的元素的数量。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：nums = [5,2,6,1]
 * 输出：[2,1,1,0] 
 * 解释：
 * 5 的右侧有 2 个更小的元素 (2 和 1)
 * 2 的右侧仅有 1 个更小的元素 (1)
 * 6 的右侧有 1 个更小的元素 (1)
 * 1 的右侧有 0 个更小的元素
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 
 * 
 */

// @lc code=start
func countSmaller(nums []int) []int {

}
// 二分法超时  62/65 cases passed
// func countSmaller(nums []int) []int {
// 	var queue []int
// 	var res []int 
// 	n := len(nums)
// 	for i:=n-1;i>=0;i--{
// 		insertIndex := findIndex(queue,nums[i])
// 		addNum := insertIndex
// 		res = append([]int{addNum},res...)
// 		queue = append(queue[:insertIndex],append([]int{nums[i]},queue[insertIndex:]...)...)
// 	}
// 	return res
// }

// func findIndex(nums []int,num int) int{
// 	i,j := 0,len(nums)-1
// 	if len(nums) == 0{
// 		return 0
// 	}
// 	midIndex := 0
// 	for i<j{
// 		midIndex = (i+j)/2
// 		if nums[midIndex] < num{
// 			i = midIndex + 1
// 		}else{
// 			j = midIndex
// 		}
// 	}
// 	if nums[i] < num{
// 		return i + 1
// 	}
// 	return i
// }
// @lc code=end

