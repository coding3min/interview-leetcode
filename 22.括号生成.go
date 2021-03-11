/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0{
		return []string{}
	}
	var res []string
	var dfs func(l,r int,curStr string)
	dfs = func(l,r int,curStr string){
		// 因为要配对，左括号必须大于右括号，剪支减去错误的情况
		if l>r{
			return
		}
		if l==0 && r == 0{
			res = append(res, curStr)
			return
		}
		if l>0{
			// 这里不能改变当前l的值
			dfs(l-1,r,curStr+"(")
		}
		if r>0{
			dfs(l,r-1,curStr+")")
		}
	} 
	dfs(n,n,"")
	return res
}
// @lc code=end

