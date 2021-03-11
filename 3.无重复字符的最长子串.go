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
    sArray := strings.Split(s,"")
    // map用来存储当前字符上一次出现的下标
    countMap := make(map[string]int)
    startIndex,maxLen := -1,0
    for k,v := range sArray{
        // 如果当前字符在之前出现过，还要保证是在当前子串内部，而不是之前出现的才算是重复
        if _,ok := countMap[v];ok && countMap[v] > startIndex  {
            // 如果是出现重复
            startIndex = countMap[v]
            countMap[v] = k
        }else{
            countMap[v] = k
            if k - startIndex > maxLen{
                maxLen = k - startIndex
            }
        }
    }
    return maxLen
}

// func lengthOfLongestSubstring(s string) int {
//     sArray := strings.Split(s,"")
//     startIndex,maxLen := 0,0
//     for k,_ := range sArray{
//         len := 1
//         for j := startIndex; j<k;j++{
//             if sArray[j]!=sArray[k]{
//                 len++
//             }else{
//                 startIndex = j+1
//                 break
//             }
//         } 
//         if len>maxLen{
//             maxLen = len
//         }
//     }
//     return maxLen
// }
// @lc code=end

