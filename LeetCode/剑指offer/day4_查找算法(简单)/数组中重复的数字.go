// 题目链接：https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

package main

// 解题思路：本题最通俗的解法为哈希表，用哈希表记录nums数组中元素是否出现过，
// 若出现过，返回该元素即可。
func findRepeatNumber(nums []int) int {
	record := map[int]int{}
	for _,num := range nums{
		if record[num] > 0{
			return num
		}
		record[num] ++
	}
	return -1
}

// 方法2：设计哈希表
// 哈希表解法的时间与空间复杂度均为O(n)，该方法没有用到题目给的数字范围条件，所以应该思考下如何做改进
// 试想下如何降低复杂度，对长度为n的数组，其中所有数字都在0~n-1范围内，而0~n-1又是长度为n的数组的所有下标索引。我们刚才方法是将其值放入哈希表
// 而现在我们给定数组的长度为n，0~n-1刚好可以对应数组索引，让我们有一种将数组设计为哈希表的思路：
// 对数组进行遍历，设当前遍历到的的数字值为 x，则将索引为 x 对应的数字打标记，若之后遍历到某元素后，打标记过程中发现该元素已打标记，说明该数字已出现过，返回即可。
// 现在要做的就是设计标记，标记为取负值，流程如下：
// 一次遍历数组中元素，设当前元素值为num，它可能已经被打了标记，我们取用其原数x=abs(num)
// 如果 nums[x] < 0，说明 x 已经出现过，返回 x 即可，否则，打标记，nums[x] = -nums[x]
// 但还存在一个问题，就是0取负还是0，我没想到太好的解决方案，就用了最朴素的方法，单独处理0。
// 处理流程如下：初始化0的下标为 zero_index，遍历数组，得到0的下标zero_index，然后第二次遍历数组，若数组元素值为 zero_index的元素个数大于1，说明该元素重复，返回0的下标即可。
// 这样的解题思路来自LeetCode41题：缺失的第一个正数
// 算法核心在于将输入数组设计为哈希表，另外，我在此题的题解部分还没有看人有人用类似的方案解此题
// 在进行此操作前，请务必与面试官进行交流，询问是否可以修改输入数组，确定可以的话，再用此方案。
// 为避免命名冲突，次函数名添加后缀 “_2"
func findRepeatNumber_2(nums []int) int {
	// 先处理0的情况
	// m为nums中0的下标出现的次数，zero_index为0的下标初始值
	m := 0
	zero_index := -1
	// 第一次遍历得到0的下标
	for i:=0;i<len(nums);i++{
		if nums[i] == 0{
			zero_index = i
			break
		}
	}
	// 第二次遍历判断0的下标是否出现两次及两次以上
	for i:=0;i<len(nums);i++{
		if nums[i]==zero_index{
			m ++
		}
	}
	if m > 1{
		return zero_index
	}
	// 将输入数组设计为哈希表
	for _,num := range nums{
		x := abs(num)
		// nums[x]<0，说明x元素此前出现过
		if nums[x] < 0{
			return x
		}
		// 若未出现过，打标记
		nums[x] = -nums[x]
	}
	return -1
}

func abs(num int) int {
	if num < 0{
		return -num
	}
	return num
}


// 方法3：原地交换 次题解方法来自：https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/mian-shi-ti-03-shu-zu-zhong-zhong-fu-de-shu-zi-yua/
// 长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内 。 此说明含义：数组元素的 索引 和 值 是 一对多 的关系。
// 因此，可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应（即 nums[i] = i）。因而，就能通过索引映射对应的值，起到与字典等价的作用。
// 遍历中，第一次遇到数字 x 时，将其交换至索引 x 处；而当第二次遇到数字 x 时，一定有 nums[x] = x，此时即可得到一组重复数字。。
// 算法流程：
// 遍历数组 nums，设索引初始值为 i = 0 :
// 若 nums[i] = i： 说明此数字已在对应索引位置，无需交换，因此跳过；
// 若 nums[nums[i]] = nums[i]： 代表索引nums[i]处和索引i处的元素值都 nums[i]，即找到一组重复值，返回此值 nums[i]；
// 否则，交换索引为i和nums[i]的元素值，将此数字交换至对应索引位置。
// 若遍历完毕尚未返回，则返回 -1 。
// 为避免命名冲突，次函数名添加后缀 “_3"
func findRepeatNumber_3(nums []int) int {
	i := 0
	for i < len(nums){
		// 交换 nums[i]至索引i出后才进行下一索引处的交换
		if nums[i] == i{
			i ++
			continue
		}
		if nums[nums[i]] == nums[i]{
			return nums[i]
		}
		nums[nums[i]],nums[i] = nums[i],nums[nums[i]]
	}
	return -1
}