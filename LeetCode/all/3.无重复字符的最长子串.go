/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (36.32%)
 * Likes:    4933
 * Dislikes: 0
 * Total Accepted:    822.3K
 * Total Submissions: 2.3M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: s = "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 * 示例 4:
 * 
 * 
 * 输入: s = ""
 * 输出: 0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * s 由英文字母、数字、符号和空格组成
 * 
 * 
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0{
        return 0
    }
    // map用来存储当前字符上一次出现的下标
   queryMap := make(map[byte]int)
   // startIndex是字符串起始位置，maxLen是最大不重复子串长度
   startIndex,maxLen := 0,1
   for k := range s{
       v := s[k]
        // 如果当前字符在之前出现过，还要保证是在当前子串内部，而不是之前出现的才算是重复
       if _,ok := queryMap[v];ok && queryMap[v]>= startIndex{
            // 字符串起始位置应该是上一次重复字符的后一位
            startIndex = queryMap[v]+1
       }else{
           // 起始位置到当前字符的长度需要加1
           if k-startIndex+1 > maxLen{
               maxLen = k - startIndex+1
           }
       }
       queryMap[v] = k
   }
   return maxLen
}
// @lc code=end

