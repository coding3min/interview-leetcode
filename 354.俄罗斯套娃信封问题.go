/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (39.15%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    25.4K
 * Total Submissions: 64.8K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 * 
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 * 
 * 说明:
 * 不允许旋转信封。
 * 
 * 示例:
 * 
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3 
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 * 
 * 
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n==0{
		return 0
	}
	
	sort.Slice(envelopes,func(i,j int) bool{
		if envelopes[i][0]<envelopes[j][0]{
			return true
		}
		if envelopes[i][0] == envelopes[j][0]{
			if envelopes[i][1]>envelopes[j][1]{
				return true
			}
		}
		return false
	})
	dp := make([]int,n)
	for i:=0;i<n;i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if envelopes[i][1] > envelopes[j][1]{
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
	}
	resMax := math.MinInt32
	for _,v := range dp{
		resMax = max(resMax,v)
	}
	return resMax
}
func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
// @lc code=end

