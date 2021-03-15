/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (30.78%)
 * Likes:    2951
 * Dislikes: 0
 * Total Accepted:    416.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * -10^5 
 * 
 * 
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n<3{
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int,0,n/3)
	for i,l,r := 0,1,n;i<n;i++{
		// 因为排了序，nums[i]>0时后面的数是正的，不可能再加起来为0
		if nums[i] > 0{
			break
		}
		// 去重,已经找过的就不必找了
		if i>0 && nums[i] == nums[i-1]{
			continue
		}
		// 对当前元素的后面的所有元素进行处理的原因是，前面的元素已经找到了所有符合条件的可能，不需要再找
		l = i+1
		r = n-1
		for l<r{
			// 这里使用target是尽可能防止溢出，是在代码中需要考虑的点
			// 不理解的话，用下面的例子也没有问题
			target := -nums[i]
			sumLR := nums[l] + nums[r]
			if sumLR == target{
				tmpArr := []int{nums[i],nums[l],nums[r]}
				res = append(res,tmpArr)
				l++
				r--
				//跳过所有重复的项
				for l<r && nums[l] == nums[l-1]{
					l++
				}
				for l<r && nums[r] == nums[r+1]{
					r--
				}
			// 和sum比较时不符合跳出
			}else if sumLR<target{
				l++
			}else if sumLR>target{
				r--
			}
		}
	}
	return res
}

// 此方法和上面的一样
// func threeSum(nums []int) [][]int {
// 	n := len(nums)
// 	if n<3{
// 		return [][]int{}
// 	}
// 	sort.Ints(nums)
// 	res := make([][]int,0,n/3)
// 	for i,l,r := 0,1,n;i<n;i++{
// 		// 因为排了序，nums[i]>0时后面的数是正的，不可能再加起来为0
// 		if nums[i] > 0{
// 			break
// 		}
// 		// 去重,已经找过的就不必找了
// 		if i>0 && nums[i] == nums[i-1]{
// 			continue
// 		}
// 		// 对当前元素的后面的所有元素进行处理的原因是，前面的元素已经找到了所有符合条件的可能，不需要再找
// 		l = i+1
// 		r = n-1
// 		for l<r{
// 			sum := nums[i] + nums[l] + nums[r]
// 			if sum == 0{
// 				tmpArr := []int{nums[i],nums[l],nums[r]}
// 				res = append(res,tmpArr)
// 				l++
// 				r--
// 				//跳过所有重复的项
// 				for l<r && nums[l] == nums[l-1]{
// 					l++
// 				}
// 				for l<r && nums[r] == nums[r+1]{
// 					r--
// 				}
// 			// 和sum比较时不符合跳出
// 			}else if sum<0{
// 				l++
// 			}else if sum>0{
// 				r--
// 			}
// 		}
// 	}
// 	return res
// }
// @lc code=end

