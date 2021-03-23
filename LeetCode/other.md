### 其他高频

### medium
### 120.三角形最小路径和

题目：[求三角形的最小路径和](https://leetcode-cn.com/problems/triangle/description/)

要求是每一步只能找下一行的相邻节点，也就是当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 

```other
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
```

题解：从三角形底部向上遍历

```other
[
[2],
[3,4],
[6,5,7],
[4,1,8,3]
]
相邻结点：与(i, j) 点相邻的结点为 (i + 1, j) 和 (i + 1, j + 1)。
```

* 用一个一维度数组`dp [n+1]int`来存储最终结果，保证每个值都是最优解
* 方程 `min(dp[j],dp[j+1]) + triangle[i][j]`

注意：

代码：[golang](../LeetCode/all/120.三角形最小路径和.go)

引用：[递归+记忆化+DP](https://leetcode-cn.com/problems/triangle/solution/di-gui-ji-yi-hua-dp-bi-xu-miao-dong-by-sweetiee/)

### difficult
### 

题目：[]()

题解：

注意：

代码：[golang](../)