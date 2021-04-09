package main

import "fmt"

//知道各货币兑换另一种货币的汇率，要求求最大兑换和最小可兑换货币
//
//USD - RMB 1:6
//RMB - HKD 1:2
//HKD - JAN 1:20
//
//示例：
//```
//输入HKD = 65
//最大可兑换  USD=10,RMB=5
//最小可兑换  JAN=2600
//```
//
//函数
//```
// 顺序输入UDB RMB HKD 汇率，比如上面的例子，输入[6 2 20]
// moneyType = 0 1 2 3 对应 USD RMB HKD JAN，money对应货币数量
// isMax = true 时求最大可兑换 false 时求最小可兑换
// 返回兑换情况，最大[5,2,1,0]，最小[0,0,0,1300]
//func xxx(rate []int,moneyType int,money int,isMax bool) []int{
//
//}
//```

func main(){
	rate := []int{6,2,20}
	fmt.Println(xxx(rate,2,65,true))
	fmt.Println(xxx(rate,2,65,false))
}
func xxx(rate []int,moneyType int,money int,isMax bool) []int{
	res := make([]int,4)
	res[moneyType] = money
	rate = append(rate,1)
	if isMax{
		for i:=moneyType;i>0;i--{
			res[i-1] = res[i]/rate[i-1]
			res[i] = res[i]%rate[i-1]
		}
		return res
	}
	for i:=moneyType; i < len(res)-1; i++ {
		res[i+1] = res[i] * rate[i]
		res[i] = 0
	}
	return res
}


