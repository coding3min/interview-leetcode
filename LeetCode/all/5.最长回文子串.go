/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (33.18%)
 * Likes:    3325
 * Dislikes: 0
 * Total Accepted:    504.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "cbbd"
 * 输出："bb"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "a"
 * 输出："a"
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "ac"
 * 输出："a"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 仅由数字和英文字母（大写和/或小写）组成
 * 
 * 
 */

// @lc code=start
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1{
		return s
	}
	// begin 和 max_length 为当前最长回文子串的开始位置和长度
	max_len := 1
	begin := 0
	dp := make([][]bool,n)
	// 长度为 1 的字符串均为回文串
	for i := 0;i < n;i++{
		dp[i] = make([]bool,n)
		dp[i][i] = true
	}
	// 枚举子串长度，从 2 开始，最长为 n
	for L := 2;L < n + 1;L++{
		// 枚举左边界
		for left := 0;left < n;left++{
			// 固定字符串长度下的右边界
			right := L + left - 1
			// 右边界越界，则退出此长度下的循环
			if right >= n{
				break
			}
			// 如果左右边界字符不同，则一定不是回文串
			if s[left] != s[right]{
				dp[left][right] = false
			} else {
				// right - left <= 2 等价于 字符串长度 <= 3
				// 此时只需要字符串左右两端字符相等，即为回文串
				if right - left <= 2{
					dp[left][right] = true
				} else {
					// 该情况下，字符串左右边界符合回文串条件，
					// 改字符串是否为回文串要根据左右各缩进一位的字符串是否为回文串来进行判断
					dp[left][right] = dp[left + 1][right - 1]
				}
			}
			// 更新最长回文子串
			if dp[left][right] && right - left + 1 > max_len{
				max_len = right - left + 1
				begin = left
			}
		}
	}
	return s[begin:begin + max_len]
}
// @lc code=end

