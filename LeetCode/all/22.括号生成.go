/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.93%)
 * Likes:    1640
 * Dislikes: 0
 * Total Accepted:    243.4K
 * Total Submissions: 316.3K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：["()"]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0{
		return []string{}
	}
	var res []string
	var dfs func(l,r int,curr string)
	//为了使用res，需要内置函数
	dfs = func(l,r int,curr string){
		// 因为要配对，左括号必须大于右括号，剪支减去错误的情况
		if l>r{
			return
		}
		if l == 0 && r == 0{
			res = append(res,curr)
			return
		}
		if l>0{
			// 注意必须传入l-1，而不能传入 --l 这里不能改变当前l的值
			dfs(l-1,r,curr+"(")
		}
		if r>0{
			dfs(l,r-1,curr+")")
		}
	}
	dfs(n,n,"")
	return res
}
// @lc code=end

