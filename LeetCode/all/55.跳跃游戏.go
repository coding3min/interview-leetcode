/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (41.65%)
 * Likes:    1037
 * Dislikes: 0
 * Total Accepted:    192.7K
 * Total Submissions: 462.5K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个下标。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
func canJump(nums []int) bool {
	// 剩余能量
	resEnergy := 0
	// 最后一格是终点，所以只用计算到倒数第二格
	for i:=0;i<len(nums)-1;i++{
		// 每次达到当前格时，计算剩余能量和当前能量哪个大，每次取最大即可
		// 因为每次我们都是自己选择下一步跳多长的
		if resEnergy < nums[i]{
			resEnergy = nums[i]
		}
		if resEnergy==0{
			return false
		}
		// 假如上次能量没用完，相当于判断上次能量和他跨过的所有能量块进行比较
		resEnergy--
	}
	// len < 2时一定是true
	return true 
}
// @lc code=end

