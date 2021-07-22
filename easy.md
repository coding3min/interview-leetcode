# LeetCode-hot100-easy

## 介绍

一句话总结算法思路，LeetCode hot100 easy

总计21题，题目列表请戳 [LeetCode Hot100 easy 列表](https://leetcode-cn.com/problemset/leetcode-hot-100/?difficulty=%E7%AE%80%E5%8D%95)。

全部源码可见我的GitHub [interview-leetcode](https://github.com/minibear2333/interview-leetcode/tree/master/LeetCode/all)

注：

有下划线标志的都是超链接。 点击下列题目标题可以跳转到LeetCode中文官网直接阅读题目，提交代码。 点击下列代码链接，可以直接跳转到我的GitHub代码页面。每道题一般精选一种解法，我的GitHub中可能收录多种解法代码，请自行查看。

## 题解

### 1.两数之和\(高频\)

题目：[数据中是否有两个数之和为目标值](https://leetcode-cn.com/problems/two-sum)

题解：遍历数组，用目标值减去当前值，判断HashMap是否有值存在，如果有则创建新数组返回两者，如果没有循环遍历完返回空数组

时间复杂度：O\(1\) 代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/1.两数之和.go)

### 20.有效的括号\(高频\)

题目：[存在`[]{})(`的字符串，判断是否合法](https://leetcode-cn.com/problems/valid-parentheses)

题解： 存储左括号和右括号的映射，用栈统计左括号，出现左括号就入栈，出现右括号就和栈顶在map中映射的右括号比较，如果匹配就出栈，不匹配返回false，最后遍历完栈空为false

注意：go语言可以用byte代表单个字符

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/20.有效的括号.go)

### 21.合并两个有序链表\(高频\)

题目：[两个升序链表，合并成一个](https://leetcode-cn.com/problems/merge-two-sorted-lists)

题解：

* 需要一个空的头节点做辅助，`head.Next`就是结果
* 每次遍历始终维护上一个节点`prev`，初始值`prev=head`
* 循环遍历两个链表，循环条件都不为空，每次把当前节点更小的取出来即可

```go
prev.Next = l1
prev = l1
l1 = l1.Next
```

* 最后加入有不为空的节点，则直接赋值

```go
if l1!=nil{
    prev.Next = l1
}else{
    prev.Next = l2
}
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/21.合并两个有序链表.go)

### 53.最大子序和\(高频\)

题目：[求和加起来最大的连续子数组](https://leetcode-cn.com/problems/maximum-subarray)

题解：

* 一次循环，数组里有可能出现负数，且只需要统计和即可
* 需要两个计数器，一个存储全局最大的子序列和，只要出现更大的就更新
* 另一个存储当前最大的子序和，判断当前最大子序和的方法是比较子序和与当前值哪个大
* 有可能当前值比子序列和大，就更新子序 `max(nums[i],nums[i]+last)`

核心代码

```go
last = max(nums[i],nums[i]+last)
resMax = max(resMax,last)
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/53.最大子序和.go)

### 70.爬楼梯

题目：[一次可以上一阶或者二阶，如果是n阶有多少种爬法](https://leetcode-cn.com/problems/climbing-stairs/)

题解：斐波那契数列，返回结果是前两个值的和，只需要保存前两个值和当前结果，递推赋值即可

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/70.爬楼梯.go)

### 101.对称二叉树

题目：[二叉树是不是镜像的](https://leetcode-cn.com/problems/symmetric-tree/description/)

题解：

* 递归，相当于使用了前序遍历和后续遍历相等的性质
* 函数签名`isMirror(root,root)`，判断前序和后续相等

```go
q.Val == p.Val && isMirror(p.Left,q.Right) && isMirror(p.Right,q.Left)
```

* 注意都为空时`true`，其中一个为空时`false`

```go
if p == nil && q == nil {
    return true
}
if p==nil || q==nil{
    return false
}
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/101.对称二叉树.go)

### 104.二叉树的最大深度

题目：[根节点到最远叶子节点的最长路径上的节点数](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/)

题解：递归，前序遍历，返回值为左右节点最大深度+1，退出条件为null节点返回0，左右子树都为空返回1

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/104.二叉树的最大深度.go)

### 121.买卖股票的最佳时机\(高频\)

题目：[给定整数数组表示每天股票价格，买一次卖一次求最大收益](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/)，要求必须先买再卖

题解：与目前最小值做差，得到当前最大值，更新最大值，一次循环。核心代码如下

```go
if v<minNum{
    minNum = v
}else if v - minNum > maxNum{
    maxNum = v - minNum
}
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/121.买卖股票的最佳时机.go)

### 136.只出现一次的数字

题目：[数组中某元素只出现一次，其余两次。找那个元素](https://leetcode-cn.com/problems/single-number/description/)

题解：

* 方法一、直接用map统计，超过一次就删掉
* 方法二、每个值都异或，最终得到的就是答案
* 异或的性质：`a^a=0, a^0=a`

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/136.只出现一次的数字.go)

### 141.环形链表\(高频\)

题目：[判断链表中是否有环](https://leetcode-cn.com/problems/linked-list-cycle/description/)

题解：快慢指针，相等时退出，注意循环条件

```go
p.Next != nil && q.Next != nil && q.Next.Next != nil
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/141.环形链表.go)

### 155.最小栈\(中频\)

题目：[设计一个栈，支持push、pop、top、getMin获取栈内最小值操作](https://leetcode-cn.com/problems/min-stack/description/)

题解：要自定义结构体，结构体内两个栈，一个存储push的元素的栈1，另一个存储最小值栈2，栈2有一个性质，每个栈2内元素位置为栈顶时，始终表示当前栈1最小的元素，比如

```text
栈1 [1 2 3 4 0]
栈2 [1 1 1 1 0]
```

这样出栈同同步出栈，始终最小值就是栈2的栈顶，使用了一个当前最优解的思路

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/155.最小栈.go)

### 160.相交链表\(高频\)

题目：[假定链表中没有循环，判断两个链表是否相交](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/)

题解：

法1（笨办法）

* 双指针同步走，肯定有一个先结束，假如另一个指针所指到结束位置的距离就是长链表比短链表多出来的部分
* 长链表跑到和短链表等长，然后同步走，找出相等节点（地址相同非值相同）

法2（最优）

* pA走过的路径为A链+B链, pB走过的路径为B链+A链
* 分别将另一个链表拼到尾部，相当于两个链表等长
* pA和pB走过的长度都相同，都是A链和B链的长度之和，相当于将两条链从尾端对齐，如果相交，则会提前在相交点相遇，如果没有相交点，则会在最后相遇。

```text
pA:1->2->3->4->5->6->null->9->5->6->null
pB:9->5->6->null->1->2->3->4->5->6->null
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/160.相交链表.go)

### 169.多数元素

题目：[数组中有一个数超过元素的一半，找出那个数](https://leetcode-cn.com/problems/majority-element/description/)

题解：

* 最先想到的是hash表，更好的办法是投票
* 随便选个人当选，和它不同就反对，票数--
* 和他相同就赞成，票数++
* 票数为0则换届，最终票数肯定是正的，当选的肯定是众数，代码较短我直接贴上来了

```go
func majorityElement(nums []int) int {
    var count,num int
    for _,v := range nums{
        if count == 0 || v == num{
            num = v
            count++
        }else{
            count--
        }
    }
    return num
}
```

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/169.多数元素.go)

### 206.反转链表\(高频\)

题目：

题解：

* 其实是234. 回文链表 的一部分,就是在考基本功

方法1、递归

* 首先要确认递归返回链表头，所以退出条件：head为空直接退出，递归内部head.Next为空返回头节点
* 递归传入`head.Next`返回已经反转好的链表头
* 递归后动作，需要让最后的头节点指空，也就是`head.Next = nil`在指空前，需要让`head.Next.Next=head`

```text
1->2->null
1<-2
nil<-1<-2
```

关键代码

```go
newHead := reverseList(head.Next)
head.Next.Next = head
head.Next = nil
return newHead
```

方法2、非递归（最快，省空间）

* 非递归两件套，`prev,curr = nil,head`
* 每次保存下一个节点，让当前节点指向上一个节点
* 然后向下走一位,`prev=curr; curr=next`,`curr`为空时停止，`prev`就是新的头节点

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/206.反转链表.go)

### 226.翻转二叉树

题目：[翻转二叉树，每一层都要完全翻转](https://leetcode-cn.com/problems/invert-binary-tree/description/)

题解：

方法1 非递归 不推荐 较为麻烦，容易漏掉

* 层遍历保存每层的状态，反着读出来每层就是翻转结果
* 但不能改变层遍历的过程，所以要两个栈，栈1保存层遍历，栈2保存该层的反向结果作为下次的`before`层
* 遍历栈2，左右节点指向倒着遍历栈1
* 要注意有可能出现部分空节点的情况，栈1，空位要留出来，所以要用nil来占位

方法2 递归:其实子节点还是属于父节点，只要翻转左右节点位置就行了

自身递归，“交换”左右子树时记得备份。 代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/226.翻转二叉树.go)

### 234.回文链表

题目：[判断一个链表是否为回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/description/)

题解： 法1、最简单，直接遍历一次转换成数组，判断数组是否回文即可 法2、快慢指针（整除器）

* 把剩下的一半变成逆序，再进行比较。注意奇偶情况讨论。递归非递归都行，想起来哪个用哪个，判断完后恢复链表
* 如果要快就边跑边让慢指针翻转链表，结束后也不用恢复

  代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/234.回文链表.go)

