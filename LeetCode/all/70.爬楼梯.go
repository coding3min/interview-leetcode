/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (51.49%)
 * Likes:    1526
 * Dislikes: 0
 * Total Accepted:    384.5K
 * Total Submissions: 745.4K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 注意：给定 n 是一个正整数。
 * 
 * 示例 1：
 * 
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 * 
 * 示例 2：
 * 
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 * 
 * 
 */

// @lc code=start
func climbStairs(n int) int {
	if n<=0{
		return 0
	}
	q,p,res := 0,0,1
	for i:=0;i<n;i++{
		p=q
		q=res
		res = p+q
	}
	return res
}
// 递归超时 21/45 cases passed (N/A)
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2{
// 		return 2
// 	}
// 	if n<=0{
// 		return 0
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }
// @lc code=end

