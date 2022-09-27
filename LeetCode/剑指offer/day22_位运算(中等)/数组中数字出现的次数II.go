//题目链接：https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/?envType=study-plan&id=lcof
package main

//本题最通俗的解法是哈希表记录每个数字出现的次数，我们还是用位运算进行优化，将空间复杂度从 O(n) 降低至 O(1)
//
//解题思路：依次确定二进制位
//记数组中出现一次的数字为 res，题目给定数组中整数范围为 1~2^31-1，我们可以从最低位开始，依次确定 res 当前位 是 0 还是 1 。
//在每一位中，我们遍历数组所有元素，统计当前位为 1 的个数 num，若 num 对 3 取余 为 0，说明 res 当前位为 0，否则为 1 。

func singleNumber(nums []int) int {
	n := len(nums)
	res := 0
	for i:=0;i<32;i++{
		num := 0
		bit := 1<<i
		for j:=0;j<n;j++{
			if nums[j] & bit != 0{
				num ++
			}
		}
		if num % 3 == 1{
			res |= bit
		}
	}
	return res
}