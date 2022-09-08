//题目链接：https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/?envType=study-plan&id=lcof

package main

//我们可以用一个字典存储每个字符最近一次出现的下标
//一次遍历字符串，变量 start 代表当前不含重复字符的子字符串的起始下标-1，初始化为 -1，
//为什么是 -1呢，因为我们遍历第一个字符的时候，其下标为 0，为不包含重复字符的子字符串，长度为0-（-1），返回变量res初始化为 0
//
//一次遍历字符串，for i,x := range s ，维护 s[start+1;i+1] 为不含重复字符的子串
//若当前遍历到的字符的最近一次出现下标大于 start（不会出现等于的情况），说明当前加入当前字符后，该字符串将包含重复字符，
//我们的做法是，将 start 更新到其最近一次出现的下标，这样就保证了加入当前字符后，我们目前的子字符串仍然不包含重复字符，
//
//否则（小于 start 或者 该字符在字符串中是第一次出现），当前不含重复字符的子串长度增加，更新 res 变量，

//最后，更新当前字符的最近出现下标。
func lengthOfLongestSubstring(s string) int {
	start := -1
	record := map[rune]int{}
	res := 0
	for i,x := range s{
		if _,ok := record[x];ok && record[x] > start{
			start = record[x]
		} else {
			res = max(res,i-start)
		}
		record[x] = i
	}
	return res
}

func max(x,y int) int {
	if x > y{
		return x
	}
	return y
}