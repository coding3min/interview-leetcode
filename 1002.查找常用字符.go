/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 *
 * https://leetcode-cn.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (74.30%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    48.8K
 * Total Submissions: 65.8K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * 给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3
 * 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
 * 
 * 你可以按任意顺序返回答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：["bella","label","roller"]
 * 输出：["e","l","l"]
 * 
 * 
 * 示例 2：
 * 
 * 输入：["cool","lock","cook"]
 * 输出：["c","o"]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= A.length <= 100
 * 1 <= A[i].length <= 100
 * A[i][j] 是小写字母
 * 
 * 
 */

// @lc code=start
func commonChars(A []string) []string {
	queryMap := make(map[rune]int)
	for _,str := range A{
		freqMap := make(map[rune]int)
		for _,c := range str{
			freqMap[c]++
		}
		if len(queryMap)==0{
			queryMap = freqMap
			continue
		}
		for c,cnt := range queryMap{
			if freqMap[c]<cnt{
				queryMap[c] = freqMap[c]
			}
		}
	}
	var res []string
	for c,cnt := range queryMap{
		for i:=0;i<cnt;i++{
			res = append(res,string(c))
		}
	}
	return res
}
// @lc code=end

