/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (42.66%)
 * Likes:    1182
 * Dislikes: 0
 * Total Accepted:    121.3K
 * Total Submissions: 284K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 保证入栈的是顺序的，一个比一个大
	var stack []int
	max := 0
	for k,v := range heights{
		for len(stack)>0 && heights[stack[len(stack)-1]] < v{
			stack = stack[]
		} 

		stack = append(stack,v)
	}

	return max
}


// 方法1 超时 89/96 cases passed (N/A) 存在很多重复比较
// func largestRectangleArea(heights []int) int {
// 	// 这一题莫名的像接雨水
// 	MIN_INT := ^int(^uint(0)>>1)
// 	max := MIN_INT
// 	n := len(heights)
// 	for k,v := range heights{
// 		tmp := v
// 		// 保证左右都比当前大，就是他的面积
// 		stop := false
// 		for i,j := k-1,k+1; !stop && (i>=0 || j < n);{
// 			stop = true
// 			if i>=0 && heights[i]>= v{
// 				tmp += v
// 				i--
// 				stop = false
// 			}
// 			if j< n && heights[j]>=v{
// 				tmp += v
// 				j++
// 				stop = false
// 			}
// 		}
// 		if tmp > max{
// 			max = tmp
// 		}
// 	}
// 	return max
// }
// @lc code=end

