// 题目链接：https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/?envType=study-plan&id=lcof
// day17/31
// 第 17 天主题为：排序（中等）
// 包含两道题目：
// 剑指offer40.最小的k个数
// 剑指offer41.数据流的中位数

package main

import "container/heap"

//解题思路：用两个堆来维护数据流，将数据流根据元素的大小一分为二
//大根堆维护数据流元素值较小的一半，小根堆维护数据流元素值较大的一半
//当数据流长度为 偶数 时，大小根堆长度相同，当数据流长度为奇数时，我们规定，小根堆存储中位数。

//向数据流添加元素 num 时，可分为两种情况
//
//- 1.大根堆与小根堆长度相等，此时，应向小根堆添加元素，但添加的元素并不一定是 num
//- 	若 num 大于小根堆堆顶元素值，则 num 属于元素值较大的一部分，num 直接插入大根堆
//- 	否则，先将 num 插入大根堆，然后取出小根堆堆顶元素，插入大根堆
//- 2.大根堆与小根堆长度不相等，此时，应向大根堆添加元素，但添加的元素同样并不一定是 num
//- 	若 num 大于小根堆堆顶元素值，将 num 插入小根堆，然后取出小根堆堆顶元素，插入大根堆
//- 	否则，num 直接插入大根堆


type maxHeap []int  // 大顶堆
type minHeap []int  // 小顶堆

// 每个堆都要heap.Interface的五个方法：Len, Less, Swap, Push, Pop
// 其实只有Less的区别。

// Len 返回堆的大小
func (m maxHeap) Len() int {
	return len(m)
}
func (m minHeap) Len() int {
	return len(m)
}

// Less 决定是大优先还是小优先
func (m maxHeap) Less(i, j int) bool {  // 大根堆
	return m[i] > m[j]
}
func (m minHeap) Less(i, j int) bool {  // 小根堆
	return m[i] < m[j]
}

// Swap 交换下标i, j元素的顺序
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Push 在堆的末尾添加一个元素，注意和heap.Push(heap.Interface, interface{})区分
func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}
func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

// Pop 删除堆尾的元素，注意和heap.Pop()区分
func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n - 1]
	*m = old[:n - 1]
	return v
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	v := old[n - 1]
	*m = old[:n - 1]
	return v
}

// MedianFinder 维护两个堆，前一半是大顶堆，后一半是小顶堆，中位数由两个堆顶决定
type MedianFinder struct {
	maxH *maxHeap
	minH *minHeap
}

// Constructor 建两个空堆
func Constructor() MedianFinder {
	return MedianFinder{
		new(maxHeap),
		new(minHeap),
	}
}


func (m *MedianFinder) AddNum(num int)  {
	if m.maxH.Len() == m.minH.Len() {
		if m.minH.Len() == 0 || num >= (*m.minH)[0] {
			heap.Push(m.minH, num)
		} else {
			heap.Push(m.maxH, num)
			top := heap.Pop(m.maxH).(int)
			heap.Push(m.minH, top)
		}
	} else {
		if num > (*m.minH)[0] {
			heap.Push(m.minH, num)
			bottle := heap.Pop(m.minH).(int)
			heap.Push(m.maxH, bottle)
		} else {
			heap.Push(m.maxH, num)
		}
	}
}

// FindMediam 输出中位数
func (m *MedianFinder) FindMedian() float64 {
	if m.minH.Len() == m.maxH.Len() {
		return float64((*m.maxH)[0]) / 2.0 + float64((*m.minH)[0]) / 2.0
	} else {
		return float64((*m.minH)[0])
	}
}