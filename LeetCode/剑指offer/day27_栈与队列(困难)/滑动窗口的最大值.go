// 题目链接：https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/?envType=study-plan&id=lcof
// day27/31
// 第 27 天主题为：栈与队列（困难）
// 包含两道题目：
// 剑指offer59-I.滑动窗口的最大值
// 剑指offer59-II.队列的最大值
package main

//可以先去暴力解决，对每个滑动窗口，遍历所有元素得到最大值。
//对长度为 n 的数组 nums，窗口数量为 n-k+1，暴力解法时间复杂度为 O((n-k+1)*k) = O(nk)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0{
		return []int{}
	}
	res := []int{}
	for i := 0;i < n-k+1;i++{
		res = append(res,max(nums[i:i+k]))
	}
	return res
}


func max(sli []int) (res int){
	res = sli[0]
	for _,i := range sli{
		if i > res{
			res = i
		}
	}
	return
}

//法二：单调队列
//然后想一想，暴力没有用到题目中哪些条件，可以针对这些条件去做改进
//相邻的两个窗口，有重叠元素 k-1 个，这 k-1 个元素中的最大值对两个窗口来说是相同的，这就是我们改进的方向
//
//设想一下，一个窗口中两个元素下标分别为 i 和 j，其中 i 在 j 的左侧（i<j)
//- 若 nums[i] < nums[j]，在滑动窗口右移的过程中，只要 nums[i] 在窗口中，则 nums[j] 也一定还在窗口中，这是 i<j 所保证的。
//   因此，由于 nums[j] 的存在，nums[i] 一定不会是滑动窗口最大值，我们可以将 nums[i] 永久删除。
//- 若 nums[i] >= nums[j]，nums[i] 出队后，nums[j] 是有可能成为队列最大值的，nums[j] 需要保留
//
//因此，我们可以使用一个队列来存储还没有被删除的元素。在队列中，这些元素的值是单调递减的
//当窗口向右滑动时，我们需要把一个新的元素放入队列中，为了保持队列的性质，我们需要不断将新元素与队列队尾元素进行比较，
//若新元素较大，则队尾元素出队，不断进行此操作，直至队列为空 或 新的元素小于等于队尾的元素
//
//由于队列的元素是严格单调递减的，且队列中元素属于该窗口，所以队首元素就是该窗口的最大值
//此时，我们需要考虑最后一个问题，若当前窗口的最大值为窗口最左侧元素，那进入下一个窗口前队列中该元素应该出队，因为该元素并不属于下一个窗口
//
//该队列元素单调递减，满足这种单调性的队列一般称作 单调队列。
func maxSlidingWindow_2(nums []int, k int) []int {
	n := len(nums)
	monoQ := []int{}
	res := make([]int,0,n-k+1)
	for i:=0;i<n;i++{
		for len(monoQ) > 0 && nums[i] > monoQ[len(monoQ)-1]{
			monoQ = monoQ[:len(monoQ)-1]
		}
		monoQ = append(monoQ,nums[i])
		// 若单调队列最大值非窗口元素
		// 将该元素出队
		if i >= k && monoQ[0] == nums[i-k]{
			monoQ = monoQ[1:]
		}
		// i=k-1时，窗口元素数量达到 k
		// 开始向 res 数组添加窗口最大值
		if i >= k-1{
			res = append(res,monoQ[0])
		}
	}
	return res
}