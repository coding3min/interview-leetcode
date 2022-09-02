//题目链接：https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

package main

// 最通俗的解法，遍历数组，统计target出现的次数
// 但这样的话完全没用到题目中数组是排序的 这个条件
// 显然，本题想要考察二分查找算法！
// 看到有序数组就应该第一时间想到二分查找
func search(nums []int, target int) int {
	n := len(nums)
	ans := 0
	for i := 0;i < n;i++{
		if nums[i] == target{
			ans += 1
		}
	}
	return ans
}

// 方法2：两次二分查找
// 统计一个数字在排序数组中出现的次数，那我们只需要知道其在数组中第一次和最后一次出现的下标
// 设为 left 和 right，出现次数即为：right-left+1
// 函数 first_equal_search 和 last_equal_search 分别用于查找数字第一次和最后一次出现的位置下标。
// 实现上述两个函数时，网上有很多花里花哨的写法，如果去死记硬背，没过几天就会全部忘光
// 下面我的写法非常好理解
// 为避免命名冲突，次函数名添加后缀 “_2”
func search_2(nums []int, target int) int {
	left := first_equal_search(nums,target)
	right := last_equal_search(nums,target)
	// 若出现次数为 0，需额外处理
	if left == -1{
		return 0
	}
	return right-left+1
}

func first_equal_search(nums []int,target int) int{
	n := len(nums)
	left,right := 0,n-1
	for left <= right{
		mid := left + (right-left)>>1
		if mid==0 && nums[mid]==target{
			return mid
		} else if mid>0 && nums[mid]==target && nums[mid-1]!=target{
			return mid
		} else if nums[mid] < target{
			left = mid + 1
		} else if nums[mid] > target{
			right = mid - 1
		} else {
			// 此种情况说明 mid 指向元素为众多与target相等元素的其中一个
			// 而且不在起始点，我们要找第一个，所以移动右指针
			right = mid -1
		}
	}
	return -1
}

func last_equal_search(nums []int,target int) int{
	n := len(nums)
	left,right := 0,n-1
	for left <= right{
		mid := left + (right-left)>>1
		if mid==n-1 && nums[mid]==target{
			return mid
		} else if mid<n-1 && nums[mid]==target && nums[mid+1]!=target{
			return mid
		} else if nums[mid] < target{
			left = mid + 1
		} else if nums[mid] > target{
			right = mid - 1
		} else {
			// 与找第一个相等的元素同理，不在最后一个，所以移动左指针
			left = mid + 1
		}
	}
	return -1
}