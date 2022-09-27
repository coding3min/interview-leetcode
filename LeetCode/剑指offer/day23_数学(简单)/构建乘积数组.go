//题目链接：https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/?envType=study-plan&id=lcof

package main

//若可以使用除法的话，我们可以先一次遍历，得到所有数字的乘积 x，然后创建新的数组，
//一次遍历，每次 x 除以当前遍历到的元素，添加到新数组即可。
//
//不使用乘法，我的想法是，构建前缀乘积数组 prefix 和后缀乘积数组 suffix，
//长度均与给定数组长度相同，设 n 为给定数组长度
//
//- prefix[i] = a[0] * a[1] * ... * a[i-1]
//- suffix[i] = a[i+1] * a[i+2] * ... * a[n-1]
//
//然后创建乘积数组 res, res[i] = prefix[i] * suffix[i]

func constructArr(a []int) []int {
	n := len(a)
	//当输入数组长度为 0 时，直接返回空数组，
	//若不进行处理，会在下面代码 suffix[n-1]报错超出索引
	if n == 0{
		return []int{}
	}
	// 构建前缀乘积数组和后缀乘积数组
	prefix,suffix := make([]int,n),make([]int,n)
	res := make([]int,n)
	prefix[0],suffix[n-1] = 1,1
	for i:=1;i<n;i++{
		prefix[i] = prefix[i-1] * a[i-1]
	}
	for i:=n-2;i>=0;i--{
		suffix[i] = suffix[i+1] * a[i+1]
	}
	for i:=0;i<n;i++{
		res[i] = prefix[i] * suffix[i]
	}
	return res
}

//时间空间复杂度均为 O(n)
//
//思考如何进行优化，时间复杂度的 O(n) 是没有优化空间的，只能从空间复杂度下手
//
//通过两轮循环，分别计算 prefix 和 suffix 的乘积。
//先从上到下计算前缀乘积数组，再从下到上计算后缀乘积数组，因为不需要额外空间存储，
//只用 res 数组即可，将空间复杂度从 O(n) 降低至 O(1) （此空间复杂度不包含返回数组）。

func constructArr_2(a []int) []int {
	n := len(a)
	if n == 0{
		return []int{}
	}
	res := make([]int,n)
	res[0] = 1
	// 计算前缀树组乘积
	for i:=1;i<n;i++{
		res[i] = res[i-1] * a[i-1]
	}
	//计算后缀数组乘积
	temp := 1
	for i:=n-2;i>=0;i--{
		temp *= a[i+1]
		res[i] *= temp
	}
	return res
}
