//题目链接：https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/?envType=study-plan&id=lcof
//仔细看下链接URL：diao-zheng，明明应该是 tiao-zheng 这是程序员能干出来的事吗？
package main


//解题思路：双指针
//
//left 和 right 初始化指向 nums 的第一个和最后一个元素下标，循环条件 left < right
//- 左指针从当前位置向右移动寻找第一个偶数
//- 右指针从当前位置向左移动寻找第一个奇数
//以上过程要始终保持 left < right，不符合时直接跳出，然后 left 和 right 指向的元素互换位置。
//left++，right-- 持续以上步骤，直至 left <= right，说明全部奇数已位于偶数前面。
func exchange(nums []int) []int {
	n := len(nums)
	if n <= 1{
		return nums
	}
	left,right := 0,n-1
	for left < right{
		for left<right && nums[left]&1==1{
			left ++
		}
		for left<right && nums[right]&1==0{
			right --
		}
		nums[left],nums[right] = nums[right],nums[left]
		left ++
		right --
	}
	return nums
}

//我曾经犯过一个错，这里描述一下，在 left 向右移 和 right 向左移的过程中，没有写 left<right，只有最外部的 left < right 条件，这会造成什么情况呢？
//对 题目所给实例 [1,2,3,4] 来说，没有问题，left 和 right 分别指向 2 和  3 后，两元素互换位置，left 和 right 值分别为 2 和 1，跳出循环。
//
//但考虑这样一种情况，数组元素全部为奇数，比如：[1,3,5]，那左指针在寻找偶数时，索引会超出数组长度，引发 panic。
//
//再看一种情况，调整完的两个元素，left和right向内移动后，left<right，比如：[1,4,3,2,5]，left 寻找偶数找到 4，
//索引为 1，right寻找奇数位 5，索引为 4，交换位置后，各自向内移动，left=2，right=3，left继续寻找偶数元素，找到2，
//下标为3（此时已经=right了），right向左寻找奇数，找到 3，下标为2，两者交换位置后，left=4,right=1，结束循环
//数组变为：[1,5,2,3,4]，发现多了一次数组元素交换位置。错因就在寻找元素的过程中没有保证left<right。
//
//所以，大家一定要引以为鉴，不要犯类似的错误啦，寻找元素的过程一定要保证left<right