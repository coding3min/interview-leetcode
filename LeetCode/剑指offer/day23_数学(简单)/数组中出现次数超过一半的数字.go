//题目链接：https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
// day23/31
// 第 23 天主题为：数学（简单）
// 包含两道题目：
// 剑指offer39.数组中出现次数超过一半的数字
// 剑指offer66.构建乘积数组

package main
//本题通俗解法为哈希表统计数组中每个元素出现次数，然后遍历哈希表键值对，找到值大于数组长度一般的键，返回即可。时间空间复杂度均为 O(n)
//
//方法二：排序，将 nums 元素从小到大排序，那下标为 n/2 的元素为要寻找的元素,时间复杂度O(nlogn)，空间 O(1)
//
//考虑一种更好的解法
//方法三：摩尔投票算法（参考自官方题解）

//算法步骤如下：
//1. 维护一个候选主要元素 candidate 和候选主要元素出现的次数 count，初始时 candidate 为任意值，count 为 0；
//2. 遍历数组 nums 中的元素，设元素值为 x，依次对元素执行如下操作：
//	如果 count = 0，则将 x 赋给 candidate，否则不更新 candidate 的值；
//	如果 x = candidate，count 自增 1，否则将 count 减 1 。
//3. 遍历结束后，如果 nums 中存在主要元素，则 candidate 为主要元素，否则 candidate 可 能为数组中的任意一个元素。

//由于数组不一定存在主要元素，因此需要第二次遍历数组，验证 candidate 是否为主要元素。这次遍历统计 candidate 在数组中出现的次数，
//若大于数组长度的一半，则 candidate 是主要元素，否则，不是。
//
//为什么当数组中存在主要元素时，Boyer-Moore 投票算法可以确保得到主要元素？
//在 Boyer-Moore 投票算法中，遇到相同的数则将 count 加 1，遇到不同的数则将 count 减 1。
//根据主要元素的定义，主要元素的出现次数大于其他元素的出现次数之和，因此在遍历过程中，主要元素和其他元素两两抵消，
//最后一定剩下至少一个主要元素，此时 candidate 为主要元素，且 count≥1。
func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _,num := range nums{
		if count == 0{
			candidate = num
			count = 1
			continue
		}
		if num == candidate{
			count ++
		} else {
			count --
		}
	}
	return candidate
}