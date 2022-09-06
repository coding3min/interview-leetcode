// 题目链接：https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/?envType=study-plan&id=lcof


package main

//这道题的通俗解法是双层循环（也就是暴力解法）
//假设共有 n 天，最大利润初始化为 0，第一层循环中设定买股票的时间点，从第 1 天遍历到 第 n-1 天，
//第二层循环遍历卖股票的时间点，从买股票的第二天遍历到最后一天，若买股票的价格大于买股票的价格，更新最大利润。
func maxProfit(prices []int) int {
	res := 0
	n := len(prices)
	for i:=0;i<n-1;i++{
		for j:=i+1;j<n;j++{
			if prices[j] - prices[i] > 0{
				res = max(res,prices[j]-prices[i])
			}
		}
	}
	return res
}


//方法2：一次遍历
//遍历股票价格的过程中，维护一个股票最小价格，若当前遍历到的价格小于该最小价格，则更新最小价格，否则，更新最大利润。
//为避免函数命名冲突，次函数名添加后缀 “_2”
func maxProfit_2(prices []int) int {
	n := len(prices)
	if n == 0{
		return 0
	}
	min_price := prices[0]
	res := 0
	for i:=1;i<n;i++{
		if prices[i] < min_price{
			min_price = prices[i]
		} else {
			res = max(res,prices[i]-min_price)
		}
	}
	return res
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

//这道题和动态规划有关系吗？我觉得关系不大吧，用递推更加合适一些。
//主要是场景过于简单，只需要维护一个最小价格，每个阶段只有一个状态，且每个阶段的状态与之前阶段均无关。