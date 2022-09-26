// 题目链接：https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/?envType=study-plan&id=lcof
// day28/31
// 第 28 天主题为：搜索与回溯算法（困难）
// 包含两道题目：
// 剑指offer37.序列化二叉树
// 剑指offer38.字符串的排列
package main

import (
	"fmt"
	"sort"
)

//全排列问题，如果题目所给字符不包含重复字符，用下面简单回溯即可：
func permutation(s string) []string {
	res := []string{}
	n := len(s)
	used := make([]bool,n)
	var backtrace func(path string)
	backtrace = func(path string){
		if len(path) == n{
			res = append(res,path)
			return
		}
		for i:=0;i<n;i++{
			if used[i]{
				continue
			}
			path += string(s[i])
			used[i] = true
			backtrace(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrace("")
	return res
}

//本题的难点在于不能有重复元素，这也就意味着输入字符串中包含重复元素。上面的代码不再可行
//用集合去重是可行的，但太过笨重，试想是否可以改善回溯过程，得到不含重复元素的排列。
//
//去重，很容易想到要先对元素进行排序，这样才能判断相邻的节点是否重复使用了。
//回溯的问题在纸上都可以画成一个树的结构，横向迭代+纵向递归。
//针对本题，画出该树结构，可以发现同一树层上，若当前节点与上一个节点值相同，则对该节点回溯的结果与上一节点相同，意味着可以跳过此节点，迭代至下一节点。
//那同一个树枝呢？同一树枝上，若上一个节点使用过，该节点还是可以使用的，因为这是一个正常的排列。
//接下来就是代码实现，当前元素与上一个元素值相同的情况下，判断是在同一树枝，还是同一树层。
//	used[i-1]=true 说明上一节点已使用，说明是同一树枝
//	used[i]=false，说明为同一树层，可跳过此次回溯

func permutation_2(s string) []string {
	runeS := []byte(s)
	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})
	n := len(runeS)
	res := []string{}
	used := make([]bool,n)
	var backtrace func(path []byte)
	backtrace = func(path []byte) {
		if len(path) == n{
			res = append(res,string(path))
			return
		}
		for i:=0;i<n;i++{
			if used[i]{
				continue
			}
			if i>0 && runeS[i-1]==runeS[i] && !used[i-1]{
				continue
			}
			used[i] = true
			path = append(path,runeS[i])
			backtrace(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrace([]byte{})
	fmt.Println(runeS)
	return res
}

//本题解题思路与 LeetCode 48：全排列 II 完全一致，有兴趣可以再去做做练习一下