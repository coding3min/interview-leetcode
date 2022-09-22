//题目链接：https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/?envType=study-plan&id=lcof

package main

//约瑟夫环问题，模拟整个删除过程最直观，这里我用数组进行模拟，首先构建长度为 n 的模拟数组 nums，
//元素值分别为 0 ~ n-1，start 为每次循环中数组第一个元素下标，初始化为 0，之后开始模拟
//
//先获取当前数组的长度 clen，删除第 m 个数字，m 可能大于等于 n，我们将 m-1 对 n 取余， 直接得到要删除的元素下标，
//考虑到初始元素下标不一定为 0，最终待删除的元素下表 loc_del = (m-1+start) % clen，
//将删除 该元素后的数组重新赋值给 nums，start 更新为 loc_del，开启下一次循环，直至数组长度为 1，
//可得到圆圈剩下的最后一个数组，返回该元素即可。

func lastRemaining(n int, m int) int {
	// 每次循环中数组的第一个元素下标
	start := 0
	// 构建 nums 数组
	nums := []int{}
	for i:= 0;i <n;i++{
		nums = append(nums,i)
	}
	// 数组长度等于1时，循环结束
	for len(nums) > 1{
		clen := len(nums)
		loc_del := (m-1+start) % clen
		nums = append(nums[:loc_del],nums[loc_del+1:]...)
		start = loc_del
	}
	return nums[0]
}

//提交后会超时，个人感觉时间是浪费在数组的拼接上。
//
//下述解法参考自官方题解
//现在我们建模 n 个数字删除 第 m 个元素的问题为 f(n,m)，f(n,m) 的值为最后剩余元素值（同元素下标）
//那 f(1,m) 的答案是一定的，数组只存在一个元素，f(1,m) = 0，我们可以试想下，可不可以通过 f(1,n) 推出 f(n,m) 的值。
//进一步来说，如果能通过 f(n-1,m) 推出 f(n,m)，那就一定能从 f(1,n) 推出 f(n,m)
//因为 f(n,m) 要删除第 m%n 个元素后，长度就变成了 n-1，那自然而然就与 f(n-1,m) 扯上了关联，它们的区别在哪里呢?
//对应元素下标不同！ 只要我们把它们的下标相对应起来，那我们就可以求解改题目。
//下标不同在哪里呢？起始位置，f(n-1,m) 起始元素下标为 0，而 f(n,m) 删除一个元素后的起始元素下标为 m%n，
//我们把它们对应起来，就可以通过 f(n-1,m) 求得 f(n,m)
//同理，由于 f(1,m) 值是固定的，我们可以从 f(1,m) 递归到 f(n,m)

func lastRemaining_2(n int, m int) int {
	// n=1时，剩余元素下标为0
	res := 0
	// 从前向后推，求n=2,3,...,n时，剩余元素下标
	// i 为数组长度
	for i:= 2;i <= n;i++{
		// 将i-1数组元素下标与长度为i的数组下标对应起来
		// 求得长度为 i 的数组剩余元素
		res = (res + m) % i
	}
	return res
}