package main

import "fmt"

//21 22 23 24 25 26
//20  7  8  9 10
//19  6  1  2 11
//18  5  4  3 12
//17 16 15 14 13
//-2 -1 0  1  2
//
//示例
//
//输入样例 ：
//
//0 0
//-1 -1
//输出样例 ：
//
//1
//5
func main(){
	fmt.Println(getNum(0,0))
	fmt.Println(getNum(1,0))
	fmt.Println(getNum(1,-1))
	fmt.Println(getNum(0,-1))
	fmt.Println(getNum(-1,-1))
	fmt.Println(getNum(-1,0))
	fmt.Println(getNum(-1,1))
	fmt.Println(getNum(0,1))
	fmt.Println(getNum(1,1))
}

func getNum(x,y int) int{
	if x == y && x == 0{
		return 1
	}
	//先去掉内层面积，可以得到起始点值
	//内层面积先求内层边长，side := (max-1) * 2 + 1
	//内层面积：area := side*side 所以起始点为 begin := area + 1
	//可以根据外层最大坐标值算出起始点坐标，(max,max-1)
	//从起始点位置开始遍历坐标，规律为下左上右，对应 y- x- y+ x+，同时num++
	//当达到边角以后需要转向，需要一个计数器统计当前方向 0 1 2 3 对应  y- x- y+ x+
	//可以发现0 2是偶数改变y，1 3是奇数改变x，此时再确定符号
	// 0 2平移一位变成 -1 1 可以控制y-1或者y+1，1 3平移两位变成-1 1控制x-1或者x+1
	//计数器数值变化需要在outSide-1的情况变化一次 所以 / (outSide-1)
	//计数器需要在outSide-1的时候就要变化，所以到达outSide-1时必须计数为outSide-1的倍数
	//num - begin = 0 所以 num - begin + 1
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
func abs(x int) int{
	if x<0{
		return -x
	}
	return x
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}


