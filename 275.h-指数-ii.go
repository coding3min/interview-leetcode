/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 *
 * https://leetcode-cn.com/problems/h-index-ii/description/
 *
 * algorithms
 * Medium (40.96%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 42.3K
 * Testcase Example:  '[0,1,3,5,6]'
 *
 * 给定一位研究者论文被引用次数的数组（被引用次数是非负整数），数组已经按照 升序排列 。编写一个方法，计算出研究者的 h 指数。
 * 
 * h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）总共有 h
 * 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: citations = [0,1,3,5,6]
 * 输出: 3 
 * 解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 0, 1, 3, 5, 6 次。
 * 由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
 * 
 * 
 * 
 * 说明:
 * 
 * 如果 h 有多有种可能的值 ，h 指数是其中最大的那个。
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 这是 H 指数 的延伸题目，本题中的 citations 数组是保证有序的。
 * 你可以优化你的算法到对数时间复杂度吗？
 * 
 * 
 */

// @lc code=start
 func hIndex(citations []int) int {
	n := len(citations)
	for h:=n;h>0;h--{
		// 先确定期望值，当期望值出现时跳出
		// 不要反着来
		if citations[n-h] >= h{
			return h
		}
	}
	return 0
 }


// func hIndex(citations []int) int {
// 	n := len(citations)
// 	if n==1{
// 		return citations[0]
// 	}
// 	if n==0{
// 		return 0
// 	}
// 	return bSearch(citations,0,n)
// }

// func bSearch(citations []int,x,y int) int{
// 	ans := 0
// 	m := x + (y-x)/2
// 	if x>y{
// 		return 0
// 	}

// 	if citations[m] == m{
// 		return m
// 	}

// 	if citations[m] > m{
// 		ans = bSearch(citations,x,m-1)
// 	}else{
// 		ans = bSearch(citations,m+1,y)
// 	}
// 	return ans
// }
// @lc code=end

