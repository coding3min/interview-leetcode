/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (43.00%)
 * Likes:    1855
 * Dislikes: 0
 * Total Accepted:    409.2K
 * Total Submissions: 951.7K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 * 
 * 有效字符串需满足：
 * 
 * 
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 * 
 * 
 * 注意空字符串可被认为是有效字符串。
 * 
 * 示例 1:
 * 
 * 输入: "()"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "()[]{}"
 * 输出: true
 * 
 * 
 * 示例 3:
 * 
 * 输入: "(]"
 * 输出: false
 * 
 * 
 * 示例 4:
 * 
 * 输入: "([)]"
 * 输出: false
 * 
 * 
 * 示例 5:
 * 
 * 输入: "{[]}"
 * 输出: true
 * 
 */

// @lc code=start
func isValid(s string) bool {
	var stack []byte
	queryMap := map[byte]byte{'{':'}','[':']','(':')'}
	for i := range s{
		if _,ok := queryMap[s[i]];ok{
			stack = append(stack,s[i])
		}else{
			n := len(stack)
			if n>0 && queryMap[stack[len(stack)-1]] == s[i]{
				stack = stack[:n-1]
			}else{
				return false
			}
		}
	}
	return len(stack) == 0
}
// @lc code=end

