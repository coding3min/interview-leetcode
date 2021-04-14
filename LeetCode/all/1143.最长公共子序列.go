/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 *
 * https://leetcode-cn.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (61.31%)
 * Likes:    520
 * Dislikes: 0
 * Total Accepted:    110.3K
 * Total Submissions: 177.1K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * 给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
 * 
 * 一个字符串的 子序列
 * 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
 * 
 * 
 * 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
 * 
 * 
 * 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：text1 = "abcde", text2 = "ace" 
 * 输出：3  
 * 解释：最长公共子序列是 "ace" ，它的长度为 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：text1 = "abc", text2 = "abc"
 * 输出：3
 * 解释：最长公共子序列是 "abc" ，它的长度为 3 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：text1 = "abc", text2 = "def"
 * 输出：0
 * 解释：两个字符串没有公共子序列，返回 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * text1 和 text2 仅由小写英文字符组成。
 * 
 * 
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	n1,n2 := len(text1),len(text2)
	dp := make([][]int,n1+1)
	dp[0] = make([]int,n2+1)
	for i:=0;i<n1;i++{
		dp[i+1] = make([]int,n2+1)
		for j:=0;j<n2;j++{
			if text1[i] == text2[j]{
				// 两个长度都-1时的最长公共子序列长度+1
				dp[i+1][j+1] = dp[i][j]+1 
			}else{
				//各自-1时的最长公共子序列对比
				dp[i+1][j+1] = max(dp[i][j+1],dp[i+1][j])
			}
		}
	}
	return dp[n1][n2]
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}
// @lc code=end

