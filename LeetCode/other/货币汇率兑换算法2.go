package main

import "fmt"

//知道各货币兑换另一种货币的汇率，要求求最大兑换和最小可兑换货币(注意，有可能有的货币无汇率信息)
//
//USD - RMB 1:6
//HKD - JAN 1:20
//
//示例：
//```
//输入HKD = 65
//最大可兑换  HKD=65
//最小可兑换  JAN=2600
//
//输入RMB=65
//最大可兑换 USD=10,RMB=5
//最小可兑换 RMB=65
//```
//
//函数
//```
//// 顺序输入UDB RMB HKD 汇率，比如上面的例子，输入[6 0 20],0表示不能兑换
//// isMax = true 时求最大可兑换 false 时求最小可兑换
//// moneyType = 0 1 2 3 对应 USD RMB HKD JAN，money对应货币数量
//// 返回兑换情况，输入moneyType=2 money=65 输出最大[0,0,65,0]，最小[0,0,0,1300]
//// 返回兑换情况，输入moneyType=1 money=65 输出最大[10,5,0,0]，最小[0,65,0,0]
//func xxx(rate []int,moneyType int,money int,isMax bool) []int{
//
//
//}
//```

func main(){
	rate := []int{6,0,20}
	fmt.Println(xxx(rate,2,65,true))
	fmt.Println(xxx(rate,2,65,false))
	fmt.Println(xxx(rate,1,65,true))
	fmt.Println(xxx(rate,1,65,false))
}
func xxx(rate []int,moneyType int,money int,isMax bool) []int{
	res := make([]int,len(rate)+1)
	res[moneyType] = money
	rate = append(rate,1)
	if isMax{
		for i:=moneyType;i>0 && rate[i-1] != 0;i--{
			res[i-1] = res[i]/rate[i-1]
			res[i] = res[i]%rate[i-1]
		}
		return res
	}
	for i:=moneyType; i < len(res)-1 && rate[i] != 0; i++ {
		res[i+1] = res[i] * rate[i]
		res[i] = 0
	}
	return res
}


