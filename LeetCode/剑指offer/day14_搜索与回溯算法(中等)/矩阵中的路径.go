// 题目链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/

package main


//解题思路：DFS+回溯
//题目出现网格，我们要寻找一条路径，很容易想到用 DFS 或 BFS，因为同一单元格内的字母不允许被重复使用，所以每使用一个字母，我们要对其进行记录，
//我们每一步很可能有多种选择，为避免使用到本条路径使用过的网格，我们需要进行回溯，用 used 二维数组记录每个元素在本条路径中是否被访问过。
//
//另外，遇到这类需要上下左右移动的题目，可提前声明一个二维数组，保存上下左右移动时的元素下标变化，方便之后操作。

//具体流程如下：
//递归参数：i、j 和 k，i 和 j 为当前访问的网格元素下标，k 为当前目标字符在 word 中的下标。

//终止条件：
//- board[i][j] != word[k] 即字符不匹配，说明此条路径走不通，返回 false，向前回溯尝试其他路径
//- board[i][j] = word[k]，且 k=len(word)-1，说明word所有字符匹配完成，存在指定路径，返回 true
//
//搜索与回溯：
//当前字符已经匹配成功，先标记当前下标为已访问，used\[i\]\[j]=true，函数结束时向前回溯需要取消此标记，结合go语言的特性，我们在这里用 defer 进行取消标记。
//之后向四个方向进行搜索，先验证移动后坐标的合法性，然后判定该位置是否已被访问过，若合法且未访问过，在该位置进行递归操作，参数需要作出相应变化（坐标变化对应方向，k+=1）。
func exist(board [][]byte, word string) bool {
	m,n := len(board),len(board[0])
	// used大小与board相同，用于记录每个字符是否被访问过
	used := make([][]bool,m)
	for i:=0;i<m;i++{
		used[i] = make([]bool,n)
	}
	directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}

	// 从board的 i，j 索引处开始寻找单词word[K:]
	var backtrace func(i,j,k int) bool
	backtrace = func(i,j,k int) (res bool) {
		//若字符不匹配，直接返回 false
		if board[i][j] != word[k]{
			return false
		}
		// 字符匹配且已经匹配到了word的最后一个字符，说明匹配成功
		if k == len(word)-1{
			return true
		}
		// 当前位置匹配成功，需要递归地匹配之后的字符
		// 将当前位置标记为 已访问过
		used[i][j] = true
		// 函数结束时取消当前位置访问标记，进行回溯
		defer func() {
			used[i][j] = false
		}()
		// 或 如下写法
		//defer func(used [][]bool){
		//	used[i][j] = false
		//}(used)
		for _,dir := range directions{
			x,y := i+dir[0],j+dir[1]
			// 验证参数的合法性
			//坐标是否超出索引范围，且该位置是否被访问过
			if 0<=x && x<=m-1 && 0<=y && y<=n-1 && !used[x][y]{
				// 如果某位置匹配成功，说明整体可以匹配成功
				if backtrace(x,y,k+1){
					return true
				}
			}
		// 若取消标记用如下注释掉的操作代替 defer 的话，会失败
		// 具体为什么，我还没搞清楚
		//used[i][j] = false
		}
		// 四个方向均未匹配成功，当前路径无法走通
		return false
	}
	// 从 board 的每个字符开始对 word 从下标 0 开始匹配
	// 只要存在匹配成功的索引，返回 true
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if backtrace(i,j,0){
				return true
			}
		}
	}
	return false
}

//刚开始说到，出现网格，很容易想到用 DFS 或 BFS，那为什么我们选择使用 DFS 而没有选择使用 BFS 呢？
//因为我们要进行回溯操作！
//在 DFS 中，我们在一条路走不通后进行回溯，契合我们对位置打标记记录其是否被访问过的思路，刚进入时做标记，在此道路尝试过所有路径后，进行回溯，取消标记。
//如果用 BFS 的话，很难实现这种打标记的思路。