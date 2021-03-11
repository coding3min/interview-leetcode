/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 *
 * https://leetcode-cn.com/problems/h-index/description/
 *
 * algorithms
 * Medium (38.92%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    21.6K
 * Total Submissions: 55.4K
 * Testcase Example:  '[3,0,6,1,5]'
 *
 * 给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
 * 
 * h 指数的定义：h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）总共有 h
 * 篇论文分别被引用了至少 h 次。且其余的 N - h 篇论文每篇被引用次数 不超过 h 次。
 * 
 * 例如：某人的 h 指数是 20，这表示他已发表的论文中，每篇被引用了至少 20 次的论文总共有 20 篇。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 输入：citations = [3,0,6,1,5]
 * 输出：3 
 * 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
 * 由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
 * 
 * 
 * 
 * 提示：如果 h 有多种可能的值，h 指数是其中最大的那个。
 * 
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	queryArray := make([]int,n+1)
	for _,v := range citations{
		if v>=n{
			v=n
		}
		queryArray[v]++
	}
	i:=n
	//倒序找到不符合的点
	for nums := queryArray[n];i>=0;{
		// i是期望的h指数
		// queryArray[i]是有多少篇文章引用次数等于i
		// nums 是有多少篇文章大于i
		// i <= nums 代表正好在此处有nums篇文章符合要求
		if i <= nums{
			break
		}
		i--
		nums += queryArray[i]
	}
	return i

	// k := n
	// for s := queryArray[n]; k > s; s += queryArray[k]{
	// 	k--
	// }
    //return k
}

// func hIndex(citations []int) int {
// 	// 求h篇文章被引用h次的最值问题,先排序
// 	sort.Slice(citations,func (i,j int) bool{
// 		return citations[i]>citations[j]
// 	})
// 	ans := 0
// 	// 找到相应位置return就行，要注意一定要记录位置
// 	for k,v := range citations{
// 		if v < k+1{
// 			return ans
// 		}			
// 		ans++
// 	}
// 	return ans
// }
// @lc code=end

