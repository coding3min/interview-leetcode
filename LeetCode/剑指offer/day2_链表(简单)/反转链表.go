// 题目链接：https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/

// 解题思路：本题最通俗的解法为迭代，从头到尾遍历链表，不断更改当前节点的Next域
// 我们需要事先引入一个空节点，第一次迭代时，头结点指向pre，之后不断更新
// 更改Next域前，要记录当前节点的Next域指向的节点，防止链表出现断裂
// 做链表相关题目时，一定要谨防链表断裂的情况出现
// 此题的另一个要注意的点是代码的鲁棒性
// 表头指针为 null 或者整个链表只有一个节点时，我们要保证程序依旧能正常运行。

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 关于Go语言中关于空节点的初始化需要格外注意，本题中对应pre的初始化
// 我最初的写法是 pre:=*ListNode{},报错： indirect of leetcode.ListNode{}
// ListNode{}的间接无效，不确定该如何翻译
// 若 pre 初始化为 nil，报错：use of untyped nil (solution.go)
// pre:=new(ListNode) 也是不可以的，最后输出的切片结果会多一个0元素
// 正确写法为 var pre *ListNode，此时pre为空节点，无零值，值为nil
// 好吧，这块是我对GO理解存在问题，为什么 *ListNode{}不可以，我需要再思考一下
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil{
		pnext := cur.Next
		cur.Next = pre
		pre,cur = cur,pnext
	}
	// 遍历结束后，cur指向nil节点，pre指向原先链表的尾结点
	return pre
}

// 法2：递归
// 理解：如果链表长度为2，结构为：a->b->nil 想要反转链表，可以用
// a.Next.Next=a
// a.Next=nil
// return b
// 这三行代码实现，明白这个，那递归就好理解了
// 假设链表长度大于2，当前正在处理b节点，b往后的节点已经完成反转
// 我们希望b指向a，则 a.Next.Next=a
// 若当目前处理节点为空，或其Next域为空时，返回该节点，即新链表的头结点
// 为避免命名冲突，本函数名后添加了后缀 ”_recursion“
func reverseList_recrusion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newhead := reverseList(head.Next)
	head.Next.Next = head
	// 虽然这行代码实质上只为原链表的头结点服务，但是仍然不可缺少
	// 若无下面这行代码，链表有可能会形成环
	head.Next = nil
	return newhead
}