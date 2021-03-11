/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (41.81%)
 * Likes:    313
 * Dislikes: 0
 * Total Accepted:    75.6K
 * Total Submissions: 180.7K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 * 
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 * 
 * 
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	// 同76题，唯一区别是找到子串且子串必须连续就退出
	need := make(map[byte]int)
	windows := make(map[byte]int)
	for k := range s1{
		need[s1[k]]++
	}
	var valid,left,right int
	n := len(s2)
	for right<n{
		c := s2[right]
		right++
		if _,ok := need[c];ok{
			windows[c]++
			if windows[c] == need[c]{
				valid++
			}
		}
		// 只在增加查找长度时判断，是否找到以及是否符合标准
		if right-left==len(s1) && valid == len(need){
			return true
		}
		// 注意这里的长度只要大于等于子串都需要缩减
		for right-left>=len(s1){
			c := s2[left]
			left++
			if _,ok := need[c];ok{
				if windows[c] == need[c]{
					valid--
				}
				windows[c]--
			}
		}
	}
	return false
}
// @lc code=end

