// 题目链接：https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

package main


// 解题思路：两次遍历字符串，第一次遍历的过程中，用哈希表存储每个字符出现的次数，
// 第二次遍历字符串时，哈希表查看当前字符出现次数，若哈希表值为1，返回该字符。
func firstUniqChar(s string) byte {
	if len(s) == 0{
		return ' '
	}
	record := map[rune]int{}
	for _,x := range s{
		record[x] ++
	}
	for _,x := range s{
		if record[x] == 1{
			return byte(x)
		}
	}
	return ' '
}