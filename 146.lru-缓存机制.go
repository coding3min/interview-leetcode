/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (51.60%)
 * Likes:    1142
 * Dislikes: 0
 * Total Accepted:    129.6K
 * Total Submissions: 251K
 * Testcase Example:  '["DLinkedNode","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
 * 
 * 
 * 
 * 实现 DLinkedNode 类：
 * 
 * 
 * DLinkedNode(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value)
 * 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 
 * 
 * 
 * 
 * 
 * 
 * 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 输入
 * ["DLinkedNode", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * 输出
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 * 
 * 解释
 * DLinkedNode DLinkedNode = new DLinkedNode(2);
 * DLinkedNode.put(1, 1); // 缓存是 {1=1}
 * DLinkedNode.put(2, 2); // 缓存是 {1=1, 2=2}
 * DLinkedNode.get(1);    // 返回 1
 * DLinkedNode.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
 * DLinkedNode.get(2);    // 返回 -1 (未找到)
 * DLinkedNode.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
 * DLinkedNode.get(1);    // 返回 -1 (未找到)
 * DLinkedNode.get(3);    // 返回 3
 * DLinkedNode.get(4);    // 返回 4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 0 
 * 最多调用 3 * 10^4 次 get 和 put
 * 
 * 
 */

// @lc code=start
type DLinkedNode struct {
    key,value int
    // 双向链表
    prev,next *DLinkedNode
}

type LRUCache struct{
    cache map[int]*DLinkedNode
    capacity int
    head *DLinkedNode
    tail *DLinkedNode
}

func NewDLinkedNode(k,v int,prev *DLinkedNode,next *DLinkedNode) *DLinkedNode{
    return &DLinkedNode{
        key:k,
        value:v,
        prev: next,
        next: next,
    }
}


func Constructor(capacity int) LRUCache {
    l := LRUCache{
        cache: make(map[int]*DLinkedNode),
        capacity:capacity,
        head:NewDLinkedNode(0,0,nil,nil),
        tail:NewDLinkedNode(0,0,nil,nil),
    }

    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if _,ok := this.cache[key];!ok{
        return -1
    }
    node := this.cache[key]
    this.MoveToHead(node)
    return node.value
}

func (this *LRUCache) Put(key int, value int)  {
    if _,ok := this.cache[key];ok{
        node := this.cache[key]
        node.value = value
        this.MoveToHead(node)
    }else{
        node := NewDLinkedNode(key,value,nil,nil)
        this.AddToHead(node)
        this.cache[key] = node
        if len(this.cache)>this.capacity{
            node = this.RemoveLast()
            delete(this.cache,node.key)
        }
    }
}

func (this *LRUCache) AddToHead(node *DLinkedNode){
    // 新节点指向老头节点
    node.prev = this.head
    node.next = this.head.next
    // 老头节点prev指向新节点
    this.head.next.prev = node
    // 更新头指针
    this.head.next = node
}

func (this *LRUCache) MoveToHead(node *DLinkedNode){
    this.RemoveNode(node)
    this.AddToHead(node)
}

func (this *LRUCache) RemoveNode(node *DLinkedNode){
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) RemoveLast() *DLinkedNode{
    node := this.tail.prev
    this.RemoveNode(node)
    return node
}

/**
 * Your DLinkedNode object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

