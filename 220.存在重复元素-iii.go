/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (26.56%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    29.6K
 * Total Submissions: 111.4K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j
 * 的差的绝对值也小于等于 ķ 。
 * 
 * 如果存在则返回 true，不存在返回 false。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [1,2,3,1], k = 3, t = 0
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入: nums = [1,0,1,1], k = 1, t = 2
 * 输出: true
 * 
 * 示例 3:
 * 
 * 输入: nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出: false
 * 
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums)
	if n <= 1{
		return false
	}
	queryMap := make(map[int]int)
	w := t + 1
	for i:=0;i<n;i++{
		if i>k{
			// 维护一个窗口，窗口超过map就删掉
			// -1是为了保证窗口内有k个元素，超过的去掉
			delete(queryMap,getID(nums[i-k-1],w))
		}
		// 通过生成器生成窗口id
		id := getID(nums[i],w)
		if _,ok := queryMap[id];ok{
			return true
		}
		if _,ok := queryMap[id-1];ok{
			if nums[i] - queryMap[id-1]  <= t{
				return true
			}
		}
		if _,ok := queryMap[id+1];ok{
			if queryMap[id+1] - nums[i] <= t{
				return true
			}
		}
		queryMap[id] = nums[i]
	}
	return false
}

func getID(num,w int) int{
	if (num >= 0) {
        return num / w;
    } else {
		// /w得到桶id，但是 num<w时 num/w=0，此时会与正数/w=0的桶冲突，所以再-1，整体平移一位
        return num / w - 1;
    }
}
// @lc code=end

