//题目链接：https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/?envType=study-plan&id=lcof
package main

func reversePairs(nums []int) int {
	n := len(nums)
	res := 0
	for i:=0;i<n-1;i++{
		for j:=i+1;j<n;j++{
			if nums[i] > nums[j]{
				res ++
			}
		}
	}
	return res
}

//提交后发现超时，需要想其他的方法
//参考自LeetCode官方题解
//
//相比第一题，我认为这道题更体现出分治的思想。
//解题的前提是理解归并排序，归并排序的核心在于合并两个有序序列时，我们只需要 O(n) 的时间复杂度，当序列长度为 1 时，自然有序，然后向上回溯。

//「归并排序」是分治思想的典型应用，它包含这样三个步骤：
//- 分解： 待排序的区间为 [l, r]，令 m=(l+r)/2 ，我们把 [l, r] 分成[l,m] 和[m+1,r]
//- 解决： 使用归并排序递归地排序两个子序列
//- 合并： 把两个已经排好序的子序列 [l,m] 和 [m+1,r] 合并起来
//
//在待排序序列长度为 1 的时候，递归开始「回升」，因为我们默认长度为 1 的序列是排好序的。
//那么求逆序对和归并排序又有什么关系呢？关键就在于「归并」当中「并」的过程。我们通过一个实例来看看。
//假设我们有两个已排序的序列等待合并，分别是 L= { 8,12,16,22,100 } 和 R = { 9, 26, 55, 64, 91}。
//一开始我们用指针 lPtr = 0 指向 L 的首部，rPtr = 0 指向 R 的头部。记已经合并好的部分为 M。
//
//L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = []
//我们发现 lPtr 指向的元素 8 小于 rPtr 指向的元素 9，于是把 lPtr 指向的元素放入 M，并把 lPtr 后移一位。
//这个时候我们把左边的 8 加入了 M，我们发现右边没有数比 8 小，所以 8 对逆序对总数的「贡献」为 0。

//接着我们继续合并，把 9 加入了答案，此时 lPtr 指向 12，rPtr 指向 26。
//此时 lPtr 比 rPtr 小，把 lPtr 对应的数加入答案，并考虑它对逆序对总数的贡献为 rPtr 相对 R 首位置的偏移 1（即右边只有一个数比 12 小，
//所以只有它和 12 构成逆序对），以此类推。
//
//我们发现用这种「算贡献」的思想在合并的过程中计算逆序对的数量的时候，只在 lPtr 右移的时候计算，是基于这样的事实：
//当前 lPtr 指向的数字比 rPtr 小，但是比 R 中 [0 ... rPtr - 1] 的其他数字大，所以这里就贡献了 rPtr 个逆序对。
//以上就是官方题解的内容。



//我想再补充一下自己的理解：
//设数组 nums 长度 为 n，左半序列 left=nums[:n/2]，右半序列 right=nums[n/2:]，两者都是有序的，
//那在当前数组中，右半序列的所有元素贡献度均为 0，因为其右方数组均大于等于自身
//在左半序列中，每个元素相对于左半序列来说贡献度为 0，但我们要求的是针对数组 nums 每个元素的贡献度，
//左半序列每个元素大于右半序列元素的个数就是其贡献度，就可以合并两个序列，每次左指针移动的时候，得到左半序列某个元素的贡献度。
//很巧妙的分治思想

func reversePairs_2(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end - start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid + 1, end)
	tmp := []int{}
	i, j := start, mid + 1
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i - start]
	}
	return cnt
}

//然后，我想着，能够在在常见的归并排序写法中，通过加入一个整型指针，归并时修改该指针的值，来达到统计逆序对的目的呢
//尝试后，发现是可以的，占位符 _ 代表排序后的数组，因为用不到，所以使用了占位符。
//需要特别注意注意一种情况，就是左序列的指针指向的元素值和右序列指向的元素值相等时，如何处理，在常规的归并排序中，
//在这种情况下，无论我们先处理左指针指向的元素，还是右指针指向的元素，都是可以的。
//但在这里不行，我们必须优先移动左指针，因为相同的情况下，左指针指向的元素已经和右指针指向的元素无法形成逆序对。这点需要特别注意
func reversePairs_3(nums []int) int {
	res := 0
	_ = merge_sort(nums,&res)
	return res
}

func merge_sort(nums []int,cnt *int) []int {
	n := len(nums)
	if n < 2{
		return nums
	}
	mid := n/2
	left_part := merge_sort(nums[:mid],cnt)
	right_part := merge_sort(nums[mid:],cnt)
	result := merge(left_part,right_part,cnt)
	return result
}

func merge(left_part,right_part []int,cnt *int) []int {
	result := make([]int,0)
	i,j := 0,0
	for i<len(left_part) && j<len(right_part){
		// 重点在这里 需要=
		// 不然 [1,3,2,3,1]过不了
		if left_part[i] <= right_part[j]{
			result = append(result,left_part[i])
			*cnt = *cnt + j
			i++
		} else {
			result = append(result,right_part[j])
			j++
		}
	}
	for k:=0;k<len(left_part)-i;k++{
		*cnt += len(right_part)
	}
	result = append(result,left_part[i:]...)
	result = append(result,right_part[j:]...)
	return result
}