### 283.移动零

题目：[把数组里的零全部移动到结尾](https://leetcode-cn.com/problems/move-zeroes/description/)

题解：两个下标，使用类似于选择插入排序的方法，不断扩充非零列，剩余的元素用0填充

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/283.移动零.go)

### 448.找到所有数组中消失的数字

题目：[1-n的数字存储在长度为n的数组里，有的数字重复出现了，所以有的数字没有出现，找出没有出现的数字](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/)

题解： 法1、直接用hash表，比较简单，不过建议还是用法2可以体现水平 法2、用占位方法，遍历，出现的abs\(数字\)-1作为下标的数字改为负，如果已经是负的就不用改了，最后再遍历一次数组把存储数字为正位置的\(下标+1\)存储到结果集里

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/448.找到所有数组中消失的数字.go)

### 461.汉明距离

题目：[求两个数字二进制位不同的有多少个](https://leetcode-cn.com/problems/hamming-distance/description/)

题解：先亦或，然后`%2=1`时统计，`>>1`代表`/2`去掉一位

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/461.汉明距离.go)

### 543.二叉树的直径

题目：[求两个叶子之间最大距离](https://leetcode-cn.com/problems/diameter-of-binary-tree/description/)

题解：

* 深度优先dps，必须用递归，递归返回的是左右子树最深深度
* 维护一个最大值，递归返回后判断左右子树贡献的深度和，与最大值哪个大，更新最大值，这样可以保证直径是当前最大 `max(x+y,maxRes)`
* 返回当前子树最大深度`return max(x,y)+1`

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/543.二叉树的直径.go)

### 617.合并二叉树

题目：[合并二叉树，同位置都有值就加起来，有一个为空就只合并](https://leetcode-cn.com/problems/merge-two-binary-trees/description/)

题解：

* 递归，返回条件：其中一个为空返回另一个节点
* 如果两个都不为空，加起来
* 处理递归函数，分别传入两棵树的左子树或右子树，赋值给当前节点左右子树
* 跟437 路径总和III的思想是一样的。

代码：[golang](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/all/617.合并二叉树.go)

## 最后

如果文中有误，欢迎提pr或者issue，**一旦合并或采纳作为贡献奖励可以联系我直接无门槛**加入[技术交流群](https://mp.weixin.qq.com/s/ErQFjJbIsMVGjIRWbQCD1Q)

我是小熊，关注我，知道更多不知道的技术

![](https://github.com/minibear2333/interview-leetcode/tree/f218b102b41d5ce6b95e9b305fdc326205b4e2f9/LeetCode/hot100/res/2021-03-17-19-57-33.png)

