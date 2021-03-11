/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (34.94%)
 * Likes:    2522
 * Dislikes: 0
 * Total Accepted:    579.7K
 * Total Submissions: 1.7M
 * Testcase Example:  '123'
 *
 * 给你一个 32 位的有符号整数 x ，返回 x 中每位上的数字反转后的结果。
 * 
 * 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
 * 假设环境不允许存储 64 位整数（有符号或无符号）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：x = 123
 * 输出：321
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：x = -123
 * 输出：-321
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：x = 120
 * 输出：21
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：x = 0
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * -2^31 
 * 
 * 
 */

// @lc code=start
func reverse(x int) int {
	num := 0
	for x!=0{
		tmp := x % 10
		x = x/10
		num = num * 10 + tmp
	}
	INT32_MAX := int(^uint32(0)>>1)
	if num < int(^INT32_MAX) || num > INT32_MAX{
		return 0
	}
	return num
}
// @lc code=end

