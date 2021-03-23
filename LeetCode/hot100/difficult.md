### 介绍

一句话总结算法思路，LeetCode hot100 difficult

按频率排序，排序依据参考 [字节跳动后端高频面试题目](https://github.com/afatcoder/LeetcodeTop/blob/master/bytedance/backend.md)。

全部源码可见我的GitHub [interview-leetcode](https://github.com/minibear2333/interview-leetcode/tree/master/LeetCode/all)

注：

有下划线标志的都是超链接。
点击下列题目标题可以跳转到LeetCode中文官网直接阅读题目，提交代码。
点击下列代码链接，可以直接跳转到我的GitHub代码页面。每道题一般精选一种解法，我的GitHub中可能收录多种解法代码，请自行查看。

### 题解

### 42.接雨水

题目：[接雨水，给定整数数组，把他想象成柱状图，凹槽部分接雨水总合](https://leetcode-cn.com/problems/trapping-rain-water/)

题解：用栈的思路比较难理解，我写在代码里了，有兴趣可以自己看看，这里推荐双指针的思路
* 双指针下标`l`和`r`指向`0`和`len-1`，记录左侧最大值`height[0]`和右侧最大值`height[len-1]`
* 1. 判断左侧和右侧最大值哪个更小，更小侧向内测移动，比如 `if lMax<rMax then l++`
* 2. 如果当前指针位置比`lMax`要小，又因为此时`rMax>lMax`说明此处一定会出现低洼，它被两侧包住了，一定不会渗水；水的高度是由更小的最大值决定，现在是`lMax`
* 3. 统计水量`res += lMax - height[l]`，如果当前指针比`lMax`大，说明左侧包不住，更新左侧最大值
* 上面1、2、3步骤的动作`else` 反过来，`r--; if height[r]<rMax then res+= rMax-height[r] else rMax = height[r]`
* 循环条件`l<r`,函数至少要3个值才有可能形成低洼，所以边界是`len<3 return 0` 


代码：[golang](../all/42.接雨水.go)


### 

题目：[]()

题解：

注意：

代码：[golang](..)

### 最后

如果文中有误，欢迎提pr或者issue，**一旦合并或采纳作为贡献奖励可以联系我直接无门槛**加入[技术交流群](https://mp.weixin.qq.com/s/ErQFjJbIsMVGjIRWbQCD1Q)

我是小熊，关注我，知道更多不知道的技术

![](res/2021-03-17-19-57-33.png)
