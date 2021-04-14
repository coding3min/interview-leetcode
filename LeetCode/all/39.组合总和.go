/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.15%)
 * Likes:    1290
 * Dislikes: 0
 * Total Accepted:    239.8K
 * Total Submissions: 331.7K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的数字可以无限制重复被选取。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1：
 * 
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * 示例 2：
 * 
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 * 
 * 
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	n := len(candidates)
	if n == 0{
		return res
	}
	var backtrace func (curr []int,sum,startIndex int)
	backtrace = func(curr []int,sum,startIndex int){
		if sum > target{
			return
		}
		if sum == target{
			var tmp []int
			tmp = append(tmp,curr...)
			res = append(res,tmp)
			return
		}
		for i:=startIndex;i<n;i++{
			curr = append(curr,candidates[i])
			backtrace(curr,sum+candidates[i],i)
			curr = curr[:len(curr)-1]
		}
	}
	backtrace([]int{},0,0)
	return res
}
// @lc code=end

