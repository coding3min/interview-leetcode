//题目链接：https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/?envType=study-plan&id=lcof
package main

import "math"

//本题与 day24-剪绳子 题目区别仅在于答案需要取模，解题思路不做复述了


//我刚开始的想法是，对运行结果进行一次取余就可以，代码如下

func cuttingRope(n int) int {
	mod := math.Pow(10,9) + 7
	if n <= 3{
		return n-1
	}
	a,b := n/3,n%3
	res := float64(0)
	if b == 0{
		res = math.Pow(3,float64(a))
	} else if b == 1{
		res = math.Pow(3,float64(a-1)) * 4
	} else {
		res = math.Pow(3,float64(a)) * 2
	}
	return int(res) % int(mod)
}
//提交之后发现结果错误，我忽视了指数运算的增长速度

//本题解参考自链接：https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/
//原来在求 3 ^ a 时已经溢出，现在需要考虑的是取模的处理。本题需要取模的原因是 n 值的取值范围扩大了，导致计算过程中求 3 的次方时，数值指数扩大，
//需考虑溢出情况。我们的取余应该在做指数运算后进行
//
//两种解决方案：循环求余 和 快速幂求余，两种方法均基于以下运算规则
//(x∗y)%p=[(x%p)(y%p)]%p
//
//一种比较通俗的解法是 循环取余：
//
//x^a % p=[(x^(a−1)%p) * (x%p)] % p=[(x^(a−1)%p) * x]%p
//
//本题中 a=3，p=10^9+7，利用此公式，可通过循环操作依次取x、x^2、x^3、...、x^a 对 p 的余数，保证每轮中间值在取值范围内。
func cuttingRope_2(n int) int {
	x := int(math.Pow(10,9) + 7)
	if n<=3{
		return n-1
	}
	a,b := n/3,n%3
	if b == 0{
		return remainder(3,a,x)
	}
	if b == 1{
		return remainder(3,a-1,x) * 4 %x
	}
	return remainder(3,a,x)%x *2 %x
}
// 循环求余法
func remainder(a,b,p int) int {
	rem := 1
	for i:=0;i<b;i++{
		rem = (rem*a)%p
	}
	return rem
}


//第二种解决方案是快速幂求余
//我们在 day20 的 求数值的整数次方 中已经讲解过了快速幂，这里就不再赘述了，再将其每次操作进行取余即可。
// 快速幂求余法
func fastmul_remainder(a,b,p int) int {
	rem := 1
	for b>0{
		if b&1==1{
			rem = (rem*a)%p
		}
		a = a*a % p
		b /= 2
	}
	return rem
}

//对这道题的取余我想再多说几句，关于为什么可以这样做，让我困惑了挺长一段时间，下面假设很不合理，但这确实是我的困惑所在。
//我做了一个这样的假设，a=3，b=6，p=90，范围最大值为100，3^6=486，远大于100，产生溢出情况，用循环求余解决溢出，3^4=81，对90取余还是81,3^5=243，
//会直接溢出，还是没有避免溢出的情况啊。
//然后，我是这么想的，题目中 p=10^9+7，MaxInt32=2147483647，题目没有设定数据范围，那先假设整数为最大为 int64，p要小于 MaxInt32，
//那一个小于p的数乘以一个小于p的数，结果一定在 int64 的取值范围内！这下就反应过来，我之前的假设是非常不合理的。
//
//另外，10^9+7这个数好像经常用用于取模，为什么是这么数字，我查了一下大概是这么说的，一方面，该数字要足够大，
//另一方面，需要是一个大的质数来减少冲突。而10^9+7这个数字，相加不超过 int32，相乘不超过 int64.
//一般来说x的选取只要10^x＋7保证比初始输入数据的范围大就可以。比如有些数据范围小的题为了避免用long long而把模数设定为10007。
//
//这道题中主要用到 10^9+7 相乘不超过 int64 的性质。