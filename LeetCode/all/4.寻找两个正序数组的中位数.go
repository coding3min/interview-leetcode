/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (39.56%)
 * Likes:    3679
 * Dislikes: 0
 * Total Accepted:    330.7K
 * Total Submissions: 835.8K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
 * 
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 * 
 * 
 * 示例 4：
 * 
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 * 
 * 
 * 示例 5：
 * 
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 * 
 * 
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i,j,n1,n2 := 0,0,len(nums1),len(nums2)
	if n1 == 0{
		fmt.Println("n1 is empte")
		return getMid(nums2)
	}
	if n2 == 0{
		fmt.Println("n2 is empte")
		return getMid(nums1)
	}
	sumLen := n1 + n2
	k := sumLen / 2
	singleInt := sumLen % 2
	for i+j < sumLen{
		if k > n1{
			i = n1 - 1
		}
		if k > n2{
			j = n2 - 1
		}	
		if nums1[i] < nums2[j]{
			nums1 = nums1[i+1:]
			k = k - (n1 - len(nums1))
			n1 = len(nums1)
		}else if nums1[i] > nums2[j]{
			nums2 = nums2[j+1:]
			k = k - (n2 - len(nums2))
			n2 = len(nums2)
		}else{
			if (n1 == 1 && n2 == 1) || k == 1{
				return getAns(singleInt,nums1[0],nums2[0])
			}
			nums1 = nums1[i+1:]
			nums2 = nums2[j+1:]
			k = k - (n1 - len(nums1) - len(nums2))
			n1 = len(nums1)
			n2 = len(nums2)
		}
		fmt.Println(i,j,k,n1,n2)
		if (n1 == 1 && n2 == 1) || k == 1{
			return getAns(singleInt,nums1[0],nums2[0])
		}
		if n1 == 0{
			return getMid(nums2)
		}
		if n2 == 0{
			return getMid(nums1)
		}

	}
	return 0
}

func getMid(nums []int) float64{
	n := len(nums)
	if n==0{
		return 0
	}

	if n % 2 == 1{
		return float64(nums[n/2])
	}else{
		return float64(nums[n/2] + nums[n/2-1]) / 2
	}
	return 0
}

func getAns(singleInt, x,y int) float64{
	if singleInt == 1{
		if x<y{
			return float64(x)
		}
		return float64(y)
	}else{
		return float64(x + y) / 2 
	}
}
// @lc code=end

