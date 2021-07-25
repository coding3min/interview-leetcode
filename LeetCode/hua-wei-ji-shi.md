# 华为机试

## 字符串余

输入若干字符串和一个数字n，按8切分字符串，长度不足补0

示例：

```text
输入：asdf 123456789
输出：asdf0000 12345678 90000000
```

```go
func splitStr(str string,n int){
    zeroStr := "00000000"
    for len(str)>8{
        fmt.Println(str[:8])
        str = str[8:]
    }
    fmt.Println(str+zeroStr[:8-len(str)])
}
```

## 货币汇率兑换算法

知道各货币兑换另一种货币的汇率，要求求最大兑换和最小可兑换货币

USD - RMB 1:6 RMB - HKD 1:2 HKD - JAN 1:20

示例：

```text
输入HKD = 65
最大可兑换  USD=10,RMB=5
最小可兑换  JAN=2600
```

函数

```text
// 顺序输入UDB RMB HKD 汇率，比如上面的例子，输入[6 2 20]
// moneyType = 0 1 2 3 对应 USD RMB HKD JAN，money对应货币数量
// isMax = true 时求最大可兑换 false 时求最小可兑换
// 返回兑换情况，最大[5,2,1,0]，最小[0,0,0,1300]
func xxx(rate []int,moneyType int,money int,isMax bool) []int{

}
```

解法,代码见[货币汇率兑换算法.go](https://github.com/coding3min/interview-leetcode/tree/5cb2e5224e55a017aa569f71ee31714371f1a8cd/LeetCode/other/货币汇率兑换算法.go)

```go
func xxx(rate []int,moneyType int,money int,isMax bool) []int{
    res := make([]int,len(rate)+1)
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
```

## 货币汇率兑换算法2

知道各货币兑换另一种货币的汇率，要求求最大兑换和最小可兑换货币\(注意，有可能有的货币无汇率信息\)

USD - RMB 1:6 HKD - JAN 1:20

示例：

```text
输入HKD = 65
最大可兑换  HKD=65
最小可兑换  JAN=2600

输入RMB=65
最大可兑换 USD=10,RMB=5
最小可兑换 RMB=65
```

函数

```text
// 顺序输入UDB RMB HKD 汇率，比如上面的例子，输入[6 0 20],0表示不能兑换
// isMax = true 时求最大可兑换 false 时求最小可兑换
// moneyType = 0 1 2 3 对应 USD RMB HKD JAN，money对应货币数量
// 返回兑换情况，输入moneyType=2 money=65 输出最大[0,0,65,0]，最小[0,0,0,1300]
// 返回兑换情况，输入moneyType=1 money=65 输出最大[10,5,0,0]，最小[0,65,0,0]
func xxx(rate []int,moneyType int,money int,isMax bool) []int{


}
```

解法，代码见[货币汇率兑换算法2](https://github.com/coding3min/interview-leetcode/tree/5cb2e5224e55a017aa569f71ee31714371f1a8cd/LeetCode/other/货币汇率兑换算法2.go)

## 螺旋矩阵变体

输入一个坐标，直接输出对应的值，比如输入横坐标-4 纵坐标6输出什么（里面是个正方形，边长乘以边长就是结果）

21 22…… 20 7 8 9 10 19 6 1 2 11 18 5 4 3 12 17 16 15 14 13

示例

```text
输入样例 ：

0 0
-1 -1
输出样例 ：

1
5
```

先去掉内层面积，可以得到起始点值

内层面积先求内层边长，side := \(max-1\) \* 2 + 1

内层面积：area := side\*side 所以起始点为 begin := area + 1

可以根据外层最大坐标值算出起始点坐标，\(max,max-1\)

从起始点位置开始遍历坐标，规律为下左上右，对应 y- x- y+ x+，同时num++

当达到边角以后需要转向，需要一个计数器统计当前方向 0 1 2 3 对应 y- x- y+ x+

可以发现0 2是偶数改变y，1 3是奇数改变x，此时再确定符号

0 2平移一位变成 -1 1 可以控制y-1或者y+1，1 3平移两位变成-1 1控制x-1或者x+1

计数器数值变化需要在outSide-1的情况变化一次 所以 / \(outSide-1\)

计数器需要在outSide-1的时候就要变化，所以到达outSide-1时必须计数为outSide-1的倍数时 num - begin = 0 所以 num - begin + 1

```go
func getNum(x,y int) int{
    if x == y && x == 0{
        return 1
    }
    maxNum := max(abs(x),abs(y))
    side := (maxNum - 1) * 2 + 1
    begin  := side * side + 1
    outSide := side + 2
    num := begin
    for cnt,i,j:=0,maxNum,maxNum-1;cnt<4;{
        if i == x && j == y{
            return num
        }
        if cnt%2 == 0{
            j = j + cnt - 1
        }else{
            i = i + cnt - 2
        }
        num++
        cnt = (num - begin + 1) / (outSide - 1)
    }
    return num
}
```

测试用例,输出1 2 3 4 5 6 7 8 9

```go
    fmt.Println(getNum(0,0))
    fmt.Println(getNum(1,0))
    fmt.Println(getNum(1,-1))
    fmt.Println(getNum(0,-1))
    fmt.Println(getNum(-1,-1))
    fmt.Println(getNum(-1,0))
    fmt.Println(getNum(-1,1))
    fmt.Println(getNum(0,1))
    fmt.Println(getNum(1,1))
```

