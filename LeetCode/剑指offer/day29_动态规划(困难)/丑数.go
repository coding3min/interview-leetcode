//题目链接：https://leetcode.cn/problems/chou-shu-lcof/?envType=study-plan&id=lcof
package main

import "container/heap"

//丑数的递推性质： 丑数只包含因子 2, 3, 5，因此有 “丑数 = 某较小丑数 * 某因子” （因子为 2、3、5）。
//该性质是解题的关键

//刚开始我们只拥有丑数 1，然后 经过 1 与因子 2、3、5 相乘，得到三个丑数，将该三个丑数再次与因子相乘，
//又得到部分丑数（数目无法确定，因为可能有重复出现的数字），我们不断经过此步骤，可以得到很多丑数
//
//那如何得到第 n 个呢？
//上面得到的丑数是无序的，我们可以使用动态规划得到排好序的丑数，丑数数字 dp 初始化时只有一个元素 1，三个因子的指针 p2、p3、p5 均指向 1，
//计算得到 三个因子指向的丑数 与 该因子的 乘积，取最小值，就是下一个丑数，然后找到对应的指针，将该指针自增1，指向下一个丑数，
//每次可以得到一个丑数，n-1次循环可得到第 n 个丑数。
//
//动态规划三步：
//1. 定义dp数组大小及下表含义：dp[i] 代表第 i 个丑数
//2. dp 数组状态初始化：dp[1] = 1，三个指针 p2,p3,p5=1,1,1
//3. 状态转移方程，dp[i] = min(dp[p2]*2,dp[p3]\*3,dp[p5]\*5)，之后找到对应的指针，将该指针自增1

func nthUglyNumber(n int) int {
	dp := make([]int,n+1)
	// dp 数组初始化，只有 1 一个丑数
	dp[1] = 1
	// 三个指针初始化指向第一个丑数
	p2,p3,p5:= 1,1,1
	for i:=2;i<=n;i++{
		// 寻找三个指针指向元素与对应因子乘积的最小值
		num := min(dp[p2]*2,min(dp[p3]*3,dp[p5]*5))
		dp[i] = num
		// 找到对应指针，该指针右移（即自增1）
		// 可能对应不止一个指针
		if num == dp[p2] * 2{
			p2++
		}
		if num == dp[p3] * 3{
			p3++
		}
		if num == dp[p5] * 5{
			p5++
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y{
		return x
	}
	return y
}


//此题我们还可以用小根堆来解决
//
//初始时堆为中只有第一个丑数 1 。
//
//每次取出堆顶元素 x，则 x 是堆中最小的丑数，由于 2x, 3x, 5x 也是丑数，因此将 2x, 3x, 5x 加入堆。
//
//上述做法会导致堆中出现重复元素的情况。为了避免重复元素，可以使用哈希集合去重，避免相同元素多次加入堆。
//
//在排除重复元素的情况下，第 n 次从最小堆中取出的元素即为第 n 个丑数。

type maxH []int

func (this maxH) Len() int{ return len(this)}
func (this maxH) Less (i,j int) bool {
	return this[i] < this[j]
}
func (this maxH) Swap (i,j int){
	this[i],this[j] = this[j],this[i]
}
func (this *maxH) Push(v interface{}){
	*this = append(*this,v.(int))
}
func (this *maxH) Pop() interface{}{
	old := *this
	n := len(old)
	res := old[n-1]
	*this = old[:n-1]
	return res
}


func nthUglyNumber_2(n int) int {
	uglys := &maxH{1}
	factors := []int{2,3,5}
	record := map[int]struct{}{}
	record[1] = struct{}{}
	for i:=1;;i++{
		num := heap.Pop(uglys).(int)
		if i == n{
			return num
		}
		for _,f := range factors{
			next := num * f
			if _,has := record[next];!has{
				heap.Push(uglys,next)
				record[next] = struct{}{}
			}
		}
	}
}