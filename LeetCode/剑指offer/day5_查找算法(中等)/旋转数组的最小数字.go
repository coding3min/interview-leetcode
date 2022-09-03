//题目链接：https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/?plan=lcof&plan_progress=klczit3

package main

// 尽管题目描述看起来并不复杂，但本题是一道实实在在的困难题
// 我个人在这道题上花费时间超过5h，现在有时候还是会绕进去
//

// 二分思想很简单，细节是魔鬼！各种边界处理问题
// Although the basic idea of binary search is comparatively straightforward, the details can be surprisingly tricky...            ----Knuth
//
// 解题思路：二分查找，参照题解 [旋转数组的最小数字 - 旋转数组的最小数字 - 力扣（LeetCode）](https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-by-leetcode-s/)
// 强烈建议观看官方题解，有图片作参考会好理解一些
// 左右边界 left 和 right 初始化为 0 和 len(numbers)-1，假设最小值出现在 mid 位置，我们考虑数组中的最后一个元素 x：
// 在最小值右侧的元素，它们的值一定都小于等于 x；而在最小值左侧的元素，它们的值一定都大于等于 x。因此，我们可以根据这一条性质，通过二分查找的方法找出最小值。

// 将中间元素 number[mid] 和 numbers[right] 作比较，可能会出现以下三种情况：（为什么拿这两个位置的数字作比较，题解没有说，但这绝对不是简简单单就能想到的，一定是做了推理，或者经验所得）
//1. numbers[mid] < numbers[right]，说明 numbers[mid] 是最小值右侧的元素（也可能就是numbers[mid])，因此我们忽略当前整个搜索区间的右半部分。right=mid
//2. number[mid] > numbers[right]，说明最小值在 mid 右侧，我们忽略二分查找区间的左半部分。left = mid + 1
//3. numbers[mid] = numbers[right]，在这种情况下，我们无法确定最小值在mid的左侧还是右侧，但我们能确定的是，由于它们值相同，所以无论 numbers[high]是不是最小值，
//   其都有一个替代值 numbers[mid]，因此我们可以忽略当前搜索区间的右端点。right--
// 最终 left 与 right 落于同一位置，即最小值所在的索引。
//
// 在这里，谈一下个人的想法：看题解，有时候我们觉得看懂了，但关掉题解，写代码的时候就忘了该怎么做，或者隔几天就忘了，很多时候就是因为你只看到了题解部分，
// 没有看到题解之外的，作者是如何从题目所给条件推导到题解这一步的，就比如，我第一次看这道题的官方题解时，就没有思考过为什么查找的条件是 left < right，而不是 left <= right，
// 还有，为什么比较的是 mid 元素 和 right 元素，而不是 left 元素，这些题解都没有说，如果自己不去想清楚，那就做不到对这道题的完美把控。
// 当然，现在我明白了，left=right 时，我们已经找到了最小值所在的索引。
// 这种能力的获得不是一蹴而就的，而是慢慢获取的。我也在练习的过程中...



func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high - low) / 2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

