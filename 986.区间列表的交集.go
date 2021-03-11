/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] 区间列表的交集
 *
 * https://leetcode-cn.com/problems/interval-list-intersections/description/
 *
 * algorithms
 * Medium (66.25%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    13.4K
 * Total Submissions: 20.2K
 * Testcase Example:  '[[0,2],[5,10],[13,23],[24,25]]\n[[1,5],[8,12],[15,24],[25,26]]'
 *
 * 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而
 * secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。
 * 
 * 返回这 两个区间列表的交集 。
 * 
 * 形式上，闭区间 [a, b]（其中 a ）表示实数 x 的集合，而 a  。
 * 
 * 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList =
 * [[1,5],[8,12],[15,24],[25,26]]
 * 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：firstList = [[1,3],[5,9]], secondList = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：firstList = [], secondList = [[4,8],[10,12]]
 * 输出：[]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：firstList = [[1,7]], secondList = [[3,10]]
 * 输出：[[3,7]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * firstList.length + secondList.length >= 1
 * 0 i < endi 
 * endi < starti+1
 * 0 j < endj 
 * endj < startj+1
 * 
 * 
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	n1,n2 := len(firstList),len(secondList)
	var res [][]int
	if n1 ==0 || n2==0{
		return res
	}
	for i,j := 0,0; i < n1 && j < n2;{
		var lo,hi int
		// 求两个区间的交集
		lo = max(firstList[i][0],secondList[j][0])
		hi = min(firstList[i][1],secondList[j][1])
		if lo<=hi{
			// 找到交集
			res = append(res,[]int{lo,hi})
		}
		// 短区间 向后跑
		if firstList[i][1]<secondList[j][1]{
			i++
		}else{
			j++
		}
	}
	return res
}
func min(x,y int) int{
	if x<y{
		return x
	}
	return y
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}
/**
 * 扩展题目，先遍历更正区间，再前左闭区间为主排序，然后可以复用以上方法，奇偶下标区间当作上面的两个区间
 * 1.计算多个区间的交集
 *   区间用长度为2的数字数组表示，如[2, 5]表示区间2到5（包括2和5）；
 *   区间不限定方向，如[5, 2]等同于[2, 5]；
 *   实现`getIntersection 函数`
 *   可接收多个区间，并返回所有区间的交集（用区间表示），如空集用null表示
 * 示例：
 *   getIntersection([5, 2], [4, 9], [3, 6]); // [4, 5]
 *   getIntersection([1, 7], [8, 9]); // null
 */
// @lc code=end

