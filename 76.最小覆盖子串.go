/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (40.75%)
 * Likes:    979
 * Dislikes: 0
 * Total Accepted:    111.5K
 * Total Submissions: 273.5K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
 * 
 * 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 和 t 由英文字母组成
 * 
 * 
 * 
 * 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
 */

// @lc code=start
func minWindow(s string, t string) string {
	// 统计子串各个字符出现频率（需要凑齐的字符）
	need := make(map[byte]int)
	// 统计窗口中凑齐了多少复合要求的子串字符
	windows := make(map[byte]int)
	for k := range t{
		// 遍历统计子串需要凑齐字符数量
		need[t[k]]++
	}
	// left和right是双指针，寻找子串覆盖范围用的
	var valid,left,right,start,length int
	// length 保存最小子串长度，start保存最小子串起始位置
	length = math.MaxInt32
	n := len(s)
	for right<n{
		c := s[right]
		// 扩大右窗口，包含更多的字符
		right++
		if _,ok := need[c];ok{
			// 出现了需要凑齐的统计字符，让此字符数量++
			windows[c]++
			// valid 只要某个字符的出现次数和子串中出现的次数一致则++
			if windows[c] == need[c]{
				valid++
			}
		}
		// 窗口内包含的字符频率完全一致，尝试缩小窗口找最短子串
		for valid == len(need){
			// 每次复合要求，判断是否找到了更短子串
			if right-left < length{
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _,ok := need[d];ok{
				// 一旦windows数量即将不匹配，让valid--,直到下次匹配
				if windows[d] == need[d]{
					valid--
				} 
				windows[d]--
			}
		}
	}
	if length != math.MaxInt32{
		return s[start:start+length]
	}
	return ""
}
// @lc code=end

