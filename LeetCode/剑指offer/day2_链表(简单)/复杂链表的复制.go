//题目链接：https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/?envType=study-plan&id=lcof

// 感觉这道题题目描述地不够清晰，可以查看主站138题的题目描述
// 刚看到这个题目我是有些困惑的，遍历原链表，然后逐个生成新节点不就可以了吗，思考过后，才发现不是这样的
// 如果没有 random域，确实可以在遍历的同时创建节点，设当前遍历到的节点为 cur，对 Next 指向的节点 pnext
// 先创建Val相同，Next为空的节点next，然后将当前节点 cur 指向next，然后 cur=next，遍历完成后即可完成复制
// 但本题中 Node 节点包含了 Random 域，Random 指向的节点位置是随机的，可能该节点还未创建，无法进行指向
// 本题的难点就在于构建新链表节点的 Random 域。
// 所以，我们需要进行两次遍历，第一次创建节点，给 Val 域复制，第二次遍历给 Next 域 和 Random 域 进行指向。
// 还有一个问题，如何保存第一次遍历创建的节点呢？哈希表应该是第一个浮现在脑海中的数据结构，键为原链表节点，值为新创建的原节点对应的节点。

// 上面讲的有些啰嗦了，下面看具体实现，参考自题解：
// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/solution/jian-zhi-offer-35-fu-za-lian-biao-de-fu-zhi-ha-xi-/
// 利用哈希表的查询特点，第一次遍历构建原链表节点和新链表节对应节点的键值对映射关系，
// 第二次遍历构建新链表各节点的 Next 与 Random 指向即可。
//1. 若头结点head为空，直接返回 head
//2. 构建哈希表 record
//3. 复制链表，进行第一次遍历，构建新节点，Next和Random域均为空
//4. 第二次遍历，所有节点已创建完成，根据record进行Next和Random域的指向
//5. 返回新链表的头结点record[head]
package main


//Definition for a Node.
type Node struct {
    Val int
    Next *Node
    Random *Node
}


func copyRandomList(head *Node) *Node {
	if head == nil{
		return head
	}
	record := map[*Node]*Node{}
	cur := head
	for cur != nil{
		node := Node{cur.Val,nil,nil}
		record[cur] = &node
		cur = cur.Next
	}
	cur = head
	for cur != nil{
		if cur.Next != nil{
			record[cur].Next = record[cur.Next]
		}
		if cur.Random != nil{
			record[cur].Random = record[cur.Random]
		}
		cur = cur.Next
	}
	return record[head]
}


// 上述解法为哈希表，是一种通俗的解法，下面看一种非常巧妙的解法
// 法2：拼接+拆分
// 这里建议大家多画图去理解这种解法
// 考虑构建 原节点 1 -> 新节点 1 -> 原节点 2 -> 新节点 2 -> …… 的拼接链表
// 如此便可在访问原节点的 random 指向节点的同时找到新对应新节点的 random 指向节点。
//算法流程：
//1. 复制各节点，构建拼接链表：设原链表为 a->b->...，构建的拼接链表为：a->a'->b->b'->...
//2. 构建新链表各节点的random指向：当访问原节点cur的随机指向节点cur.random时，对应新节点cur.next的随机指向节点为cur.random.next 。
//3. 拆分原/新链表：设置pre/cur分别指向原/新链表头节点，遍历执行pre.next=pre.next.next和cur.next=cur.next.next将两链表拆分开。
//4. 返回新链表的头节点 res 即可。

// 为避免命名冲突，此函数名后添加了后缀 ”_2“
func copyRandomList_2(head *Node) *Node {
	if head == nil{
		return head
	}
	cur := head
	for cur != nil{
		temp := &Node{cur.Val,nil,nil}
		temp.Next = cur.Next
		cur.Next = temp
		cur = temp.Next
		// 同 cur = cur.Next.Next
	}
	cur = head
	for cur != nil{
		if cur.Random != nil{
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur,res := head.Next,head.Next
	pre := head
	for cur.Next != nil{
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	// 处理原链表尾结点
	pre.Next = nil
	// 返回新链表头结点
	return res
}