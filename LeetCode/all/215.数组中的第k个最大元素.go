/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.72%)
 * Likes:    945
 * Dislikes: 0
 * Total Accepted:    279K
 * Total Submissions: 431.4K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 
 * 示例 1:
 * 
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 * 
 * 说明: 
 * 
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 * 
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// 随机数种子
    rand.Seed(time.Now().UnixNano())
    return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// 快排剪枝
func quickSelect(a []int, l, r, index int) int {
    q := randomPartition(a, l, r)
    if q == index {
        return a[q]
    } else if q < index {
        return quickSelect(a, q + 1, r, index)
    }
    return quickSelect(a, l, q - 1, index)
}

//随机中枢
func randomPartition(a []int, l, r int) int {
    i := rand.Int() % (r - l + 1) + l
	// 因为把r的位置作为中枢，所以要交换
    a[i], a[r] = a[r], a[i]
    return partition(a, l, r)
}

//快排核心代码
func partition(a []int, l, r int) int {
	// 因为传入的中枢是r位置
    x := a[r]
	// 先把i设置为l-1区间之外，保证每次更新先++再更新
	// i 的作用是记录小于中枢的区间
    i := l - 1
	// j不能遍历到r位置，因为r是中枢本身，查找完之后再做交换
    for j := l; j < r; j++ {
		// 遍历l到r，保证i左侧包括i位置记录的都小于等于x,右侧全部大于x
        if a[j] <= x {
            i++
            a[i], a[j] = a[j], a[i]
        }
    }
	// 最后遍历完交换r和i+1位置，把中枢放到正确的地方
    a[i+1], a[r] = a[r], a[i+1]
    return i + 1
}
// 肯定不能用这种无赖的办法
// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }
// @lc code=end

