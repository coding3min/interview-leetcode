// 题目链接：https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/?plan=lcof&plan_progress=klczit3
// day3/31
// 第 3 天主题为：字符串（简单）
// 包含两道题目：
// 剑指offer05.替换空格
// 剑指offer58-II.左旋转字符串


// 此题题解前三个方法参考自：https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/solution/mian-shi-ti-58-ii-zuo-xuan-zhuan-zi-fu-chuan-qie-p/
// 解题思路：设字符串s的长度为n，首先，很容易想到，s左旋n位还是原本的字符串，我们只需要左旋k%n次即可。
// 此题解法较多，请耐心观看
package main

// 方法1：也是最通俗的解法，字符串切片进行拼接
func reverseLeftWords(s string, n int) string {
	n = n % len(s)
	return s[n:] + s[:n]
}

// 如果面试不允许对字符串进行切片，那我们可以对列表进行遍历，然后进行拼接，记为方法2：
// 先遍历从下标n到末尾的字符串元素，加入返回的rune切片res，然后遍历下标0-n，添加至res。最后返回将res转换为string返回即可
// 为避免命名冲突，此函数名添加后缀 “_2"
func reverseLeftWords_2(s string, n int) string {
	res := make([]rune,len(s))
	n = n % len(s)
	k := 0
	for i:=n;i<len(s);i++{
		res[k] = rune(s[i])
		k ++
	}
	for i:=0;i<n;i++{
		res[k] = rune(s[i])
		k++
	}
	return string(res)
}

// 方法三：还是遍历字符串，一次遍历完成
// 主要是通过取余操作对方法2的第二次遍历进行优化，感觉算是一种写法技巧，并不能提升算法性能。
// 为避免命名冲突，此函数名添加后缀 “_3"
func reverseLeftWords_3(s string, n int) string {
	res := make([]rune,len(s))
	n = n % len(s)
	k := 0
	for i:=n;i<len(s)+n;i++{
		res[k] = rune(s[i%len(s)])
		k ++
	}
	return string(res)
}

// 方法4：两次翻转字符串
// 前三种解法本质上是完全一样的，都是拆解字符串然后进行拼接。
// 本解法来自《剑指offer》，做本题前建议先做 剑指offer58题：翻转单词顺序
// 以“abcdefg”为例，对该字符串进行两次翻转，仍为原字符串，如果想左移2位，我们先将其分为两部分，“ab”为第一部分，“cdefg”为第二部分，
// 先不管每一部分内部的字符，对整个字符串进行翻转得到 “gfedcba”，第一部分到了右侧，第二部分到了左侧，得到了大体上的相对顺序，
// 由于进行了一次翻转，每部分字符串内部字符顺序是颠倒的，题目要求的每一部分相对顺序不变，那我们再对两部分字符串，分别进行一次翻转即可。
// 为避免命名冲突，此函数名添加后缀 “_4"
func reverseLeftWords_4(s string, n int) string {
	n = n % len(s)
	x := reverseString(s)
	res := reverseString(x[:len(s)-n]) + reverseString(x[len(s)-n:])
	return res
}

func reverseString(s string) string{
	res := []rune(s)
	left,right := 0,len(s)-1
	for left < right{
		res[left],res[right] = res[right],res[left]
		left ++
		right --
	}
	return string(res)
}

// 很久以前就看到这个解法了，但我一直有个疑惑，为什么两次翻转对应的是左旋，而不是右旋呢？翻转和左右又没有关系，后面手写试了几次，发现了规律，第一次整体翻转后
// 若要左移，原字符串左边的 k 位到了右边 k 位，我们要做的是翻转 s[:n-k] 和 s[n-k:]
// 若要右移，原字符串右边的 k 位到了左边 k 位，我们要做的是翻转 s[:k] 和 s[k:]

// 如果要右移字符串，对应如下代码：
// 右移字符串
func reverseRightWords(s string, n int) string {
	n = n % len(s)
	x := reverseString(s)
	res := reverseString(x[:n]) + reverseString(x[n:])
	return res
}
// 这下就搞清楚了！另外，左移1位，也等价于右移 len(s)-1位。
