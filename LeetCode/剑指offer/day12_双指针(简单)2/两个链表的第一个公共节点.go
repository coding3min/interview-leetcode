// 题目链接：https://leetcode.cn/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
package main

//这道题比较通俗的解法是使用两个切片保存 listA 和 listB 的所有节点，然后双层循环判断节点是否相等，若相等直接返回。
//这种解法时间复杂度 O(mn)，空间复杂度（max(m,n))，m和n分别为两链表长度，和题目要求的复杂度相差甚远...

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	recordA := []*ListNode{}
	cur := headA
	for cur != nil{
		recordA = append(recordA,cur)
		cur = cur.Next
	}
	recordB := []*ListNode{}
	cur = headB
	for cur != nil{
		recordB = append(recordB,cur)
		cur = cur.Next
	}
	for i:=0;i<len(recordA);i++{
		for j:=0;j<len(recordB);j++{
			if recordA[i] == recordB[j]{
				return recordA[i]
			}
		}
	}
	return nil
}


//解法2：双指针
//这种解法非常巧妙，感觉第一次见这道题的话，很难想到
//
//声明链表指针类型变量 x 和 y ，分别指向 A 和 B，x 和 y 先在各自链表上向后移动，移动到空节点时，
//跳到对方链表，依旧同时移动，若两链表存在公共节点，x 和 y 会指向同一节点。存在两种情况
//- 若该节点非空，为两链表的第一个公共节点，返回该节点即可；
//- 该节点为空，说明两指针均走完了两个链表，未发现公共链表，返回 nil。
//
//下面来解释一下：假设链表 A 和 B 的长度分别为 m 和 n，若存在公共节点，设公共链表的长度为 x，将两个链表合并，长度为 m+n，
//若存在公共节点，即 x>0，那 x 和 y 走 m+n-x 步后到达该节点，若不存在公共节点，x 和 y 始终不相等，知道 x 和 y 共同走向合并链表的尽头，
//也就是空节点，此时，返回 A or B 都是空节点。
//
//针对细节，在做一些描述（也是自己曾经的疑惑）：
//1. 对一个链表（长度为n），从头结点走到尾结点需要 n-1 步，走 n 步会到达尾结点的Next域（即空节点），在本题的双指针解法中，真是要走到该空节点位置，然后再走到另一条链表的道路，和另一个指针一起判断是否存在公共节点。
//2. 公共节点，该节点不仅Val域相同，Next域也是相同的，所以以该节点作为头结点的链表长度也是相同的。
//3. 若两链表长度相同呢，感觉都走不到对方的链表上面？确实是这样，因为这种情况下我们也不需要走到对方的链表上，具体分两种情况
//	1. 存在公共节点，两指针 x 和 y 在第一条路上就能找到公共节点
//	2. 不存在公共节点，x 和 y 走到各自链表的尾结点的 Next 域时，已经相同，返回空节点即可

// 此解法时间复杂度O(m+n)，空间复杂度O(1)，满足题目要求(题目描述是O(n)，这里我认为已经满足)
func getIntersectionNode_2(headA, headB *ListNode) *ListNode {
	x,y := headA,headB
	for x != y{
		if x != nil{
			x = x.Next
		} else {
			x = headB
		}
		if y != nil{
			y = y.Next
		} else {
			y = headA
		}
	}
	return x
}