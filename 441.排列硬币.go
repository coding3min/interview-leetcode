/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 *
 * https://leetcode-cn.com/problems/arranging-coins/description/
 *
 * algorithms
 * Easy (41.92%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    35.9K
 * Total Submissions: 85.7K
 * Testcase Example:  '5'
 *
 * 你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。
 * 
 * 给定一个数字 n，找出可形成完整阶梯行的总行数。
 * 
 * n 是一个非负整数，并且在32位有符号整型的范围内。
 * 
 * 示例 1:
 * 
 * 
 * n = 5
 * 
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤
 * 
 * 因为第三行不完整，所以返回2.
 * 
 * 
 * 示例 2:
 * 
 * 
 * n = 8
 * 
 * 硬币可排列成以下几行:
 * ¤
 * ¤ ¤
 * ¤ ¤ ¤
 * ¤ ¤
 * 
 * 因为第四行不完整，所以返回3.
 * 
 * 
 */

// @lc code=start
func arrangeCoins(n int) int {
	// 1+2+3+4... 是等差数列，如果=n返回数列长度，n在结果中间取左侧
	// 要求使用二分法
	left,right := 1,n
	for left<=right{
		// 二分法：每次找到中间值，再验证中间值是否正确
		mid := left + (right-left)/2
		sum := mid * (1+mid)/2
		if sum == n{
			return mid
		}
		if sum < n{
			// 验证值和最终值比较，以此来修改区间
			left = mid + 1
		}else{
			right = mid - 1
		}
		fmt.Println(left,right)
	}
	return right
}
// @lc code=end

