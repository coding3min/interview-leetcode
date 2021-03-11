/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (49.28%)
 * Likes:    755
 * Dislikes: 0
 * Total Accepted:    266.1K
 * Total Submissions: 539.5K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 * 
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自
 * nums2 的元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * 输出：[1,2,2,3,5,6]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums1 = [1], m = 1, nums2 = [], n = 0
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * nums1.length == m + n
 * nums2.length == n
 * 0 
 * 1 
 * -10^9 
 * 
 * 
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n==0{
		return 
	}
	if m == 0{
		copyArray(nums1,nums2,0,0)
		return
	}
	curr := 0
	temp := make([]int,m+n)
	for i,j,k := 0,0,0;;{
		if i<m && j <n{
			if nums1[i] < nums2[j]{
				curr = nums1[i]
				i++
			}else{
				curr = nums2[j]
				j++
			}
		}else{
			if i >= m{
				copyArray(temp,nums2,k,j)
				break
			}else{
				copyArray(temp,nums1,k,i)
				break
			}
		}
		temp[k] = curr
		k++
	}

	copyArray(nums1,temp,0,0)
}

func copyArray(dest,src []int,i,j int){
	for i<len(dest){
		dest[i] = src[j]
		i++
		j++
	}
}
// @lc code=end

