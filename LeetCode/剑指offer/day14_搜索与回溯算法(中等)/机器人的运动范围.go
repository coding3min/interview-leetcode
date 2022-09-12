// 题目链接：https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/?envType=study-plan&id=lcof
// day14/31
// 第 14 天主题为：搜索与回溯算法（中等）
// 包含两道题目：
// 剑指offer12.矩阵中的路径
// 面试题13.机器人的运动范围

package main

//解题思路：BFS 或 DFS，要频繁地计算数字的各位数之和，我们可以声明一个函数专门用于计算该数字的位数之和，具体操作时不断取余和对 10 整除。
//
//方法1：DFS
//我们要统计机器人能够到达的格子数量，（0，0）是机器人的起始坐标，那我们从该坐标开始，向外 BFS，同时用一个used数组记录到达过的格子，避免重复统计一个格子。此外，我们要走的尽可能远，所以我们是不需要往上或者往左走的，只需要往远处行进（右或者下）。
//
//具体操作：
//DFS 的参数：i 和 j，代表当前即将访问的格子下标
//DFS 结束条件：
//索引越界、两个索引的位数之和大于 k 或者 该索引已经访问过，这些情况下，说明该格子无法到达或者已访问过，返回 0 即可。
//搜索：首先，当前位置可访问，标记当前位置，used[i][j] = true，然后利用系统栈实现DFS，向右和向下出发。
func movingCount(m int, n int, k int) int {
	used := make([][]bool,m)
	for i:=0;i<m;i++{
		used[i] = make([]bool,n)
	}
	var dfs func(i,j int) int
	dfs = func(i,j int) int{
		if i >=m || j >= n || bitsum(i)+bitsum(j)>k{
			return 0
		}
		if used[i][j]{
			return 0
		}
		used[i][j] = true
		return 1 + dfs(i,j+1) + dfs(i+1,j)
	}
	return dfs(0,0)
}

func bitsum(x int) int {
	res := 0
	for x > 0{
		res += x % 10
		x /= 10
	}
	return res
}


//方法2：BFS
//本题中，虽然我们也使用 used 数组记录每个位置是否被访问过，但是我们不需要回溯操作，所以是可以使用 BFS 的。
//BFS通常需要借助队列的先进先出特性来实现

func movingCount_2(m int, n int, k int) int {
	used := make([][]bool,m)
	for i:=0;i<m;i++{
		used[i] = make([]bool,n)
	}
	res := 0
	// 初始化队列只有元素(0,0)
	q := [][]int{{0,0}}
	for len(q) != 0{
		x,y := q[0][0],q[0][1]
		if used[x][y]{
			q = q[1:]
			continue
		}
		res ++
		used[x][y] = true
		// 检查参数合法性 及 该位置是否未被访问过
		// try right
		if y+1 < n && (bitsum(x) + bitsum(y+1)) <= k{
			q = append(q,[]int{x,y+1})
		}
		// try blow
		if x+1 < m && (bitsum(x+1) + bitsum(y)) <= k{
			q = append(q,[]int{x+1,y})
		}
		q = q[1:]
	}
	return res
}


// 有个大坑！错误代码如下，我排查了一个多小时才发现问题所在，我这里就不具体说了，好心累
// 我这里错误就是一个未被访问过的元素多次进入队列！
func movingCount_3(m int, n int, k int) int {
	used := make([][]bool,m)
	for i:=0;i<m;i++{
		used[i] = make([]bool,n)
	}
	res := 0
	q := [][]int{{0,0}}
	for len(q) != 0{
		x,y := q[0][0],q[0][1]
		res ++
		used[x][y] = true
		// 检查参数合法性 及 该位置是否未被访问过
		// try right
		if y+1 < n && (bitsum(x) + bitsum(y+1)) <= k && !used[x][y+1]{
			q = append(q,[]int{x,y+1})
		}
		// try blow
		if x+1 < m && (bitsum(x+1) + bitsum(y)) <= k && !used[x+1][y]{
			q = append(q,[]int{x+1,y})
		}
		q = q[1:]
	}
	return res
}