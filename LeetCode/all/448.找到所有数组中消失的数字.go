/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (63.68%)
 * Likes:    650
 * Dislikes: 0
 * Total Accepted:    94.6K
 * Total Submissions: 148.6K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 * 
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 * 
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 * 
 * 示例:
 * 
 * 
 * 输入:
 * [4,3,2,7,8,2,3,1]
 * 
 * 输出:
 * [5,6]
 * 
 * 
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for _,v := range nums{
		x := nums[abs(v)-1]
		nums[abs(v)-1] = -abs(x)
	}
	var res []int
	for k,v := range nums{
		if v>0{
			res = append(res,k+1)
		}
	}
	return res
}

func abs(x int) int{
	if x<0{
		x = -x
	}
	return x
}

// hash表
// func findDisappearedNumbers(nums []int) []int {
// 	queryMap := make([]int,len(nums))
// 	for _,v := range nums{
// 		queryMap[v-1] = 1
// 	}
// 	var res []int
// 	for k,v := range queryMap{
// 		if v==0{
// 			res = append(res,k+1)
// 		}
// 	}
// 	return res
// }
// @lc code=end

