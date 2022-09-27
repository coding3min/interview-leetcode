// 题目链接：https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/
package main

//最通俗的解法就是排序了
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr,0,len(arr)-1,k)
	return arr[:k]
}

//这么做能满足要求，但没有意义
//
//接下来思考一下，题目要找的是最小的 k 个数，但没有要求这 k 个数字按顺序，很容易想到快排，每一趟快排会确定一个元素的最终位置，
//其左侧元素值均小于该元素，右侧元素值均大于该元素，而左右侧元素是否有序快排并不关心，这与本题的要求不谋而合。
//
//下面用快排的思想来解题。

func getLeastNumbers_2(arr []int, k int) []int {
	quickSort(arr,0,len(arr)-1,k)
	return arr[:k]
}

func quickSort(nums []int,left,right,k int) []int {
	if left < right{
		// partitionIndex为该次分区后，基准元素所在下标
		partitionIndex := partition(nums,left,right)
		//若 k大于该下标，我们只需要对左侧区间递归进行快排，
		// 右侧区间无需处理
		if partitionIndex < k{
			return quickSort(nums,partitionIndex+1,right,k)
		//	k小于该下标，同理
		} else if partitionIndex > k{
			return quickSort(nums,left,partitionIndex-1,k)
		} else {
			// 当相等时，达到我们的要求，返回“排好序”的数组
			return nums
		}
	}
	return nums
}

//移动左右指针，按照基准（这里使用nums[left]）划分区域。最后返回基准所在的下标
func partition(nums []int,left,right int) int {
	// 基准 选择 left 指向的元素
	pivot := nums[left]
	for left < right{
		for left<right && nums[right]>=pivot{
			right --
		}
		nums[left] = nums[right]
		for left<right && nums[left]<=pivot{
			left ++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}