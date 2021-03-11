/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (65.42%)
 * Likes:    652
 * Dislikes: 0
 * Total Accepted:    162.9K
 * Total Submissions: 248.9K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 * 
 * 示例:
 * 
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * 说明：
 * 
 * 
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 * 
 * 
 */

// @lc code=start
// 原来这一题用hash也更快
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	queryMap := make(map[string][]string,len(strs))
	for _,v := range strs{
		queryKey := encode(v)
		queryMap[queryKey] =  append(queryMap[queryKey], v)
	}
	 for _,v:= range queryMap {
        res = append(res, v)
    }
	return res
}
func encode(s string) string {
	// 关键是byte还可以转成hash
	encodeSlice := make([]byte, 26)
    for _, char := range s {
        encodeSlice[char - 'a']++
    }

    return string(encodeSlice)
}

// 正常方法
// func groupAnagrams(strs []string) [][]string {
// 	n := len(strs)
// 	anagramsFlags := make([]bool,n)
// 	var res [][]string
// 	for i:=0; i<n;{
// 		var tmp []string
// 		tmp = append(tmp,strs[i])
// 		for j:=i+1;j<n;j++{
// 			if isAnagram(strs[i],strs[j]){
// 				tmp = append(tmp,strs[j])
// 				anagramsFlags[j] = true
// 			}
// 		}
// 		res = append(res,tmp)
// 		i++
// 		for i<n && anagramsFlags[i]{
// 			i++
// 		}
// 	}
// 	return res
// }
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t){
// 		return false
// 	}
// 	var x,y [26]byte
// 	for k := range s{
// 		x[s[k]-'a']++
// 		y[t[k]-'a']++
// 	}
// 	return  x == y
// }
// @lc code=end

