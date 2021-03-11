/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 *
 * https://leetcode-cn.com/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (52.47%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    23.6K
 * Total Submissions: 44.9K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * 给一非空的单词列表，返回前 k 个出现次数最多的单词。
 * 
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
 * 
 * 示例 1：
 * 
 * 
 * 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 * ⁠   注意，按字母顺序 "i" 在 "love" 之前。
 * 
 * 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
 * k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 * ⁠   出现次数依次为 4, 3, 2 和 1 次。
 * 
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
 * 输入的单词均由小写字母组成。
 * 
 * 
 * 
 * 
 * 扩展练习：
 * 
 * 
 * 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
 * 
 * 
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	queryMap := make(map[string]int)
	var wordArray []string
	for _,v := range words{
		if _,ok := queryMap[v];!ok{
			wordArray = append(wordArray,v)
		}
		queryMap[v]++
	}
	sort.Slice(wordArray,func (i,j int)(bool){
		a,b := wordArray[i],wordArray[j]
		countA,countB := queryMap[a],queryMap[b]
		if countA == countB{
			return strings.Compare(a,b) < 0
		}
		return countA > countB
	})
	return wordArray[:k]
}

// 比较完善的排序方式
// type WordStruct struct{
// 	Word string
// 	Cnt int
// }
// type WordStructs []WordStruct
// func (w WordStructs) Len() int{return len(w)}
// func (w WordStructs) Swap(i,j int) {w[i],w[j]=w[j],w[i]}
// func (w WordStructs) Less(i,j int) bool{
// 	a,b := w[i],w[j]
// 	if a.Cnt == b.Cnt{
// 		return strings.Compare(a.Word,b.Word) < 0
// 	}
// 	return a.Cnt>b.Cnt
// }

// func topKFrequent(words []string, k int) []string {
// 	queryMap := make(map[string]int)
// 	for _,v := range words{
// 		queryMap[v]++
// 	}
// 	var WordStructArray WordStructs
// 	for word,cnt := range queryMap{
// 		w := WordStruct{
// 			Word:word,
// 			Cnt:cnt,
// 		}
// 		WordStructArray = append(WordStructArray,w)
// 	}
// 	sort.Sort(WordStructArray)
// 	var ans []string
// 	for i,word := range WordStructArray{
// 		if i<k{
// 			ans = append(ans,word.Word)
// 		}
// 	}
// 	return ans
// }
// @lc code=end

