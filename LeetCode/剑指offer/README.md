# LeetCode-剑指offer

技术面试环节，面试官关注应聘者的 5 种素质：

- 扎实的基础知识
- 能写高质量的代码
- 分析问题时思路清晰
- 能优化时间效率和空间效率
- 学习沟通等各方面的能力

应聘者在面试之前需要做足准备，对编程语言、数据结构和算法等基础知识有全面的了解。面试的时候如果遇到简单的问题，应聘者一定要注重细节，写出完整、鲁棒的代码。如果遇到复杂的问题，应聘者可以通过画图、举具体例子分析和分解复杂问题等方法先理清思路再动手编程。除此之外，应聘者还应该不断优化时间效率和空间效率，力求找到最优的解法。在面试过程中，应聘者还应该主动提问，以弄清楚题目的要求，表现自己的沟通能力。当面试官前后问的两个问题有相关性的时候，尽量把解决前面问题的思路迁移到后面的问题中去，展示自己良好的学习能力。如果能做到这么几点，那么通过面试获得心仪的职位将是水到渠成的事情。

一、扎实的基础知识

数据结构通常是编程面试过程中考察的重点。应聘者应熟练掌握链表、树、栈、队列、哈希表等数据结构 以及 它们的操作。面试官很喜欢问链表和二叉树的问题，这方面问题看似比较简单，但想要真正掌握也不容易，特别适合在这么短的面试时间内检验应聘者的基本功。

我们应该事先对链表的插入和删除节点了如指掌，对二叉树的各种遍历方法的循环和递归都烂熟于胸。另外，查找、排序算法也应格外注意。重点掌握二分查找、归并排序 和 快速排序，很多面试题都是这类算法的变体。

比如，旋转数组的最小数字 和 数字在排序树组中出现的次数 本质是考察二分查找

数组中的逆序对 实际上是考察归并排序

二、高质量的代码

面试官会格外关注边界条件、特殊输入等看似细枝末节但实质至关重要的地方，以考察应聘者是否注重代码质量。

比如，把一个字符串转换成整数，面试官至少会期待应聘者能够在不需要提示的情况下，考虑到输入的字符串中有非数字字符和正负号，要考虑到最大的正整数和最小的负整数以及溢出。同时面试官还期待应聘者能够考虑到当输入的字符串不能转换成整数时，应该如何做错误处理。当把这个问题的方方面面都考虑到的时候，我们就不会再认为这道题简单了。

除了问题考虑不全面之外，还有一个面试官不能容忍的错误就是程序不够鲁棒。以前面的那段代码为例，只要输入一个空指针，程序立即崩溃。这样的代码如果加入到软件当中，将是灾难。因此当面试官看到代码中对空指针没有判断并加以特殊处理的时候，通常他连往下看的兴趣都没有。

比如，求链表的倒数第 k 个节点，我们应该首先判断输入指针是否为空并做特殊处理，之后，当链表节点总数小于 k 时，是否有对应的处理方案。

三、清晰的思路

有些时候面试官会有意出一些比较复杂的问题，以考查应聘者能否在短时间内形成清晰的思路并解决问题。对于确实很复杂的问题，面试官甚至不期待应聘者能在面试不到一个小时的时间里给出完整的答案，他更看重的可能还是应聘者是否有清晰的思路。面试官通常不喜欢应聘者在没有形成清晰思路之前就草率地开始写代码，这样写出来的代码容易逻辑混乱、错误百出。

应聘者可以用几个简单的方法帮助自己形成清晰的思路。首先是举几个简单的具体例子让自己理解问题。当我们一眼看不出问题中隐藏的规律的时候，可以试着用一两个具体的例子模拟操作的过程，这样说不定就能通过具体的例子找到抽象的规律。其次可以试着用图形表示抽象的数据结构。像分析与链表、二叉树相关的题目，我们都可以画出它们的结构来简化题目。最后可以试着把复杂的问题分解成若干个简单的子问题，再一一解决。很多基于递归的思路，包括分治法和动态规划，都是把复杂的问题分解成一个或者多个简单的子问题。

比如：二叉搜索树转换为双向链表，遇到这个问题，我们不妨先画出一两个具体的二叉搜索树，直观地感受二叉搜索树和排序的双向链表有哪些联系。如果一下子找不出转换的规律，我们可以把整个二叉树看成 3 个部分：根结点、左子树和右子树。当我们递归地把转换左右子树这两个子问题解决之后，再把转换左右子树得到的链表和根结点链接起来，整个问题也就解决了

如果在面试的时候遇到难题，我们有 3 种办法分析、解决复杂的问题：画图能使抽象问题形象化，举例使抽象问题具体化，分解使复杂问题简单化。

四、优化效率的能力

优秀的程序员对时间和内存的消耗锱铢必较，他们很有激情地不断优化自己的代码。当面试官出的题目有多种解法的时候，通常他会期待应聘者最终能够找到最优解。当面试官提示还有更好的解法的时候，应聘者不能放弃思考，而应该努力寻找在时间消耗或者空间消耗上可以优化的地方。

比如，自顶向下的递归和自底向上的递归 效率天差地别，哈希表 是 典型的用空间换时间的策略，看到输入数组是排序，首先想到二分查找

五、优秀的综合能力

应聘者除了展示自己的编程能力和技术功底之外，还需要展示自己的软技能，诸如自己的沟通能力和学习能力。在面试过程中，面试官会观察应聘者在介绍项目经验或者算法思路时是否观点明确、逻辑清晰，并以此判断其沟通能力的强弱。另外，面试官也会从应聘者说话的神态和语气来判断他是否有团队合作的意识。通常面试官不会喜欢高傲或者轻视合作者的人。

知识迁移能力是一种特殊的学习能力。如果我们能够把已经掌握的知识迁移到其他领域，那么学习新技术或者解决新问题就会变得容易。面试官经常会先问一个简单的问题，再问一个很复杂但和前面的简单问题相关的问题。这个时候面试官期待应聘者能够从简单问题中得到启示，从而找到解决复杂问题的窍门。比如面试官先要求应聘者写一个函数求斐波那契数列，再问一个青蛙跳台阶的问题：一只青蛙一次可以跳上 1 级台阶，也可以跳上2 级台阶。请问这只青蛙跳上n 级台阶总共有多少种跳法。应聘者如果具有较强的知识迁移能力，就能分析出青蛙跳台阶问题实质上只是斐波那契数列的一个应用。

还有不少面试官喜欢考查应聘者的抽象建模能力和发散思维能力。面试官从日常生活中提炼出问题，比如面试题44“扑克牌中的顺子”，考查应聘者能不能把问题抽象出来用合理的数据结构表示，并找到其中的规律解决这个问题。面试官也可以限制应聘者不得使用常规方法，这要求应聘者具备创新精神，能够打开思路从多角度去分析、解决问题。比如在面试题47“不用加减乘除做加法”中，面试官期待应聘者能够打开思路，用位运算实现整数的加法。

LeetCode 剑指offer 系列题目，31天的主题分别为：

栈与队列、链表、字符串、查找算法、搜索与回溯算法、动态规划、双指针、排序、分治算法、位运算、数学、模拟。部分主题包含简单、中等和困难部分。

## day1：栈与队列（简单）

### 09.用两个栈实现队列

> 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

很经典的一道题目，两个栈实现队列。思路：两个栈，一个用来入栈（支持插入操作），一个用来出栈（支持删除操作），
栈的特点是先进后出，队列是先进先出，不考虑中间的出入，那一批数据通过栈和队列后的顺序的刚好相反的。我们需要将栈的先进后出再倒一次顺序，就和队列的顺序相同了。
官方描述：根据栈先进后出的特性，我们每次往第一个栈里插入元素后，栈顶是最后插入的元素，栈底是下一个待删除的元素。
为了维护队列先进先出的特性，我们引入第二个栈，用第二个栈维护待删除的元素，在执行删除操作的时候我们首先看下第二个栈是否为空。
如果为空，我们将第一个栈里的元素一个个弹出插入到第二个栈里，
这样第二个栈里元素的顺序就是待删除的元素的顺序，要执行删除操作的时候我们直接弹出第二个栈的元素返回即可。

具体实现：维护两个栈stack1和stack2，其中，stack1用来入队，stack2用于出队，初始时，两个栈均为空，插入元素时，stack1 插入元素即可，删除元素时，若stack2为空，将stack1所有元素出栈至stack2，若stack2仍然为空，返回-1，否则，从stack2出栈一个元素并返回。

```go
// stack1与stack2分别用于入队和出队
type CQueue struct {
   stack1 []int
   stack2 []int
}

// 构造函数，初始时，两个栈均为空
func Constructor() CQueue {
   return CQueue{[]int{},[]int{}}
}

// stack1用于入队
func (this *CQueue) AppendTail(value int)  {
   this.stack1 = append(this.stack1,value)
}

// stack2用于出队
func (this *CQueue) DeleteHead() int {
   // 若stack2长度为0，将stack1所有元素出栈至stack2
   if len(this.stack2) == 0{
      for len(this.stack1) > 0{
         x := this.stack1[len(this.stack1)-1]
         this.stack2 = append(this.stack2,x)
         this.stack1 = this.stack1[:len(this.stack1)-1]
      }
   }
   // stack2出栈的元素即为队首元素
   if len(this.stack2) > 0{
      res := this.stack2[len(this.stack2)-1]
      this.stack2 = this.stack2[:len(this.stack2)-1]
      return res
   }
   // 若stack2长度仍为0，说明队列为空，返回-1
   return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
```

### 30.包含min函数的栈

> 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

思路：建立辅助栈，与存储数据的栈大小相同，在向栈中存数据时，辅助栈同时存入一个数字
该辅助栈入栈元素的序列是一个非严格递增序列
如果该数据小于辅助栈栈顶元素，则辅助栈存入该数据，否则辅助栈还存入一个辅助栈的栈顶元素。
如果是存入第一个元素，辅助栈直接入栈即可，无比较操作）

```go
type MinStack struct {
   nums []int     //储存栈
   min []int      //辅助储存栈，存储最小值
}

/** initialize your data structure here. */
func Constructor() MinStack {
   return MinStack{
      []int{},
      []int{},
   }
}

// 入栈时，存储栈直接入栈
// 对辅助栈，若栈长度为0，直接入栈
// 否则，与栈顶元素进行比较，若大于栈顶元素，入栈，否则，辅助栈再次存入栈顶元素
func (this *MinStack) Push(x int)  {
   this.nums=append(this.nums,x)
   if len(this.min)==0{
      this.min=append(this.min,x)
   }else if this.min[len(this.min)-1]<x{
      this.min=append(this.min,this.min[len(this.min)-1])
   }else{
      this.min=append(this.min,x)
   }
}

// 出栈时，两栈均常规出栈即可
func (this *MinStack) Pop()  {
   this.nums=this.nums[:len(this.nums)-1]
   this.min=this.min[:len(this.min)-1]
}

// 返回存储栈栈顶元素
func (this *MinStack) Top() int {
   return this.nums[len(this.nums)-1]
}

// 求min值时，返回辅助栈栈顶元素即可
func (this *MinStack) Min() int {
   return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
```



## day2：链表（简单）

### 06.从尾到头打印链表

> 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
>
> **示例 1：**
>
> ```
> 输入：head = [1,3,2]
> 输出：[2,3,1]
> ```
>
> **限制：**
>
> ```
> 0 <= 链表长度 <= 10000
> ```

解题思路：从尾到头打印链表，最后返回整数切片，也就是说，链表头节点的值为切片最后的元素，而链表尾结点的值为切片第一个元素，与链表的遍历结合，很容易想到先进后出的解题策略
说到”先进后出“，那必然会用到栈，先用栈来解题
从头到尾遍历链表，将节点依次入栈
遍历结束后，在从栈顶逐个输出节点的值至整数切片即可

```Go
type ListNode struct {
    Val int
    Next *ListNode
}


func reversePrint(head *ListNode) []int {
   stack := []*ListNode{}
   res := []int{}
   if head == nil{
      return []int{}
   }
   for head != nil{
      stack = append(stack,head)
      head = head.Next
   }
   for len(stack) > 0{
      res = append(res,stack[len(stack)-1].Val)
      stack = stack[:len(stack)-1]
   }
   return res
}
```

法2：递归在本质上是一个栈结构，针对本题，我们也可以使用递归来实现
容易想到，要实现反过来输出链表，每当我们访问到一个节点时，要将当前节点的值放在返回切片的末尾，先递归输出其之后的节点，当访问到当前的节点为空节点时，返回空切片即可
利用系统栈，实现了从尾到头打印链表

```Go
func reversePrint(head *ListNode) []int {
   if head == nil{
      return []int{}
   }
   return append(reversePrint(head.Next),head.Val)
}
```

### 24.反转链表

> 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
>
> 示例:
>
> 输入: 1->2->3->4->5->NULL
> 输出: 5->4->3->2->1->NULL
>
>
> 限制：
>
> 0 <= 节点个数 <= 5000

解题思路：本题最通俗的解法为迭代，从头到尾遍历链表，不断更改当前节点的Next域

我们需要事先引入一个空节点，第一次迭代时，头结点指向pre，之后不断更新,更改Next域前，要记录当前节点的Next域指向的节点，防止链表出现断裂

做链表相关题目时，一定要谨防链表断裂的情况出现

此题的另一个要注意的点是代码的鲁棒性，表头指针为 null 或者整个链表只有一个节点时，我们要保证程序依旧能正常运行。

```Go
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
```

法2：递归
理解：如果链表长度为2，结构为：a->b->nil 想要反转链表，可以用
a.Next.Next=a
a.Next=nil
return b
这三行代码实现，明白这个，那递归就好理解了
假设链表长度大于2，当前正在处理b节点，b往后的节点已经完成反转
我们希望b指向a，则 a.Next.Next=a
若当目前处理节点为空，或其Next域为空时，返回该节点，即新链表的头结点

```Go
func reverseList(head *ListNode) *ListNode {
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
```

### 35.复杂链表的复制

> 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
>
> **提示：**
>
> - `-10000 <= Node.val <= 10000`
> - `Node.random` 为空（null）或指向链表中的节点。
> - 节点数目不超过 1000 。

感觉这道题题目描述地不够清晰，可以查看主站138题的题目描述
刚看到这个题目我是有些困惑的，遍历原链表，然后逐个生成新节点不就可以了吗，思考过后，才发现不是这样的，如果没有 random域，确实可以边遍历边创建节点，设当前遍历到的节点为 cur，对 Next 指向的节点 pnext，先创建Val相同，Next为空的节点next，然后将当前节点 cur 指向next，然后 cur=next，遍历完成后即可完成复制。

但本题中 Node 节点包含了 Random 域，Random 指向的节点位置是随机的，可能该节点还未创建，无法进行指向，本题的难点就在于构建新链表节点的 Random 域。所以，我们需要进行两次遍历，第一次创建节点，给 Val 域复制，第二次遍历给 Next 域 和 Random 域 进行指向。还有一个问题，如何保存第一次遍历创建的节点呢？哈希表应该是第一个浮现在脑海中的数据结构，键为原链表节点，值为新创建的原节点对应的节点。

上面讲的有些啰嗦了，下面看具体实现，参考自题解：[剑指 Offer 35. 复杂链表的复制（哈希表 / 拼接与拆分，清晰图解） - 复杂链表的复制 - 力扣（LeetCode）](https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/solution/jian-zhi-offer-35-fu-za-lian-biao-de-fu-zhi-ha-xi-/)

具体实现：

利用哈希表的查询特点，第一次遍历构建原链表节点和新链表节对应节点的键值对映射关系，第二次遍历构建新链表各节点的 Next 与 Random 指向即可。

1. 若头结点head为空，直接返回 head
2. 构建哈希表 record
3. 复制链表，进行第一次遍历，构建新节点，Next和Random域均为空
4. 第二次遍历，所有节点已创建完成，根据record进行Next和Random域的指向
5. 返回新链表的头结点record[head]

```Go
/**

 * Definition for a Node.
 * type Node struct {
 * Val int
 * Next *Node
 * Random *Node
 * }
*/

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
```

上述解法为哈希表，是一种通俗的解法，下面看一种非常巧妙的解法

法2：拼接+拆分

这里建议大家多画图去理解这种解法，考虑构建 原节点 1 -> 新节点 1 -> 原节点 2 -> 新节点 2 -> …… 的拼接链表，如此便可在访问原节点的 random 指向节点的同时找到新对应新节点的 random 指向节点。

算法流程：

1. 复制各节点，构建拼接链表：设原链表为 a->b->...，构建的拼接链表为：a->a'->b->b'->...
2. 构建新链表各节点的 random 指向：当访问原节点 cur 的随机指向节点 cur.random 时，对应新节点 cur.next 的随机指向节点为 cur.random.next 。
3. 拆分原 / 新链表：设置 pre / cur 分别指向原 / 新链表头节点，遍历执行 pre.next = pre.next.next 和 cur.next = cur.next.next 将两链表拆分开。
4. 返回新链表的头节点 res 即可。

```go
func copyRandomList(head *Node) *Node {
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
```



## day3：字符串（简单）

### 05.替换空格

> 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
>
> 示例：
>
> 输入：s = "We are happy."
> 输出："We%20are%20happy."
>
>
> 限制：
>
> - 0 <= s 的长度 <= 10000
>

先说点题外话：为什么要替换空格，在网络编程中，如果URL参数含有特殊字符，如空格、'#'等，可能导致服务端无法获得正确的参数值。我们需要将这些特殊符号转换成服务器可以识别的字符。转换规则是在百分号后跟上ASCII码的两位十六进制的表示。比如空格的ASCII码是32，即十六进制的0x20，因此空格被替换成“%20”

回到本题：字符串的替换，设返回字符串为 res，输入字符串 s 的长度为 n，首先我们要统计 s 中空格的数量，用来计算 res 的长度，每出现一次空格，字符串的长度+2，设 len(res)=m

然后准备一次遍历，使用两个指针 i 和 j 分别指向 s 和 res 当前遍历到的位置下标

指针从前往后或从后往前都可以，这里先从后往前，那么 i 和 j 分别初始化为 n-1 和 m-1，当指针 i 指向的位置对应字符为空格时，我们需要在 j 指向的位置及前两个位置插入“%20”，之后，i向前移动1格，j向前移动3格；若i对应位置不是空格，res[j]=s[i]即可，i和j各自向前移动1格。

```Go
func replaceSpace(s string) string {
	num_Space := 0
	for _,x := range s{
		if x == ' '{
			num_Space ++
		}
	}
	m := len(s) + 2 * num_Space
	res := make([]rune,m)
	j := m - 1
	for i:=len(s)-1;i>=0;i--{
		if s[i]==' '{
			res[j-2] = '%'
			res[j-1] = '2'
			res[j] = '0'
			j -= 3
		} else {
			res[j] = rune(s[i])
			j --
		}
	}
	return string(res)
}
```

若遍历方向改为从前往后，同理

```Go
func replaceSpace(s string) string {
   nums := 0
   for _,x := range s{
      if x == ' '{
         nums ++
      }
   }
   m := len(s) + 2 * nums
   res := make([]rune,m)
   j := 0
   for i:=0;i<len(s);i++{
      if s[i]==' '{
         res[j] = '%'
         res[j+1] = '2'
         res[j+2] = '0'
         j += 3
      } else {
         res[j] = rune(s[i])
         j ++
      }
   }
   return string(res)
}
```

另外，还可以皮一下，直接调用库函数api

库函数：可以直接调用替换字符串的api 

```golang
func replaceSpace(s string) string {
  res := strings.Replace(s," ","%20",-1)
  return res
}
```

最后，我们还应该思考一下从前往后遍历和从后往前遍历这两种解法的区别

设想这样一种场景，题目要求原地修改输入的参数，并且保证改参数有充足的空间保存修改后的元素，并且仅允许使用 O(1) 的空间复杂度。

这时从前向后遍历解法就不再生效了，而从后往前依然可以使用。这种从后往前遍历的解法是一类题目的解决方案，最好能牢记于心。

另外，由于字符串本身是不可修改的，所以这里谈到的输入参数不会再是字符串，可以为 rune切片，int切片等。

### 58.左旋转字符串

> 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
>
> 示例 1：
>
> 输入: s = "abcdefg", k = 2
> 输出: "cdefgab"
> 示例 2：
>
> 输入: s = "lrloseumgh", k = 6
> 输出: "umghlrlose"
>
>
> 限制：
>
> 1 <= k < s.length <= 10000

此题解法较多，请耐心观看

方法1：也是最通俗的解法，字符串切片进行拼接

```Go
func reverseLeftWords(s string, n int) string {
   n = n % len(s)
   return s[n:] + s[:n]
}
```

如果面试不允许对字符串进行切片，那我们可以对列表进行遍历，然后进行拼接，记为方法2：

先遍历从下标n到末尾的字符串元素，加入返回的rune切片res，然后遍历下标0-n，添加至res。最后返回将res转换为string返回即可

```Go
// 遍历字符串进行拼接
func reverseLeftWords(s string, n int) string {
   res := make([]rune,len(s))
   n = n % len(s)
   k := 0
   for i:=n;i<len(s);i++{
      res[k] = rune(s[i])
      k ++
   }
   for i:=0;i<n;i++{
      res[k] = rune(s[i])
      k++
   }
   return string(res)
}
```

方法三：还是遍历字符串，一次遍历完成

主要是通过取余操作对方法2的第二次遍历进行优化，感觉算是一种写法技巧，并不能提升算法性能。

```Go
func reverseLeftWords(s string, n int) string {
   res := make([]rune,len(s))
   n = n % len(s)
   k := 0
   for i:=n;i<len(s)+n;i++{
      res[k] = rune(s[i%len(s)])
      k ++
   }
   return string(res)
}
```

**方法4：两次翻转字符串**

前三种解法本质上是完全一样的，都是拆解字符串然后进行拼接。

本解法来自《剑指offer》，做本题前建议先做 剑指offer58题：翻转单词顺序

以“abcdefg”为例，对该字符串进行两次翻转，仍为原字符串，如果想左移2位，我们先将其分为两部分，“ab”为第一部分，“cdefg”为第二部分，先不管每一部分内部的字符，对整个字符串进行翻转得到 “gfedcba”，第一部分到了右侧，第二部分到了左侧，得到了大体上的相对顺序，由于进行了一次翻转，每部分字符串内部字符顺序是颠倒的，题目要求的每一部分相对顺序不变，那我们再对两部分字符串，分别进行一次翻转即可。

```Go
// 方法4：两次翻转字符串
func reverseLeftWords(s string, n int) string {
   n = n % len(s)
   x := reverseString(s)
   res := reverseString(x[:len(s)-n]) + reverseString(x[len(s)-n:])
   return res
}

func reverseString(s string) string{
   res := []rune(s)
   left,right := 0,len(s)-1
   for left < right{
      res[left],res[right] = res[right],res[left]
      left ++
      right --
   }
   return string(res)
}
```

很久以前就看到这个解法了，但我一直有个疑惑，为什么两次翻转对应的是左旋，而不是右旋呢？翻转和左右又没有关系，后面手写试了几次，发现了规律，第一次整体翻转后

若要左移，原字符串左边的 k 位到了右边 k 位，我们要做的是翻转 s[:n-k] 和 s[n-k:]

若要右移，原字符串右边的 k 位到了左边 k 位，我们要做的是翻转 s[:k] 和 s[k:]

如果要右移字符串，那就是如下代码：

```Go
// 右移字符串
func reverseRightWords(s string, n int) string {
   n = n % len(s)
   x := reverseString(s)
   res := reverseString(x[:n]) + reverseString(x[n:])
   return res
}
```

这下就搞清楚了！另外，左移1位，也等价于右移 len(s)-1位。

## day4.查找算法（简单）

### 03.数组中重复的数组

> 找出数组中重复的数字。在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
>
> 示例：
>
> 输入：[2, 3, 1, 0, 2, 5, 3]
> 输出：2 或 3 
>
> 限制：2<=n<=100000

解题思路：本题最通俗的解法为哈希表，用哈希表记录nums数组中元素是否出现过，若出现过，返回该元素即可。

```Go
func findRepeatNumber(nums []int) int {
   record := map[int]int{}
   for _,num := range nums{
      if record[num] > 0{
         return num
      }
      record[num] ++
   }
   return -1
}
```

时间与空间复杂度均为 O(n)，该方法没有用到题目给的数字范围条件，所以应该思考下如何做改进

试想下如何降低复杂度，对长度为n的数组，其中所有数字都在0~n-1范围内，而0~n-1又是长度为n的数组的所有下标索引。我们刚才方法是将其值放入哈希表，而现在我们给定数组的长度为n，0~n-1刚好可以对应数组索引，让我们有一种将数组设计为哈希表的思路：

对数组进行遍历，设当前遍历到的的数字值为 x，则将索引为 x 对应的数字打标记，若之后遍历到某元素后，打标机过程中发现该元素已打标记，说明该数字已出现过，返回即可。

现在要做的就是设计标记，标记为取负值，流程如下：

一次遍历数组中元素，设当前元素值为 num，它可能已经被打了标记，原本对应的数字为 x 的绝对值，先取其原本的值，设为 x，如果 nums[x] < 0，说明 x 已经出现过，返回 x 即可，否则，打标记，nums[x] = -nums[x]

但还存在一个问题，就是0取负还是0，我没想到太好的解决方案，就用了最朴素的方法，单独处理0。

处理流程如下：初始化0的下标为 zero_index，遍历数组，得到0的下标zero_index，然后第二次遍历数组，若数组元素值为 zero_index的元素个数大于1，说明该元素重复，返回0的下标即可。

这样的解题思路来自LeetCode41题：缺失的第一个正数

算法核心在于将输入数组设计为哈希表，

在进行此操作前，请务必与面试官进行交流，询问是否可以修改输入数组，确定可以的话，再用此方案。

```Go
func findRepeatNumber(nums []int) int {
   // 先处理0的情况
   // m为nums中0的下标出现的次数，zero_index为0的下标初始值
   m := 0
   zero_index := -1
   // 第一次遍历得到0的下标
   for i:=0;i<len(nums);i++{
      if nums[i] == 0{
         zero_index = i
         break
      }
   }
   // 第二次遍历判断0的下标是否出现两次及两次以上
   for i:=0;i<len(nums);i++{
      if nums[i]==zero_index{
         m ++
      }
   }
   if m > 1{
      return zero_index
   }
   // 将输入数组设计为哈希表
   for _,num := range nums{
      x := abs(num)
      // nums[x]<0，说明x元素此前出现过
      if nums[x] < 0{
         return x
      }
      // 若未出现过，打标记
      nums[x] = -nums[x]
   }
   return -1
}
```

方法3：原地交换 次题解方法来自：[剑指 Offer 03. 数组中重复的数字（哈希表 / 原地交换，清晰图解） - 数组中重复的数字 - 力扣（LeetCode）](https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/mian-shi-ti-03-shu-zu-zhong-zhong-fu-de-shu-zi-yua/)

长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内 。 此说明含义：数组元素的 索引 和 值 是 一对多 的关系。
因此，可遍历数组并通过交换操作，使元素的 索引 与 值 一一对应（即 nums[i] = i）。因而，就能通过索引映射对应的值，起到与字典等价的作用。

遍历中，第一次遇到数字 x 时，将其交换至索引 x 处；而当第二次遇到数字 x 时，一定有 nums[x] = x，此时即可得到一组重复数字。。

算法流程：
遍历数组 nums，设索引初始值为 i = 0 :

若 nums[i] = i： 说明此数字已在对应索引位置，无需交换，因此跳过；
若 nums[nums[i]] = nums[i]： 代表索引nums[i]处和索引i处的元素值都 nums[i]，即找到一组重复值，返回此值 nums[i]；否则， 交换索引为i和nums[i]的元素值，将此数字交换至对应索引位置。

若遍历完毕尚未返回，则返回 -1 。

```Go
func findRepeatNumber(nums []int) int {
   i := 0
   for i < len(nums){
      // 交换 nums[i]至索引i出后才进行下一索引处的交换
      if nums[i] == i{
         i ++
         continue
      }
      if nums[nums[i]] == nums[i]{
         return nums[i]
      }
      nums[nums[i]],nums[i] = nums[i],nums[nums[i]]
   }
   return -1
}
```

### 53.在排序树组中查找数字 I

> 统计一个数字在排序数组中出现的次数。

本题最通俗的解法是遍历数组，统计该数字出现的次数，但这样做的话，没有用到题目中数组是排序的这个条件，显然，本题想要考察二分查找算法！（看到有序数组就应该第一时间想到二分查找）

```Go
//最通俗的解法，遍历数组，统计target出现的次数
func search(nums []int, target int) int {
   n := len(nums)
   ans := 0
   for i := 0;i < n;i++{
      if nums[i] == target{
         ans += 1
      }
   }
   return ans
}
```

方法2：两次二分查找，属于二分查找算法的变体

统计一个数字在排序数组中出现的次数，那我们只需要知道其在数组中第一次和最后一次出现的下标，设为 left 和 right，出现次数即为：right-left+1

函数 first_equal_search 和 last_equal_search 分别用于查找数字第一次和最后一次出现的位置下标。

实现上述两个函数时，网上有很多花里花哨的写法，如果去死记硬背，没过几天就会全部忘光，下面我的写法非常好理解。

```Go
func search_2(nums []int, target int) int {
   left := first_equal_search(nums,target)
   right := last_equal_search(nums,target)
   if left == -1{
      return 0
   }
   return right-left+1
}

func first_equal_search(nums []int,target int) int{
   n := len(nums)
   left,right := 0,n-1
   for left <= right{
      mid := left + (right-left)>>1
      if mid==0 && nums[mid]==target{
         return mid
      } else if mid>0 && nums[mid]==target && nums[mid-1]!=target{
         return mid
      } else if nums[mid] < target{
         left = mid + 1
      } else if nums[mid] > target{
         right = mid - 1
      } else {
         right = mid -1
      }
   }
   return -1
}

func last_equal_search(nums []int,target int) int{
   n := len(nums)
   left,right := 0,n-1
   for left <= right{
      mid := left + (right-left)>>1
      if mid==n-1 && nums[mid]==target{
         return mid
      } else if mid<n-1 && nums[mid]==target && nums[mid+1]!=target{
         return mid
      } else if nums[mid] < target{
         left = mid + 1
      } else if nums[mid] > target{
         right = mid - 1
      } else {
         left = mid + 1
      }
   }
   return -1
}
```

时间复杂度：O(logn)，空间复杂度：O(1)

### 53 II. 0~n-1中缺失的数字

> 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
>
> **示例:**
>
> 输入: [0,1,3]
> 输出: 2

题目出现有序，优先考虑二分

此题的通俗解法为，遍历数组，当 nums[i] 不等于下标 i 时，说明从下标 i 开始，之后的元素值均为下标值+1，返回下标 i 即可。

```go
func missingNumber(nums []int) int {
   for i:=0;i<len(nums);i++{
      if i != nums[i]{
         return i
      }
   }
   return -1
}
```

解法2：二分查找

根据题意，可将有序数组 nums 分为两部分，第一部分其值 nums[i] = i，第二部分 nums[i]=i+1

并且第二部分元素数量限制均为 大于等于0 且 小于等于 n-1（此n为函数形参n，非数组长度)。先考虑两种特殊情况：

- 当第一部分长度为0，返回0
- 第二部分长度为0，返回数组长度即可

在常规情况下，我们要找的是第二部分第一个元素的下标

```Go
func missingNumber(nums []int) int {
   n := len(nums)
   left,right := 0,n-1
   for left <= right{
      mid := left+(right-left)>>1
      //优先考虑两种极端情况
      if mid==0 && nums[mid]!=mid{
         return mid
      } else if mid==n-1 && nums[mid]==mid{
         return mid+1
      } else if nums[mid]!=mid && nums[mid-1]==mid-1{
         return mid
      } else if nums[mid] != mid{
         right = mid - 1
      } else {
         left = mid + 1
      }
   }
   return -1
}
```

## day5.查找算法（中等）

### 04.二维数组中的查找

> 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
>
> 示例:
>
> 现有矩阵 matrix 如下：
>
> [
>   [1,   4,  7, 11, 15],
>   [2,   5,  8, 12, 19],
>   [3,   6,  9, 16, 22],
>   [10, 13, 14, 17, 24],
>   [18, 21, 23, 26, 30]
> ]
> 给定 target = 5，返回 true。
>
> 给定 target = 20，返回 false。
>
> 限制：
>
> 0 <= n <= 1000
>
> 0 <= m <= 1000

通俗解法就是遍历二维数组的每个元素，查看是否存在值等于target的元素，若存在，返回 true，否则，false。但这完全没用到题目所给的递增条件，直接pass，就不附代码。 

**当我们需要解决一个复杂的问题时，一个很有效的办法就是从一个具体的问题入手，通过分析简单具体的例子，试图寻找普遍的规律。**

利用下题目所给条件，每一行从左到右递增，每一列从上到下递增。这道题，找一下规律。如果我们先选取数组右上角的数字，则该数字左侧元素小于其值，下方元素大于其值，类似二叉搜索树。如果 target > 右上角数字的数值，则剔除第一行，第二行最后一个元素作为最右上角的数字，同样的做法，每次移动，消除一行或者一列元素，直到找到目标值 target，或者 查找范围为空（即 target 不存在于二维数组中）。

```Go
func findNumberIn2DArray(matrix [][]int, target int) bool {
   if len(matrix)==0 || len(matrix[0]) == 0{
      return false
   }
   m,n := len(matrix),len(matrix[0])
   i,j := 0,n-1
   for matrix[i][j] != target{
      if matrix[i][j] > target{
         j --
      } else {
         i ++
      }
      if i>=m || j <0{
         return false
      }
   }
   return true
}
```

另外，本题从左下角分析也是可以的，同理。

```Go
func findNumberIn2DArray_2(matrix [][]int, target int) bool {
   if len(matrix)==0 || len(matrix[0]) == 0{
      return false
   }
   m,n := len(matrix),len(matrix[0])
   i,j := m-1,0
   for matrix[i][j] != target{
      if matrix[i][j] > target{
         i --
      } else {
         j ++
      }
      if i<0 || j >= n{
         return false
      }
   }
   return true
}
```

**本题考点：考察应聘者对二维数组的理解及编程能力，还考察应聘者分析问题的能力，当发现问题比较复杂时，是否能通过具体的例子找出其中的规律，是能够解决这个问题的关键所在。本题只要从一个具体的二维数组的右上角开始分析，就能找到规律所在，从而找到解决问题的突破口。**

### 11.旋转数组的最小数字

> 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
>
> 给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。请返回旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，该数组的最小值为 1。  
>
> 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
>
> 示例 1：
>
> 输入：numbers = [3,4,5,1,2]
> 输出：1
> 示例 2：
>
> 输入：numbers = [2,2,2,0,1]
> 输出：0
>
>
> 提示：
>
> - n == numbers.length
> - 1 <= n <= 5000
> - -5000 <= numbers[i] <= 5000
> - numbers 原来是一个升序排序的数组，并进行了 1 至 n 次旋转

尽管题目描述看起来并不复杂，但本题是一道实实在在的困难题，我个人在这道题上花费时间累计超过5h，现在有时候还是会绕进去。

二分思想很简单，细节是魔鬼！各种边界处理问题
Although the basic idea of binary search is comparatively straightforward, the details can be surprisingly tricky...            ----Knuth

解题思路：二分查找，参照题解 [旋转数组的最小数字 - 旋转数组的最小数字 - 力扣（LeetCode）](https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-by-leetcode-s/)

强烈建议观看官方题解，有图片作参考会好理解一些

左右边界 left 和 right 初始化为 0 和 len(numbers)-1，假设最小值出现在 mid 位置，我们考虑数组中的最后一个元素 x：在最小值右侧的元素，它们的值一定都小于等于 x；而在最小值左侧的元素，它们的值一定都大于等于 x。因此，我们可以根据这一条性质，通过二分查找的方法找出最小值。

将中间元素 number[mid] 和 numbers[right] 作比较，可能会出现以下三种情况：（为什么拿这两个位置的数字作比较，题解没有说，但这绝对不是简简单单就能想到的，一定是做了推理，或者经验所得）

1. numbers[mid] < numbers[right]，说明 numbers[mid] 是最小值右侧的元素（也可能就是numbers[mid])，因此我们忽略当前整个搜索区间的右半部分。right=mid
2. number[mid] > numbers[right]，说明最小值在 mid 右侧，我们忽略二分查找区间的左半部分。left = mid + 1
3. numbers[mid] = numbers[right]，在这种情况下，我们无法确定最小值在mid的左侧还是右侧，但我们能确定的是，由于它们值相同，所以无论 numbers[high]是不是最小值，其都有一个替代值 numbers[mid]，因此我们可以忽略当前搜索区间的右端点。right--

最终 left 与 right 落于同一位置，即最小值所在的索引。

在这里，谈一下个人的想法：看题解，有时候我们觉得看懂了，但关掉题解，写代码的时候就忘了该怎么做，或者隔几天就忘了，很多时候就是因为你只看到了题解部分，没有看到题解之外的，作者是如何从题目所给条件推导到题解这一步的，就比如，我第一次看这道题的官方题解时，就没有思考过为什么查找的条件是 left < right，而不是 left <= right，还有，为什么比较的是 mid 元素 和 right 元素，而不是 left 元素，这些题解都没有说，如果自己不去想清楚，那就做不到对这道题的完美把控。
当然，现在我明白了，left=right 时，我们已经找到了最小值所在的索引。
这种能力的获得不是一蹴而就的，而是在练习中逐渐获取的。我也在慢慢练习的过程中...

```Go
func minArray(numbers []int) int {
   low := 0
   high := len(numbers) - 1
   for low < high {
      pivot := low + (high - low) / 2
      if numbers[pivot] < numbers[high] {
         high = pivot
      } else if numbers[pivot] > numbers[high] {
         low = pivot + 1
      } else {
         high--
      }
   }
   return numbers[low]
}
```

### 50.第一个只出现一次的字符

> 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
>
> 示例 1:
>
> 输入：s = "abaccdeff"
> 输出：'b'
> 示例 2:
>
> 输入：s = "" 
> 输出：' '
>
>
> 限制：
>
> 0 <= s 的长度 <= 50000

解题思路：两次遍历字符串，第一次遍历的过程中，用哈希表存储每个字符出现的次数，第二次遍历字符串时，哈希表查看当前字符出现次数，若哈希表值为1，返回该字符。

```Go
func firstUniqChar(s string) byte {
   if len(s) == 0{
      return ' '
   }
   record := map[rune]int{}
   for _,x := range s{
      record[x] ++
   }
   for _,x := range s{
      if record[x] == 1{
         return byte(x)
      }
   }
   return ' '
}
```

## day6.搜索与回溯算法（简单）

2022/8/17 这三道题 Krahets 的题解写的非常好，值得反复查看！

### 32 I.从上到下打印二叉树

> 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
>
> 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
>
> 例如:
> 给定二叉树: [3,9,20,null,null,15,7],
>
> 3
>
>    / \
>   9  20
>     /  \
>    15   7
> 返回：
>
> [3,9,20,15,7]
>
>
> 提示：
>
> 节点总数 <= 1000
>

解题思路（此题解来自Krahets）：层序遍历，考察二叉树的BFS，BFS通常借助队列的先入先出特性来实现。

算法流程：

1. **特例处理**： 当树的根节点为空，则直接返回空列表 [] ；
2. **初始化**： 打印结果列表 res = [] ，包含根节点的队列 queue = [root] ；
3. **BFS 循环**： 当队列 queue 为空时跳出；
	1. 出队： 队首元素出队，记为 node；
	2. 打印： 将 node.val 添加至列表 res 尾部；
	3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
4. **返回值**： 返回打印结果列表 res 即可。

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
   if root == nil{
      return []int{}
   }
   res := make([]int,0)
   q := make([]*TreeNode,0)
   q = append(q,root)
   for len(q) != 0{
      cur := q[0]
      res = append(res,cur.Val)
      if cur.Left != nil{
         q = append(q,cur.Left)
      }
      if cur.Right != nil{
         q = append(q,cur.Right)
      }
      q = q[1:]
   }
   return res
}
```

时间复杂度：O(n)，n为二叉树节点数量，即BFS需要循环n次。
空间复杂度：O(n)，最差情况下，即当树为平衡二叉树时，最多有n/2个节点同时在队列中，使用 O(n）大小的额外空间。

### 32 II.从上到下打印二叉树II

> 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
>
> 例如:
> 给定二叉树: [3,9,20,null,null,15,7],
>
> 3
>
>    / \
>   9  20
>     /  \
>    15   7
> 返回其层次遍历结果：
>
> [
>   [3],
>   [9,20],
>   [15,7]
> ]
>
> **提示：**
>
> 1. `节点总数 <= 1000`

解题思路：本质还是考察层序遍历，我们依旧使用队列来实现。将每层遍历到的节点放在一个列表中。首先，获取当前 q 的长度，即当前层节点的数目 n，然后依次访问队列 q 的前 n 个节点，两步操作：

1. 将节点 val 添加到当前层的切片中；
2. 若当前节点存在孩子节点，按顺序将其添加至列表 q

最后，q = q[n:]，删除列表中已经访问过的当前层节点，若 q 非空，下一次循环进入下一层节点的处理。

算法流程：

1. 特例处理： 当根节点为空，则返回空切片 []int{} ；
2. 初始化： 打印结果列表 res:=[]int{} ，包含根节点的队列 queue:=[]*TreeNode{root} ；
3. BFS 循环： 当队列 queue 为空时跳出；
	1. 新建一个临时列表 tmp ，用于存储当前层打印结果；
	2. 当前层打印循环： 循环次数为当前层节点数（即队列 queue 长度）；
		1. 出队： 队首元素出队，记为 node；
		2. 打印： 将 node.val 添加至 tmp 尾部；
		3. 添加子节点： 若 node 的左（右）子节点不为空，则将左（右）子节点加入队列 queue ；
	3. 将当前层结果 tmp 添加入 res 。
4. 返回值： 返回打印结果列表 res 即可。

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
   if root == nil{
      return [][]int{}
   }
   res := [][]int{}
   q := []*TreeNode{root}
   for len(q) != 0{
      n := len(q)
      temp := []int{}
      for i:=0;i<n;i++{
         cur := q[i]
         temp = append(temp,cur.Val)
         if cur.Left != nil{
            q = append(q,cur.Left)
         }
         if cur.Right != nil{
            q = append(q,cur.Right)
         }
      }
      res = append(res,temp)
      q = q[n:]
   }
   return res
}
```

### 32 III.从上到下打印二叉树III

> 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
>
> 例如:
> 给定二叉树: [3,9,20,null,null,15,7],
>
> 3
>
>    / \
>   9  20
>     /  \
>    15   7
> 返回其层次遍历结果：
>
> [
>   [3],
>   [20,9],
>   [15,7]
> ]
>
>
> 提示：
>
> 节点总数 <= 1000
>

在第二版的基础上，加入了之字形的条件，解题思路：加入变量 order，初始化为 0，表示遍历方向，order % 2 为偶数时，从左至右，否则，从右至左。这只代表了当前层节点遍历的顺序，注意！添加下一层节点时始终是从左至右的顺序。第一次提交错误，就是因为添加下一层节点时，顺序搞错了。

这里加一下 Krahets 的题解，他针对这个输出的顺序给出了三种方法：

1. 层序遍历+双端队列
2. 层序遍历+双端队列（奇偶层逻辑分离）
3. 层序遍历+倒序

我用的是第三种，该方法优点是只用数组即可，无需其他数据结构

golang：

```golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
   if root == nil{
      return [][]int{}
   }
   res := [][]int{}
   q := []*TreeNode{root}
   order := 0
   for len(q) != 0{
      n := len(q)
      temp := []int{}
      for i:=0;i<n;i++{
         if q[i].Left != nil{
            q = append(q,q[i].Left)
         }
         if q[i].Right != nil{
            q = append(q,q[i].Right)
         }
      }
      if order % 2 == 0{
         for i:=0;i<n;i++{
            temp = append(temp,q[i].Val)
         }
      } else {
         for i:=n-1;i>=0;i--{
            temp = append(temp,q[i].Val)
         }
      }
      order ++
      q = q[n:]
      res = append(res,temp)
   }
   return res
}
```

然后看了下官方的操作，挺不错的，先按照第二版从上到下打印二叉树的思路解题，最后向返回数组添加某一层遍历结果时，先将该层遍历结果进行翻转，简单易懂。

## day7.搜索与回溯算法（简单）

### 26.树的子结构

> 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
>
> B是A的子结构， 即 A中有出现和B相同的结构和节点值。
>
> 例如:
> 给定的树 A:
>
>  3
> / \
>
>    4   5
>   / \
>  1   2
> 给定的树 B：
>
>    4 
>   /
>  1
> 返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
>
> 示例 1：
>
> 输入：A = [1,2,3], B = [3,1]
> 输出：false
> 示例 2：
>
> 输入：A = [3,4,5,1,2], B = [4,1]
> 输出：true
> 限制：
>
> 0 <= 节点个数 <= 10000
>

此题解参考自：[面试题26. 树的子结构（先序遍历 + 包含判断，清晰图解） - 树的子结构 - 力扣（LeetCode）](https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/solution/mian-shi-ti-26-shu-de-zi-jie-gou-xian-xu-bian-li-p/)

若树 B 是 A 的子结构，则 A 子结构的根节点可能是 A 中任意一个节点，我们判断 B 是否为 A 的子结构，就需要遍历 A 的所有节点，然后判断是否有某节点 以 该节点为 根节点的子树包含 B（有点绕，其实就是，若 A 与 B 根节点相同，B 是否为 A 的子结构，但当 B 为空的时候，B 必为 A 的子结构）。

名词规定：树 A 的根节点记作 节点 A ，树 B 的根节点称为 节点 B 。

recur(A, B) 函数，用于判断 树 A 中以 A为根节点的子结构是否包含树 B

终止条件：

- 当节点 B 为空：说明树 B 已匹配完成（越过叶子节点），因此返回 true ；
- 当节点 A 为空：说明已经越过树 A 叶子节点（而 B 节点非空），即匹配失败，返回 false ；
- 当节点 A 和 B 的值不同：说明匹配失败，返回 false ；

返回值：

- 判断 A 和 B 的左子节点是否相等，即 recur(A.left, B.left) ；
- 判断 A 和 B 的右子节点是否相等，即 recur(A.right, B.right) ；

两者取逻辑与后返回。

特例处理： 当 树 A 为空 或 树 B 为空时，直接返回 false（对应题目中的约定：空树不是任意一个树的子结构） 。

之后我们对 A 作遍历（前中后序都可以），对其中每个节点与 B 进行 recur 判断，若存在 true 结果，返回最终的 true；若全为 false，说明 B 不是 A 的子结构，返回 false。

### 27.二叉树的镜像

> 请完成一个函数，输入一个二叉树，该函数输出它的镜像。
>
> 例如输入：
>
>  4
>
>    /   \
>   2     7
>  / \   / \
> 1   3 6   9
> 镜像输出：
>
>  4
>
>    /   \
>   7     2
>  / \   / \
> 9   6 3   1
>
> 示例：
>
> 输入：root = [4,2,7,1,3,6,9]
> 输出：[4,7,2,9,6,3,1]
>
>
> 限制：
>
> 0 <= 节点个数 <= 1000
>

**二叉树的定义就是递归的，所以做二叉树的题目一定要有递归的想法，有思路后开始着手写代码，不要在脑海中模拟太多层的递归，否则很容易绕进去。--本题做完的一些经验**

观察例子可以发现，镜像 就是 每个节点的左右子树进行翻转。所以，两步完成本题：

1. 遍历二叉树保存所有节点（前中后序遍历都可以），用节点的指针类型切片保存；
2. 遍历切片，交换其左右子节点 node.Left,node.Right = node.Right,node.Left

```Go
func mirrorTree(root *TreeNode) *TreeNode {
   nodes := []*TreeNode{}
   var preorder func(root *TreeNode)
   preorder = func(root *TreeNode){
      if root == nil{
         return
      }
      nodes = append(nodes,root)
      if root.Left != nil{
         preorder(root.Left)
      }
      if root.Right != nil{
         preorder(root.Right)
      }
   }
   preorder(root)
   for _,node := range nodes{
      node.Left,node.Right = node.Right,node.Left
   }
   return root
}
```

上面这种解法的时间空间复杂度均为 O(n)

我们也可以在遍历的过程中交换节点的左右子节点，省去 nodes 切片占用的内存空间，但这本质上并不会降低时间和空间复杂度，因为我们在对二叉树进行遍历的过程中，调用了系统栈，系统需要使用 O(n) 大小的栈空间。

```Go
// 前序遍历的过程中交换左右子节点
func mirrorTree(root *TreeNode) *TreeNode {
   var preorder func(root *TreeNode)
   preorder = func(root *TreeNode){
      if root == nil{
         return
      }
      root.Left,root.Right = root.Right,root.Left
      if root.Left != nil{
         preorder(root.Left)
      }
      if root.Right != nil{
         preorder(root.Right)
      }
   }
   preorder(root)
   return root
}
```

另外，在LeetCode题解区我还看到了有人用层序遍历的方式解题，本质上和上面两种思路是一样的，就是选择了另一种遍历二叉树的方式，交换左右子节点的核心操作无任何变化，代码开头的判断 root 节点是否为空不可省略，因为 q 初始化默认加入root节点，若该节点为空，后续代码会出错。

```Go
func mirrorTree(root *TreeNode) *TreeNode {
   if root == nil{
      return root
   }
   q := []*TreeNode{root}
   for len(q) != 0{
      node := q[0]
      node.Left,node.Right = node.Right,node.Left
      if node.Left != nil{
         q = append(q,node.Left)
      }
      if node.Right != nil{
         q = append(q,node.Right)
      }
      q = q[1:]
   }
   return root
}
```

### 28.对称的二叉树

> 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
>
> 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
>
> 1
>
>    / \
>   2   2
>  / \ / \
> 3  4 4  3
> 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
>
> 1
>
>    / \
>   2   2
>    \   \
>    3    3
>
> 示例 1：
>
> 输入：root = [1,2,2,3,4,4,3]
> 输出：true
> 示例 2：
>
> 输入：root = [1,2,2,null,3,null,3]
> 输出：false
>
>
> 限制：
>
> 0 <= 节点个数 <= 1000

解题思路：如果二叉树只有一个根节点，那其必然是对称的，本题主要是判断节点的左右子树是否对称。

判断左右子树是否对称，若两子树根节点均为空，说明对称，若某子树为空而另一颗子树不为空，说明不对称，两子树根节点值不相等的话也必定不对称吗，否则，对称，之后递归地判断

- 左子树的左子树 与 右子树的右子树 是否对称
- 左子树的右子树 与 右子树的左子树 是否对称

两者均对称时，说明该左右子树对称。

```Go
func isSymmetric(root *TreeNode) bool {
	// 根节点为空，对称
	if root == nil{
		return true
	}
	var sym func(x,y *TreeNode) bool
	// 递归判断左右子树是否对称
	sym = func(x,y *TreeNode) bool {
		// 两者均为空说明对称
		if x == nil && y == nil{
			return true
		//	其中某一子树为空，而另一子树不为空，不对称
		} else if x == nil || y == nil{
			return false
		}
		// 两者根节点值不相等时也不对称
		if x.Val != y.Val{
			return false
		}
		// 递归判断x的左右子树 与 y 的右左子树是否对称
		return sym(x.Left,y.Right) && sym(x.Right,y.Left)
	}
	return sym(root.Left,root.Right)
}
```



## day8.动态规划（简单）

### 10-I.斐波那契数列

> 写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
>
> F(0) = 0,   F(1) = 1
> F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
> 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
>
> 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 示例：
>
> 输入：n = 2
> 输出：1
>
>
> 提示：
>
> 0 <= n <= 100

解题思路：最基础的动态规划问题，题目已经给出状态转移方程，或者可以直接理解为 递归，一个阶段只有一个状态

三步走：

1. 确定dp数组及下标含义：dp[i] 代表 F(i)，要求第 n 个斐波那契数，则 dp 数组长度为 n；

2. 数组初始化：题目已给出，dp[0]=0,dp[1]=1；
3. 状态转移方程：题目已给出，i>=2时，dp[i] = dp[i-1]+dp[i-2]。

最终返回 dp[n-1] 即可

```Go
func fib(n int) int {
   if n <= 1{
      return n
   }
   x := int(math.Pow(10,9)) + 7
   dp := make([]int,n+1)
   dp[0],dp[1] = 0,1
   for i:=2;i<n+1;i++{
      dp[i] = (dp[i-1] + dp[i-2]) % x
   }
   return dp[n]
}
```

由于每个阶段的状态只与前两个状态有关，所以我们可以用滚动数组代替 dp 数组解题，将空间复杂度从 O(n) 降低至 O(1)。

```Go
func fib(n int) int {
   if n <= 1{
      return n
   }
   x := int(math.Pow(10,9)) + 7
   a,b := 0,1
   var res int
   for i:=2;i<n+1;i++{
      res = (a+b) % x
      a,b = b,res
   }
   return res
}
```

### 10-II.青蛙跳台阶问题

> 一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
>
> 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 示例 1：
>
> 输入：n = 2
> 输出：2
> 示例 2：
>
> 输入：n = 7
> 输出：21
> 示例 3：
>
> 输入：n = 0
> 输出：1
> 提示：
>
> 0 <= n <= 100

本质上和斐波那契数列一样，区别在于 dp 数组初始化不同。

跳到第 n 级台阶，我们可以从第 n-1 个台阶跳上去，也可以从 n-2 个台阶跳上去，跳法，即为跳到 n-1 级台阶 和 跳到 n-2 级台阶的跳法之和。

动态规划三步：

1. 确定dp数组及下标含义：dp[i] 代表跳到第 i 级台阶的跳法数量，求跳到第 n 级台阶的跳法数量，dp[0]情况特殊，则 dp 数组长度设为 n+1；
2. dp 数组初始化：由题意可知，dp[0]=1,dp[1]=1,dp[2]=2
3. 状态转移方程：当 i > 2 时，dp[i] = dp[i-1] + dp[i-2]

最终，返回 dp[n] 即可。

```Go
func numWays(n int) int {
   if n <= 1{
      return 1
   }
   x := int(math.Pow(10,9)+7)
   dp := make([]int,n+1)
   dp[0] = 1
   dp[1] = 1
   for i:=2;i<=n;i++{
      dp[i] = (dp[i-1]+dp[i-2]) % x
   }
   return dp[n]
}
```

由于每个阶段的状态只与前两个状态有关，所以我们可以用滚动数组代替 dp 数组解题，将空间复杂度从 O(n) 降低至 O(1)。

```Go
func numWays(n int) int {
   if n <= 1{
      return 1
   }
   x := int(math.Pow(10,9)+7)
   a,b := 1,1
   var res int
   for i:=2;i<=n;i++{
      res = (a+b) % x
      a,b = b,res
   }
   return res
}
```

### 63.股票的最大利润

> 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
>
> 示例 1:
>
> 输入: [7,1,5,3,6,4]
> 输出: 5
> 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
>      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
> 示例 2:
>
> 输入: [7,6,4,3,1]
> 输出: 0
> 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
>
>
> 限制：
>
> 0 <= 数组长度 <= 10^5

这道题的通俗解法是双层循环（也就是暴力解法），假设共有 n 天，最大利润初始化为 0，第一层循环中设定买股票的时间点，从第 1 天遍历到 第 n-1 天，第二层循环遍历卖股票的时间点，从买股票的第二天遍历到最后一天，若买股票的价格大于买股票的价格，更新最大利润。

```Go
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

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}
```

方法2：一次遍历

遍历股票价格的过程中，维护一个股票最小价格，若当前遍历到的价格小于该最小价格，则更新最小价格，否则，更新最大利润。

```Go
func maxProfit(prices []int) int {
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
```

这道题和动态规划有关系吗？我觉得关系不大吧，用递推更加合适一些。

主要是场景过于简单，只需要维护一个最小价格，每个阶段只有一个状态，且每个阶段的状态与之前阶段均无关。



## day9.动态规划（中等）

### 42.连续子数组的最大和

> 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
>
> 要求时间复杂度为O(n)。
>
> 示例1:
>
> 输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
> 输出: 6
> 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
>
>
> 提示：
>
> 1 <= arr.length <= 10^5
> -100 <= arr[i] <= 100

解法：动态规划

动态规划三步：

1. 确定dp数组大小及下标含义：dp[i] 代表以下标 i 为最后元素的子数组的最大值，dp 数组长度与给定数组 nums 长度相同。
2. dp 数组初始化：每个单独元素都是一个子数组，初始化 dp[0] = nums[0]
3. 状态转移方程：从下标 1 开始，dp[i] = max(nums[i],nums[i]+dp[i-1])，若 dp[i-1]>0，说明以下标 i 截止的子数组的最大和要包含之前元素，包含多少呢？dp[i-1] 已经处理完了，我们只需要相加即可。

```Go
func maxSubArray(nums []int) int {
   n := len(nums)
   dp := make([]int,n)
   dp[0] = nums[0]
   res := nums[0]
   for i:=1;i<n;i++{
      dp[i] = max(nums[i],nums[i]+dp[i-1])
      res = max(res,dp[i])
   }
   return res
}

func max(x,y int) int {
   if x > y{
      return x
   }
   return y
}
```

### 47.礼物的最大价值

> 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
>
> 示例 1:
>
> 输入: 
> [
>   [1,3,1],
>   [1,5,1],
>   [4,2,1]
> ]
> 输出: 12
> 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
>
>
> 提示：
>
> 0 < grid.length <= 200
> 0 < grid[0].length <= 200

经典二维dp，在每个网格的位置有两个选择，向右 或 向下走，当前网格可获得的最大礼物价值，就只能从其上面的格子或左边的格子 的最大礼物价值 加上当前格子的礼物价值 得到，回溯到起始节点，容易想到用 dp 解题。

动态规划三步骤：

1. 确定dp数组大小及下标含义：dp\[i\]\[j] 代表 从棋盘左上角走到 (i,j) 下标位置可以获得的礼物最大价值，则 len(dp) = len(grid),len(dp[0])=len(grid[0])，即 大小与给定棋盘大小相等；
2. dp数组初始化：（一般情况下初始第一行和第一列）网格（0,0）为起始位置，dp\[i\]\[j] 没有别的选择，dp\[i\]\[j] = grid\[i\]\[j] ，因为计算当前网格的最大礼物价值，需要知道其上方和左方网格的最大礼物价值，所以我们要初始化第一行和第一列的 dp 数组元素，防止越界情况的发生；
3. 状态转移方程：i>1 且 j>1 时：dp\[i\]\[j] = max(dp\[i-1\]\[j] ,dp\[i\]\[j-1] ) + grid\[i\]\[j] 

最后返回 dp\[m-1\]\[n-1\] 即可。

```Go
func maxValue(grid [][]int) int {
   if len(grid)==0 || len(grid[0])==0{
      return 0
   }
   m,n := len(grid),len(grid[0])
   dp := make([][]int,m)
   for i:=0;i<m;i++{
      dp[i] = make([]int,n)
   }
   dp[0][0] = grid[0][0]
   for i:=1;i<m;i++{
      dp[i][0] = grid[i][0] + dp[i-1][0]
   }
   for i:=1;i<n;i++{
      dp[0][i] = grid[0][i] + dp[0][i-1]
   }
   for i:=1;i<m;i++{
      for j:=1;j<n;j++{
         dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + grid[i][j]
      }
   }
   return dp[m-1][n-1]
}

func max(x,y int) int {
   if x > y{
      return x
   }
   return y
}
```



## day10.动态规划（中等）

### 46.把数字翻译成字符串

> 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
>
> 示例 1:
>
> 输入: 12258
> 输出: 5
> 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
>
>
> 提示：
>
> 0 <= num < 2^31

比较明显的动态规划题目，设 dp[i]代表s[:i+1]的翻译方法数目，dp[i] 明显依赖于 dp[i-1] 与 dp[i-2]，类似于斐波那契数是吧，本题有点不一样的地方在于 dp[i]=dp[i-1]+dp[i-2] 是有条件的

- 当 s[i-1] 与 s[i] 组成的数字小于 25 时，s[i] 可以与 s[i-1] 组合翻译，也可以分开翻译，dp[i]=dp[i-1]+dp[i-2]，
- 否则 s[i] 与 s[i-1] 无法组合翻译，只能单独翻译，dp[i] = dp[i-1]

动态规划三步骤：

- 确定dp数组大小及下标含义：dp[i] 代表 s[:i+1] 的翻译方法数目，len(dp)=len(string(num))
- dp 数组初始化：dp[0]对应s[0]，单个字符只有一种翻译方法，dp[0]=1，当 s[:2] 小于26 且 s[i]!=0 时，dp[1]=2，否则 dp[1]=1
- 状态转移方程：从下标 2 开始遍历，x = strconv.Atoi(s[i-1:i+1])，并且判断 s[i-1] 是否为 0
	- 若 x < 26 且 s[i-1]!= 0，dp[i]=dp[i-1]+dp[i-2]，s[i] 可以与 s[i-1] 组合翻译，也可以单独翻译
	- 否则，dp[i] = dp[i-1]，s[i] 只能单独翻译，s[:i+1] 的翻译方法数目依赖于 s[:i]

最后，返回 dp[n-1] 即可。

```Go
// 忘记考虑 01 不能翻译成 b，只能翻译为 ab
func translateNum(num int) int {
   s := strconv.Itoa(num)
   n := len(string(s))
   if n == 0{
      return 0
   }
   dp := make([]int,n)
   dp[0] = 1
   x,_ := strconv.Atoi(s[:2])
   if x < 26{
      dp[1] = 2
   } else {
      dp[1] = 1
   }
   for i:=2;i<n;i++{
      x,_ = strconv.Atoi(s[i-1:i+1])
      zero,_ := strconv.Atoi(string(s[i-1]))
      if x < 26 && zero!= 0{
         dp[i] = dp[i-1] + dp[i-2]
      } else {
         dp[i] = dp[i-1]
      }
   }
   return dp[n-1]
}
```

有一点需要注意，在状态转移方程那里一定要判断 s[i-1] 是否为 0，因为 “01”只能翻译为 “ab”，不能翻译成 “a"，我第一次做这道题的时候就是忘记判断 s[i-1]是否为0，导致没有ac，刚才做的时候，又在这个地方跌到坑里了，有了再一再二，不会有再三再四了。

### 48.最长不含重复字符的子字符串

> 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
>
> 示例 1:
>
> 输入: "abcabcbb"
> 输出: 3 
> 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
> 示例 2:
>
> 输入: "bbbbb"
> 输出: 1
> 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
> 示例 3:
>
> 输入: "pwwkew"
> 输出: 3
> 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
>      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
>
>
> 提示：
>
> s.length <= 40000
>

我们可以用一个字典存储每个字符最近一次出现的下标

一次遍历字符串，变量 start 代表当前不含重复字符的子字符串的起始下标-1，初始化为 -1，为什么是 -1呢，因为我们遍历第一个字符的时候，其下标为 0，为不包含重复字符的子字符串，长度为 0 - （-1），返回变量 res 初始化为 0，

一次遍历字符串，for i,x := range s ，维护 s[start+1;i+1] 为不含重复字符的子串

若当前遍历到的字符的最近一次出现下标大于 start（不会出现等于的情况），说明当前加入当前字符后，该字符串将包含重复字符，我们的做法是，将 start 更新到其最近一次出现的下标，这样就保证了加入当前字符后，我们目前的子字符串仍然不包含重复字符，

若 小于 start 或者 该字符在字符串中是第一次出现，则当前不含重复字符的子串长度增加，更新 res 变量，

最后，更新当前字符的最近出现下标。

```go
    start := -1
   record := map[rune]int{}
   res := 0
   for i,x := range s{
      if _,ok := record[x];ok && record[x] > start{
         start = record[x]
      } else {
         res = max(res,i-start)
      }
      record[x] = i
   }
   return res
}

func max(x,y int) int {
   if x > y{
      return x
   }
   return y
}
```



## day11.双指针（简单）

### 18.删除链表的节点

> 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
>
> 返回删除后的链表的头节点。
>
> 注意：此题对比原题有改动
>
> 示例 1:
>
> 输入: head = [4,5,1,9], val = 5
> 输出: [4,1,9]
> 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
> 示例 2:
>
> 输入: head = [4,5,1,9], val = 1
> 输出: [4,5,9]
> 解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
>
>
> 说明：
>
> 题目保证链表中节点的值互不相同

思路很简单的题目，就是要注意代码的鲁棒性

链表相关的题目非常考察代码的鲁棒性，谨防访问空节点.Next 这种情况的发生。

要删除某个节点，我们需要记录要删除节点的前一个节点，将前一个节点的 Next域 指向要删除节点的下一个节点，所以我们要保存要删除节点的前一个节点，才能做到删除节点。

cur=head，一次遍历链表，遍历条件为 cur.Next != nil，如果 cur.Next 是我们要删除的节点，cur.Next = cur.Next.Next，然后结束遍历，否则 cur=cur.Next，向后移动节点。最后返回 head。

有一种情况需要额外判断，就是 head 节点就是要删除的节点，因为我们无法保存该节点的前一个节点，针对这种情况，直接返回 head.Next 即可，在遍历之前，我们就要对这种情况进行判断。

```Go
func deleteNode(head *ListNode, val int) *ListNode {
   if head.Val == val{
      return head.Next
   }
   cur := head
   for cur.Next != nil{
      if cur.Next.Val == val{
         cur.Next = cur.Next.Next
         // 跳出是为了应对要删除的节点为链表最后一个，会造成访问空指针的情况出现
         break
      }
      cur = cur.Next
   }
   return head
}
```

稍微来扩展一下，对应《剑指offer》

如果题目给的是一个链表的头指针 和 一个要被删除的节点，而非节点的值。那我们应该如何处理呢？

按照如上时间复杂度O(n)的思路当然是可以的，那是不是一定需要找到被删除节点的前一个节点呢？答案是否定的，我们可以很方便地找到要删除的节点的下一个节点。如果我们把下一个节点的内容复制到被删除节点上来覆盖该节点原有的内容，再把下一个节点删除，那是不是就相当于把需要删除的节点删除了？

（这个思路很新奇，在用户的需求中，删除节点的意思就是删除节点值，而非内存空间，要抓住这样的需求和抽象层析的信息不对称来寻求突破）

上述思路还有一个问题：如果被删除的节点位于链表的尾部，Next域为空，就无法使用了，在这种情况下，我们只能遍历链表，得到该节点的前一个节点，将该节点删除。

值得注意的是：上述代码仍然不是一段完美的代码，因为它基于这样一种假设：要删除的节点确定在链表中。我们需要O(n)的时间才能判断链表是否包含某一节点。在面试的过程中，我们可以和面试官讨论这个假设，这样面试官就会觉得我们考虑问题非常全面。

考察点：

- 应聘者对链表的编程能力
- 创新思维能力，这需要应聘者打破常规思维模式。当我们想要删除某一节点时，不一定要删除节点本身，可以先把下一个节点的内容复制出来覆盖要被删除节点的内容，再将下一个节点删除，这种思路不容易想到
- 代码的鲁棒性，也即应聘者思维的全面性，全面考虑到该节点是链表尾结点或头结点的情况。

### 22.链表中倒数第k个节点

> 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
>
> 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
>
> 示例：
>
> 给定一个链表: 1->2->3->4->5, 和 k = 2.
>
> 返回链表 4->5.

这道题比较通俗的解法是两次遍历

先一次遍历链表，得到链表的长度 n，倒数第 k 个节点，即正数第 n-k+1 个节点，第二次遍历输出该节点即可。

```Go
func getKthFromEnd(head *ListNode, k int) *ListNode {
   n := 0
   cur := head
   // 第一次遍历得到链表长度n
   for cur != nil{
      n ++
      cur = cur.Next
   }
   cur = head
   x := n-k+1
   //第二次遍历找到整数第n-k+1个节点并返回
   for i:=0;i<x-1;i++{
      cur = cur.Next
   }
   return cur
}
```

方法2：快慢指针

使用快慢指针后一次遍历即可得到倒数第 k 个节点

快慢指针 fast 和 slow 初始化指向 head 节点，fast 指针先走 k 步，此时 fast 和 slow 之间距离为 k，然后两指针同时向前走，当 fast 指向链表尾结点的Next域（即空节点）时，slow 指向倒数第 k 个节点，返回 slow 即可。

```Go
func getKthFromEnd(head *ListNode, k int) *ListNode {
   fast,slow := head,head
   for i:=0;i<k;i++{
      fast = fast.Next
   }
   for fast != nil{
      fast = fast.Next
      slow = slow.Next
   }
   return slow
}
```



## day12.双指针（简单）

### 25.合并两个排序的链表

> 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
>
> 示例1：
>
> 输入：1->2->4, 1->3->4
> 输出：1->1->2->3->4->4
> 限制：
>
> 0 <= 链表长度 <= 1000

很经典的链表题目了，计算机考研算法必做题

我的思路是，新建一个最终返回链表的头结点 head，Val域为0（随便设），Next域为空，之后的链表合并操作在其 Next域进行，最终返回 head.Next 即可。

再声明一个变量 cur = head，用于合并链表，具体分为两步：

1. 循环合并：当 l1 与 l2 链表头结点均非空时，判断两者头结点的Val域，谁更小，cur.Next 指向该节点，然后该链表头节点指向其下一个节点，cur也指向其Next域，如此循环，直至 l1 或 l2 为空；
2. 合并剩余尾部：此时若 l1 链表非空，则 head 指向的链表所有节点值均小于等于 l1 链表剩余节点值，cur.Next 指向 l1 即可；若是 l2 链表非空，同理。

最终返回 head.Next 即可。

```Go
//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
   head := &ListNode{0,nil}
   cur := head
   for l1 != nil && l2 != nil{
      if l1.Val <= l2.Val{
         cur.Next = l1
         cur = cur.Next
         l1 = l1.Next
      } else {
         cur.Next = l2
         cur = cur.Next
         l2 = l2.Next
      }
   }
   if l1 != nil{
      cur.Next = l1
   }
   if l2 != nil{
      cur.Next = l2
   }
   return head.Next
}
```

这道题还可以递归解决：

当 l1 或 l2 链表为空时，无需合并，直接返回另一个非空链表即可。

否则，判断 l1 和 l2 链表的头结点谁更小，改变其 Next 域，递归地指向 l1 和 l2.Next 合并链表的结果，最后返回该链表头结点。递归结束条件为某一链表为空。

```Go
func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
   if l1 == nil{
      return l2
   }
   if l2 == nil{
      return l1
   }
   if l1.Val <= l2.Val{
      l1.Next = mergeTwoLists(l1.Next,l2)
      return l1
   } else {
      l2.Next = mergeTwoLists(l1,l2.Next)
      return l2
   }
}
```

### 52.两个链表的第一个公共节点

> 输入两个链表，找出它们的第一个公共节点。
>
> 示例：
>
> 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
> 输出：Reference of the node with value = 8
> 输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
>
> 注意：
>
> - 如果两个链表没有交点，返回 null.
> - 在返回结果后，两个链表仍须保持原有的结构。
> - 可假定整个链表结构中没有循环。
> - 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

这道题比较通俗的解法是使用两个切片保存 listA 和 listB 的所有节点，然后双层循环判断节点是否相等，若相等直接返回。这种解法时间复杂度 O(mn)，空间复杂度（max(m,n))，m和n分别为两链表长度，和题目要求的复杂度相差甚远...

```Go
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
```

解法2：双指针

这种解法挺巧妙的，感觉第一次见这道题的话，很难想到

声明链表指针类型变量 x 和 y ，分别指向 A 和 B，x 和 y 先在各自链表上向后移动，移动到空节点时，跳到对方链表，依旧同时移动，若两链表存在公共节点，x 和 y 会指向同一节点。存在两种情况

- 若该节点非空，为两链表的第一个公共节点，返回该节点即可；
- 该节点为空，说明两指针均走完了两个链表，未发现公共链表，返回 nil。

下面来解释一下：假设链表 A 和 B 的长度分别为 m 和 n，若存在公共节点，设公共链表的长度为 x，将两个链表合并，长度为 m+n，若存在公共节点，即 x>0，那 x 和 y 走 m+n-x 步后到达该节点，若不存在公共节点，x 和 y 始终不相等，知道 x 和 y 共同走向合并链表的尽头，也就是空节点，此时，返回 A or B 都是空节点。

针对细节，在做一些描述（也是自己曾经的疑惑）：

1. 对一个链表（长度为n），从头结点走到尾结点需要 n-1 步，走 n 步会到达尾结点的Next域（即空节点），在本题的双指针解法中，真是要走到该空节点位置，然后再走到另一条链表的道路，和另一个指针一起判断是否存在公共节点。
2. 公共节点，该节点不仅Val域相同，Next域也是相同的，所以以该节点作为头结点的链表长度也是相同的。
3. 若两链表长度相同呢？两种情况
	1. 存在公共节点，两指针 x 和 y 不需要走到另一条链表上，在第一条路上就能找到公共节点
	2. 不存在公共节点，x 和 y 走到各自链表的尾结点的 Next 域时，已经相同，返回空节点即可

```Go
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
```



## day13.双指针(简单)

### 57.和为s的两个数字

> 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
>
> 示例 1：
>
> 输入：nums = [2,7,11,15], target = 9
> 输出：[2,7] 或者 [7,2]
>
>
> 限制：
>
> 1 <= nums.length <= 10^5
> 1 <= nums[i] <= 10^6

梦回LeetCode第一题：两数之和，有点像是吧，这道题和两数之和的区别在于：数组有序。

方法 1 ：先来暴力解法，双层遍历

第一层遍历从头开始，第二层遍历从第一层遍历元素的下一个元素开始，若两数之和等于 target，返回即可。

```Go
func twoSum(nums []int, target int) []int {
   n := len(nums)
   for i:=0;i<n;i++{
      for j:=i+1;j<n;j++{
         if nums[i]+nums[j] == target{
            return []int{nums[i],nums[j]}
         }
      }
   }
   return []int{0,0}
}
```

嗯...和预想的一样，超时了，时间复杂度O(n^2)，空间复杂度 O(1)

方法 2：哈希表

接下来用 两数之和 的解法，一次遍历，哈希表记录遍历过的数字

```Go
func twoSum(nums []int, target int) []int {
   record := map[int]struct{}{}
   for i:=0;i<len(nums);i++{
      if _,ok := record[target-nums[i]];ok{
         return []int{nums[i],target-nums[i]}
      }
      record[nums[i]] = struct{}{}
   }
   return []int{0,0}
}
```

时间复杂度O(n)，空间复杂度O(n)

方法三：双指针

上面这两种解题思路没有用到数组有序的这个非常重要的条件，一定会有更优解！

我们是如何使用双指针来进行优化的呢？这是对双层遍历的优化，在双层遍历中，第一层遍历，元素从左至右，第二层遍历，元素从第一个元素的下一个元素向右遍历。考虑使用数组有序的这个条件，双指针指向最小和最大的元素，若两数之和大于 target，最大元素移向第二大的元素，若小于 target，最小元素移向第二小的元素，这样两边同时向 target 靠拢，一次遍历就可以得到结果。

使用双指针，left 和 right 指针初始化指向数组的第一个和最后一个元素的下标

- 若两指针指向元素之和 大于 target，说明需要减小两数之和，右指针向左移
- 若两指针指向元素之和 小于 target，说明需要增大两数之和，左指针右移
- 若相等，将两指针指向元素组合为切片返回即可。

```Go
func twoSum(nums []int, target int) []int {
   left,right := 0,len(nums)-1
   for left < right{
      if nums[left]+nums[right] > target{
         right --
      } else if nums[left]+nums[right] < target{
         left ++
      } else {
         return []int{nums[left],nums[right]}
      }
   }
   return []int{0,0}
}
```

时间复杂度O(n)，空间复杂度O(1)，充分利用题目所给条件，为最优解法！

### 58-I.翻转单词顺序

> 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
>
> 示例 1：
>
> 输入: "the sky is blue"
> 输出: "blue is sky the"
> 示例 2：
>
> 输入: "  hello world!  "
> 输出: "world! hello"
> 解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
> 示例 3：
>
> 输入: "a good   example"
> 输出: "example good a"
> 解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
>
>
> 说明：
>
> 无空格字符构成一个单词。
> 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
> 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

我的想法是使用栈，每遇到一个完整的单词，添加到栈中，最后出栈。

具体过程：一次遍历，用 string.Builder 构建字符串，若当前字符不是空格，则写入将当前字符写入 sb，若是，则将 sb 的内容写入栈 并 清空内容，准备下一个单词的填充。

```Go
func reverseWords(s string) string {
   var sb strings.Builder
   n := len(s)
   if n == 0{
      return ""
   }
   stack := []string{}
   for i:=0;i<n;i++{
      if s[i]==32 && i>0 && s[i-1]!=32{
         stack = append(stack,sb.String())
         sb.Reset()
      } else if s[i]!=32{
         sb.WriteByte(s[i])
      }
   }
   if s[n-1] != 32{
      stack = append(stack,sb.String())
   }
   sb.Reset()
   for i:=len(stack)-1;i>=0;i--{
      sb.WriteString(stack[i])
      sb.WriteByte(' ')
   }
   res := sb.String()
	// 去掉末尾空格
	// 同时注意特殊情况：s全为空格，res长度为0，对其进行去除末尾空格操作会索引越界
   if len(res)==0{
      return ""
   }
   res = res[:len(res)-1]
   return res
}
```

还可以用双指针：（双指针一次过，nice）

声明变量 res 为空字符串，left 和 right 两指针初始化为 0，左右指针分别用于寻找每个单词的左右边界。

一次遍历字符串，左指针向右移动寻找第一个非空格字符，为要寻找字符串的左边界，然后开始寻找右边界，先将 left 的值赋给 right，然后 right 向右移动，直到到达字符串末尾或者寻找到空格字符，此时 left 和 right 分别指向一个单词的左右边界，将其加入 res，然后将 right 赋给 left，开始寻找下一个单词。

具体实现时，注意每加入一个单词的同时，还要添加一个空格，最后返回 res 前，处理掉最后添加的空格。

注意特殊情况的处理，当输入字符串全为空格时，res 长度为空，此时如果对其进行去除空格的操作，会索引越界。

```Go
func reverseWords(s string) string {
   n := len(s)
   left,right := 0,0
   res := ""
   for left < n{
      if s[left] == 32{
         left ++
      } else {
         right = left
         for right<n && s[right]!=32{
            right ++
         }
         res = s[left:right] + " " + res
         left = right
      }
   }
   if len(res) == 0{
      return ""
   } else {
      return res[:len(res)-1]
   }
}
```

这么一看，第一段代码其实也是双指针的思路，

### 21.调整数组顺序使奇数位于偶数前面

> 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
>
> 示例：
>
> 输入：nums = [1,2,3,4]
> 输出：[1,3,2,4] 
> 注：[3,1,2,4] 也是正确的答案之一。
>
>
> 提示：
>
> 0 <= nums.length <= 50000
> 0 <= nums[i] <= 10000

解题思路：双指针

left 和 right 初始化指向 nums 的第一个和最后一个元素下标，循环条件 left < right

- 左指针从当前位置向右移动寻找第一个偶数
- 右指针从当前位置向左移动寻找第一个奇数

以上过程要始终保持 left < right，不符合时直接跳出，然后 left 和 right 指向的元素互换位置。left++，right-- 持续以上步骤，直至 left <= right，说明全部奇数已位于偶数前面。

```Go
func exchange(nums []int) []int {
   n := len(nums)
   if n <= 1{
      return nums
   }
   left,right := 0,n-1
   for left < right{
      for left<right && nums[left]&1==1{
         left ++
      }
      for left<right && nums[right]&1==0{
         right --
      }
      nums[left],nums[right] = nums[right],nums[left]
      left ++
      right --
   }
   return nums
}
```

我曾经犯过一个错，这里描述一下，在 left 向右移 和 right 向左移的过程中，没有写 left<right，只有最外部的 left < right 条件，这会造成什么情况呢？

对 题目所给实例 [1,2,3,4] 来说，没有问题，left 和 right 分别指向 2 和  3 后，两元素互换位置，left 和 right 值分别为 2 和 1，跳出循环。

但考虑这样一种情况，数组元素全部为奇数，比如：[1,3,5]，那左指针在寻找偶数时，索引会超出数组长度，引发 panic。

再看一种情况，调整完的两个元素，left和right向内移动后，left<right，比如：[1,4,3,2,5]，left 寻找偶数找到 4，索引为 1，right寻找奇数位 5，索引为 4，交换位置后，各自向内移动，left=2，right=3，left继续寻找偶数元素，找到2，下标为3（此时已经=right了），right向左寻找奇数，找到 3，下标为2，两者交换位置后，left=4,right=1，结束循环

数组变为：[1,5,2,3,4]，发现多了一次数组元素交换位置。**错因就在寻找元素的过程中没有保证left<right。**

所以，大家一定要引以为鉴，不要犯类似的错误啦，寻找元素的过程一定要保证left<right



## day14.搜索与回溯算法（中等）

### 12.矩阵中的路径

> 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
>
> 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
>
> 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
>
> 示例 1：
>
> 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
> 输出：true
> 示例 2：
>
> 输入：board = [["a","b"],["c","d"]], word = "abcd"
> 输出：false
>
>
> 提示：
>
> m == board.length
> n = board[i].length
> 1 <= m, n <= 6
> 1 <= word.length <= 15
> board 和 word 仅由大小写英文字母组成

解题思路：DFS+回溯

题目出现网格，我们要寻找一条路径，很容易想到用 DFS，因为同一单元格内的字母不允许被重复使用，所以每使用一个字母，我们要对其进行记录，我们每一步很可能有多种选择，为避免使用到本条路径使用过的网格，我们需要进行回溯，用 used 二维数组记录每个元素在本条路径中是否被访问过。

另外，遇到这类需要上下左右移动的题目，可提前声明一个二维数组，保存上下左右移动时的元素下标变化，方便之后操作。

具体流程如下：

递归参数：i、j 和 k，i 和 j 为当前访问的网格元素下标，k 为当前目标字符在 word 中的下标。

终止条件：

- board\[i\]\[j\] != word[k] 即字符不匹配，说明此条路径走不通，返回 false，向前回溯尝试其他路径
- board\[i\]\[j\] = word[k]，且 k=len(word)-1，说明word所有字符匹配完成，存在指定路径，返回 true

搜索与回溯：

当前字符已经匹配成功，先标记当前下标为已访问，used\[i\]\[j]=true，函数结束时向前回溯需要取消此标记，结合go语言的特性，我们在这里用 defer 进行取消标记。

之后向四个方向进行搜索，先验证移动后坐标的合法性，然后判定该位置是否已被访问过，若合法且未访问过，在该位置进行递归操作，参数需要作出相应变化（坐标变化对应方向，k+=1）。

```Go
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
```

刚开始说到，出现网格，很容易想到用 DFS 或 BFS，那为什么我们选择使用 DFS 而没有选择使用 BFS 呢？因为我们要进行回溯操作！

在 DFS 中，我们在一条路走不通后进行回溯，契合我们对位置打标记记录其是否被访问过的思路，刚进入时做标记，在此道路尝试过所有路径后，进行回溯，取消标记。如果用 BFS 的话，很难实现这种打标记的思路。

### 13.机器人的运动范围

> 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
>
> 示例 1：
>
> 输入：m = 2, n = 3, k = 1
> 输出：3
> 示例 2：
>
> 输入：m = 3, n = 1, k = 0
> 输出：1
> 提示：
>
> 1 <= n,m <= 100
> 0 <= k <= 20

解题思路：BFS 或 DFS，要频繁地计算数字的各位数之和，我们可以声明一个函数专门用于计算该数字的位数之和，具体操作时不断取余和对 10 整除。

方法1：DFS

我们要统计机器人能够到达的格子数量，（0，0）是机器人的起始坐标，那我们从该坐标开始，向外 BFS，同时用一个used数组记录到达过的格子，避免重复统计一个格子。此外，我们要走的尽可能远，所以我们是不需要往上或者往左走的，只需要往远处行进（右或者下）。

具体操作：

DFS 的参数：i 和 j，代表当前即将访问的格子下标

DFS 结束条件：

索引越界、两个索引的位数之和大于 k 或者 该索引已经访问过，这些情况下，说明该格子无法到达或者已访问过，返回 0 即可。

搜索：首先，当前位置可访问，标记当前位置，used\[i\]\[j\] = true，然后利用系统栈实现DFS，向右和向下出发。

```Go
func movingCount(m int, n int, k int) int {
   used := make([][]bool,m)
   for i:=0;i<m;i++{
      used[i] = make([]bool,n)
   }
   var backtrace func(i,j int) int
   backtrace = func(i,j int) int{
      if i >=m || j >= n || bitsum(i)+bitsum(j)>k{
         return 0
      }
      if used[i][j]{
         return 0
      }
      used[i][j] = true
      return 1 + backtrace(i,j+1) + backtrace(i+1,j)
   }
   return backtrace(0,0)
}

func bitsum(x int) int {
   res := 0
   for x > 0{
      res += x % 10
      x /= 10
   }
   return res
}
```

方法2：BFS

本题中，虽然我们也使用 used 数组记录每个位置是否被访问过，但是我们不需要回溯操作，所以是可以使用 BFS 的。

下面这段代码封存起来吧。不知道为什么出错，输入 2,3,17， 输出 9，。

```Go
func movingCount_2(m int, n int, k int) int {
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
      if x==1 && y==1{
         fmt.Println(used[x][y])
      }
      // 检查参数合法性 及 该位置是否未被访问过
      // try right
      if y+1 < n && (bitsum(x) + bitsum(y+1)) <= k && !used[x][y+1]{
         q = append(q,[]int{x,y+1})
      }
      // try blow
      if x+1 < m && (bitsum(x+1) + bitsum(y)) <= k && !used[x+1][y]{
         q = append(q,[]int{x+1,y})
      }
      fmt.Println(q[0])
      q = q[1:]
   }
   return res
}
```

换了种写法，还是 9 ，不明白问题在哪里

```Go
func movingCount_2(m int, n int, k int) int {
   used := make([][]bool,m)
   for i:=0;i<m;i++{
      used[i] = make([]bool,n)
   }
   res := 0
   q := [][]int{{0,0}}
   for len(q) != 0{
      x,y := q[0][0],q[0][1]
      if x < m && y < n && bitsum(x)+bitsum(y)<=k && !used[x][y] {
         res ++
         q = append(q,[]int{x+1,y})
         q = append(q,[]int{x,y+1})
      }
      q = q[1:]
   }
   return res
}
```

知道问题在哪里了！

正确的代码如下：

```Go
func movingCount(m int, n int, k int) int {
   used := make([][]bool,m)
   for i:=0;i<m;i++{
      used[i] = make([]bool,n)
   }
   res := 0
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

func bitsum(x int) int {
   res := 0
   for x > 0{
      res += x % 10
      x /= 10
   }
   return res
}
```



## day15.搜索与回溯算法（中等）

### 34.二叉树中和为某一值的路径

> 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
>
> 叶子节点 是指没有子节点的节点。
>
> **提示：**
>
> - 树中节点总数在范围 `[0, 5000]` 内
> - `-1000 <= Node.val <= 1000`
> - `-1000 <= targetSum <= 1000`

方法一：DFS，枚举每一条从根节点到叶子节点的路径，遍历到叶子节点时，如果路径和恰好为 target，说明找到一个满足条件的路径。

注意代码中 copy(x,path) 操作，将 path 进行拷贝后再加入 res，若直接 res = append(res,path)，之后的路径将当前路径覆盖后，res 中的切片也会发生变化，这里是个小坑，需要注意一下。

```Go
//Definition for a binary tree node.
type TreeNode struct {
    Val int
   Left *TreeNode
    Right *TreeNode
}

//
func pathSum(root *TreeNode, target int) [][]int {
   res := [][]int{}
   path := []int{}
   var backtrace func(node *TreeNode,path []int,sum int)
   backtrace = func(node *TreeNode,path []int,sum int){
      if node == nil{
         return
      }
      path = append(path,node.Val)
      sum += node.Val
      if node.Left==nil && node.Right==nil{
         if sum == target{
            x := make([]int,len(path))
            copy(x,path)
            res = append(res,x)
            return
         } else {
            return
         }
      }
      backtrace(node.Left,path,sum)
      backtrace(node.Right,path,sum)
   }
   backtrace(root,path,0)
   return res
}
```

方法二：BFS

pair 结构体保存节点以及从根节点到该节点的路径值。

BFS 必定要用到队列，队列初始化时，加入根节点的 pair，sum 属性为 root.Val，同时，为了获取到叶子节点到根节点的路径，我们声明一个 parent 字典，保存节点的父节点，在进行 BFS 的同时更新 parent。

```go
func pathSum(root *TreeNode, target int) [][]int {
   res := [][]int{}
   // 判断 root 是否为空，对 root 为空的情况提前处理
   if root == nil{
      return res
   }
   // 队列初始化只保存根节点，利用队列实现 BFS
   q := []pair{{root,root.Val}}
   // parent字典帮助寻找叶子节点到根节点的路径
   parent := map[*TreeNode]*TreeNode{}
   // getPath利用parent寻找叶子节点到根节点的路径
   getPath := func(node *TreeNode) []int {
      path := []int{}
      // 不断向根节点回退
      for ;node!=nil;node=parent[node]{
         path = append(path,node.Val)
      }
      // path倒序，变为从根节点到叶子节点的路径
      left,right := 0,len(path)-1
      for left < right{
         path[left],path[right] = path[right],path[left]
         left ++
         right --
      }
      return path
   }
   // BFS
   for len(q) != 0{
      // 取队首元素
      cur := q[0]
      node := cur.node
      // 先更新 parent，将孩子节点指向该节点
      if node.Left != nil{
         parent[node.Left] = node
      }
      if node.Right != nil{
         parent[node.Right] = node
      }
      // 走到叶节点
      if node.Left==nil && node.Right==nil{
         // 且路径和等于target，更新res
         if cur.sum == target{
            res = append(res,getPath(node))
         }
      } else {
         // 若未到根节点
         // 将孩子节点加入队列
         if node.Left != nil{
            q = append(q,pair{node.Left,cur.sum+node.Left.Val})
         }
         if node.Right != nil{
            q = append(q,pair{node.Right,cur.sum+node.Right.Val})
         }
      }
      // 队首元素出队
      q = q[1:]
   }
   return res
}
```

### 36.二叉搜索树与双向链表

> 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
>
> 特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。

题解参考自：[剑指 Offer 36. 二叉搜索树与双向链表（中序遍历，清晰图解） - 二叉搜索树与双向链表 - 力扣（LeetCode）](https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/solution/mian-shi-ti-36-er-cha-sou-suo-shu-yu-shuang-xian-5/)

BST 的最重要的性质为：中序遍历为递增序列

我们一定是需要对 BST 进行中序遍历的，因为要改变指针的指向，所以遍历的同时，需要记录当前节点的前驱结点，声明变量 pre 和 head 为 *TreeNode 类型，pre 记录当前节点的前驱结点，head 记录最终返回的头结点。

中序遍历过程中，处理当前节点 cur 时，若 pre 为空，说明当前节点为中序遍历初始值，pre 为空，那自然也不需要调整 pre 和 cur 的指向，

若 pre 非空，调整 cur.Right 和 cur.Left 的指向，pre.Right = cur，cur.Left = pre；

处理完成后，pre = cur，当前节点为中序遍历中下一个节点的前驱结点。

递归结束后，pre 指向 BST 中节点值最大的节点，即中序遍历最后一个节点。

然后，调整 pre.Right 和 head.Left 的指向，形成循环链表。

很棒的一道题目。

这道题 LeetCode 上不支持 Go 语言，我去牛客网验证了一下，牛客网没有要求循环链表，代码倒数第三行和倒数第四行需要注释掉才能通过，LeetCode 要求循环，将其取消注释，理论上可以通过。

```go
package main

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

/**
 *
 * @param pRootOfTree TreeNode类
 * @return TreeNode类
 */
func Convert( pRootOfTree *TreeNode ) *TreeNode {
   // write code here
   var pre,head *TreeNode
   var inorder func(cur *TreeNode)
   inorder = func(cur *TreeNode) {
      if cur == nil{
         return
      }
      inorder(cur.Left)
      if pre == nil{
         head = cur
      } else {
         pre.Right = cur
         cur.Left = pre
      }
      pre = cur
      inorder(cur.Right)
   }
   inorder(pRootOfTree)
   // head.Left = pre
   // pre.Right = head
   return head
}
```

### 54.二叉搜索树的第k大节点

> 给定一棵二叉搜索树，请找出其中第 `k` 大的节点的值。
>
> **限制：**
>
> - 1 ≤ k ≤ 二叉搜索树元素个数

通俗解法：提到 BST，我们都知道其中序遍历是一个递增有序序列，先通过中序遍历得到该序列，然后返回倒数第 k 个值即可。

```Go
func kthLargest(root *TreeNode, k int) int {
   nums := []int{}
   var inorder func(node *TreeNode)
   inorder = func(node *TreeNode){
      if node == nil{
         return
      }
      inorder(node.Left)
      nums = append(nums,node.Val)
      inorder(node.Right)
   }
   inorder(root)
   return nums[len(nums)-k]
}
```

我们要找第k大的元素，而我们中序序列得到的是一个递增有序序列，导致我们需要保存下来中序遍历的元素才能找到第 k 大的元素，如果我们能够得到递减的有序序列，那我们就不再需要保存所有节点值，而是遍历到底 k 个有效节点时，保存该节点值，最后返回即可。

如何做到呢？中序遍历顺序是 [左子树的中序遍历，根节点，右子树的中序遍历]，其实只要把左右子树的遍历顺序颠倒一下，就可以得到从大到小的序列。我们用一个参数保存当前是遍历的第几个节点，不再需要一个数组保存中序遍历的序列，可以将空间复杂度从 O(n) 降低到 O(1)。

```Go
func kthLargest(root *TreeNode, k int) int {
   i := 0
   res := 0
   var inorder func(node *TreeNode)
   inorder = func(node *TreeNode) {
      if node == nil{
         return
      }
      inorder(node.Right)
      i ++
      if i == k{
         res = node.Val
      }
      inorder(node.Left)
   }
   inorder(root)
   return res
}
```



## day16.排序（简单）

### 45.把数组排成最小的数

> 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
>
> 示例 1:
>
> 输入: [10,2]
> 输出: "102"
> 示例 2:
>
> 输入: [3,30,34,5,9]
> 输出: "3033459"
>
>
> 提示:
>
> 0 < nums.length <= 100
> 说明:
>
> 输出结果可能非常大，所以你需要返回一个字符串而不是整数
> 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

这道题我刚开始以为数组的元素都是个位数，想着用哈希表统计每个数字出现的频率，然后从 0 开始遍历，将制定数目的数字添加至返回的在字符串中，看了下示例，才知道数组的元素可以不是个位数。

但其实原理也是类似的，小数在前，大数在后，本质上是一个排序问题，设 nums 中两整数分别为 x 和 y，先将其转化为字符串，然后拼接位 x+y 和 y+x 作比较

- 若拼接字符串 x+y > y+x，则 x > y
- 否则（包括相等的情况），x < y

根据此规则，套用任何排序算法对 nums 执行即可。

```Go
func minNumber(nums []int) string {
   n := len(nums)
   strNums := make([]string,len(nums))
   for i:=0;i<n;i++{
      strNums[i] = strconv.Itoa(nums[i])
   }
   sort.Slice(strNums, func(i, j int) bool {
      return strNums[i]+strNums[j] < strNums[j]+strNums[i]
   })
   res := ""
   for i:=0;i<n;i++{
      res += strNums[i]
   }
   return res
}
```

点过去看了下sort.Slice() 源码，在数组长度>12 的情况下，使用的是快速排序，经过递归，长度降至 12以下时，使用插入排序。

下面咱来不用标准库，自己复习下冒泡，来解这道题

```Go
func minNumber(nums []int) string {
   n := len(nums)
   strNums := make([]string,len(nums))
   for i:=0;i<n;i++{
      strNums[i] = strconv.Itoa(nums[i])
   }
   for i:=0;i<n-1;i++{
      for j:=0;j<n-1-i;j++{
         if strNums[j]+strNums[j+1] > strNums[j+1]+strNums[j]{
            strNums[j],strNums[j+1] = strNums[j+1],strNums[j]
         }
      }
   }
   res := ""
   for i:=0;i<n;i++{
      res += strNums[i]
   }
   return res
}
```

也是没有问题的

### 61.扑克牌中的顺子

> 从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
>
> 示例 1:
>
> 输入: [1,2,3,4,5]
> 输出: True
>
>
> 示例 2:
>
> 输入: [0,0,1,2,5]
> 输出: True
>
>
> 限制：
>
> 数组长度为 5 
>
> 数组的数取值为 [0, 13] 
>

解题思路：排序

给定数组长度等于 5，数据量很小，我们先对其排序

然后一次遍历，检查是否存在重复元素（0除外），若存在，一定不连续

此时，数组中除 0 外，不存在重复元素，得到其最大值和最小值（0除外），若最大值减去最小值 大于 4，说明不连续，否则，连续

```Go
func isStraight(nums []int) bool {
   sort.Ints(nums)
   for i:=1;i<5;i++{
      if nums[i] == 0 {
         continue
      }
      if nums[i]==nums[i-1] {
         return false
      }
   }
   max := nums[len(nums)-1]
   min := nums[0]
   for i:=0;nums[i]==0;i++{
      min = nums[i+1]
   }
   if max - min > 4{
      return false
   }
   return true
}
```

另外，我们也可以不用排序，而是使用一次遍历的方式求得最大最小值（0除外）
思路和排序是一致的，代码如下
这种解法注意 minNum 和 maxNum 的初始化，minNum 和 maxNum 分别初始化为最大和最小的数字

```Go
func isStraight(nums []int) bool {
   minNum,maxNum := 14,0
   record := map[int]struct{}{}
   for _,num := range nums{
      if num == 0{
         continue
      }
      if _,ok:=record[num];ok{
         return false
      }
      record[num] = struct{}{}
      minNum = min(minNum,num)
      maxNum = max(maxNum,num)
   }
   if maxNum - minNum > 4{
      return false
   }
   return true
}

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}

func min(x,y int) int{
   if x < y{
      return x
   }
   return y
}
```



## day17.排序（中等）

### 40.最小的k个数

> 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
>
> 示例 1：
>
> 输入：arr = [3,2,1], k = 2
> 输出：[1,2] 或者 [2,1]
> 示例 2：
>
> 输入：arr = [0,1,2,1], k = 1
> 输出：[0]
>
>
> 限制：
>
> 0 <= k <= arr.length <= 10000
> 0 <= arr[i] <= 10000

最通俗的解法就是排序了

```Go
func getLeastNumbers(arr []int, k int) []int {
   quickSort(arr,0,len(arr)-1,k)
   return arr[:k]
}
```

这么做能满足要求，但没有意义

接下来思考一下，题目要找的是最小的 k 个数，但没有要求这 k 个数字按顺序，很容易想到快排，每一趟快排会确定一个元素的最终位置，其左侧元素值均小于该元素，右侧元素值均大于该元素，而左右侧元素是否有序快排并不关心，这与本题的要求不谋而合。

下面用快排的思想来解题。

```go
func getLeastNumbers_2(arr []int, k int) []int {
   quickSort(arr,0,len(arr)-1,k)
   return arr[:k]
}

func quickSort(nums []int,left,right,k int) []int {
   if left < right{
      // partitionIndex为该次分区后，基准元素所在下标
      partitionIndex := partition(nums,left,right)
      //若 k大于该下标，我们只需要对左侧区间递归进行快排，
      // 右侧区间无需处理
      if partitionIndex < k{
         return quickSort(nums,partitionIndex+1,right,k)
      // k小于该下标，同理
      } else if partitionIndex > k{
         return quickSort(nums,left,partitionIndex-1,k)
      } else {
         // 当相等时，达到我们的要求，返回“排好序”的数组
         return nums
      }
   }
   return nums
}

//移动左右指针，按照基准（这里使用nums[left]）划分区域。最后返回基准所在的下标
func partition(nums []int,left,right int) int {
   // 基准 选择 left 指向的元素
   pivot := nums[left]
   for left < right{
      for left<right && nums[right]>=pivot{
         right --
      }
      nums[left] = nums[right]
      for left<right && nums[left]<=pivot{
         left ++
      }
      nums[right] = nums[left]
   }
   nums[left] = pivot
   return left
}
```

### 41.数据流中的中位数

> 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
>
> 例如，
>
> [2,3,4] 的中位数是 3
>
> [2,3] 的中位数是 (2 + 3) / 2 = 2.5
>
> 设计一个支持以下两种操作的数据结构：
>
> void addNum(int num) - 从数据流中添加一个整数到数据结构中。
> double findMedian() - 返回目前所有元素的中位数。
> 示例：
>
> 输入：
> ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
> [[],[1],[2],[],[3],[]]
> 输出：[null,null,null,1.50000,null,2.00000]

解题思路：用两个堆来维护数据流，将数据流根据元素的大小一分为二

大根堆维护数据流元素值较小的一半，小根堆维护数据流元素值较大的一半

当数据流长度为 偶数 时，大小根堆长度相同，当数据流长度为奇数时，我们规定，小根堆存储中位数。

向数据流添加元素 num 时，可分为两种情况

- 大根堆与小根堆长度相等，此时，应向小根堆添加元素，但添加的元素并不一定是 num
	- 若 num 大于小根堆堆顶元素值，则 num 属于元素值较大的一部分，num 直接插入大根堆
	- 否则，先将 num 插入大根堆，然后取出小根堆堆顶元素，插入大根堆
- 大根堆与小根堆长度不相等，此时，应向大根堆添加元素，但添加的元素同样并不一定是 num
	- 若 num 大于小根堆堆顶元素值，将 num 插入小根堆，然后取出小根堆堆顶元素，插入大根堆
	- 否则，num 直接插入大根堆

```go
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
```



## day18.搜索与回溯算法（中等）

### 55-I.二叉树的深度

> 输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
>
> 例如：
>
> 给定二叉树 [3,9,20,null,null,15,7]，
>
>    3
>
>    / \
>   9  20
>     /  \
>    15   7
> 返回它的最大深度 3 。
>
>  
>
> 提示：
>
> 节点总数 <= 10000

因为二叉树的定义是递归的，所以针对二叉树的题目，天然就有递归的解法，且代码通俗易懂，对本题，我们由递归和迭代（BFS）两种解法

树的遍历方式分为两类：DFS（前序遍历、中序遍历、后序遍历）和 BFS（层序遍历）

递归对应的是 DFS，要求某根节点为 root 的 树的深度，分两种情况

- root=nil，即该节点为空，返回 0
- root!=nil，说明该树深度起码为 1，具体是多少，还要看左右子树的深度，取其较大值，然后加上1，返回的是 1+max(maxDepth(root.Left),maxDepath(root.Right))

```Go
//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
   if root == nil{
      return 0
   }
   return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}
```

解法2：迭代（层序遍历）

二叉树的 BFS，BFS 通常借助队列的先进先出特性来实现。

首先还是特殊情况处理，若 root 为空，返回 0；

初始化变量 res=0，队列 q，仅包含 root 一个元素，因为该树第一层有且仅有 root 一个节点，然后开始常规 BFS

- 每开始一层的遍历，res+=1，
- 遍历当前队列中元素，若其存在左右子节点，将其添加至队列中
- 当前层所有元素出队

最终返回 res

```go
func maxDepth(root *TreeNode) int {
   if root == nil {
      return 0
   }
   res := 0
   q := []*TreeNode{root}
   for len(q) != 0{
      res ++
      n := len(q)
      for i:=0;i<n;i++{
         cur := q[i]
         if cur.Left != nil{
            q = append(q,cur.Left)
         }
         if cur.Right != nil{
            q = append(q,cur.Right)
         }
      }
      q = q[n:]
   }
   return res
}
```

### 55-II.平衡二叉树

> 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
>
> **限制：**0 <= 树的结点个数 <= 10000

这道题中的平衡二叉树的定义是：二叉树的每个节点的左右子树的高度差的绝对值不超过 1，则二叉树是平衡二叉树。根据定义，一棵二叉树是平衡二叉树，当且仅当其所有子树也都是平衡二叉树，因此可以使用递归的方式判断二叉树是不是平衡二叉树，递归的顺序可以是自顶向下或者自底向上。（这句话来自官方题解）

AVL 的定义是递归的，因此很容易想到递归的解法。

求当前节点左右子树的高度，然后判断高度之差是否不超过1，然后递归判断其左右子树是否为 AVL

```go
func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	leftHeight := maxDepth(root.Left)
	rightHright := maxDepth(root.Right)
	return abs(leftHeight-rightHright) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
   if root == nil{
      return 0
   }
   return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}

func abs(x int) int {
   if x < 0{
      return -x
   }
   return x
}
```

刚才这种递归是自顶向下的，对于同一个节点，函数 maxDepth 会被重复调用，调用次数为其祖先节点的个数。

如果使用自底向上的递归，则对每个节点，函数 maxDepth 只会被调用一次。

本解法参考自官方题解：[平衡二叉树 - 平衡二叉树 - 力扣（LeetCode）](https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/solution/ping-heng-er-cha-shu-by-leetcode-solutio-6r1g/)

自底向上的做法类似后序遍历，对当前遍历到的节点，先递归地判断其左右子树是否平衡，再判断以当前节点为根节点的子树是否平衡，这样，在其子树不平衡时，我们直接返回 false即可。如果一棵树是平衡的，返回其高度，否则返回 -1。只要存在子树不平衡，则整棵树必定不平衡。

```go
func isBalanced(root *TreeNode) bool {
   return depth(root) >= 0
}

func depth(node *TreeNode) int {
   if node == nil{
      return 0
   }
   leftDepth := depth(node.Left)
   rightDepth := depth(node.Right)
   if leftDepth == -1 || rightDepth == -1 {
      return -1
   }
   if abs(leftDepth-rightDepth)>1{
      return -1
   }
   return 1+max(leftDepth,rightDepth)
}

func abs(x int) int {
   if x < 0{
      return -x
   }
   return x
}

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}
```



## day19.搜索与回溯算法（中等）

### 64.求1+2+...+n

> 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
>
> 示例 1：
>
> 输入: n = 3
> 输出: 6
> 示例 2：
>
> 输入: n = 9
> 输出: 45
>
>
> 限制：
>
> 1 <= n <= 10000

先不考虑限制条件，有三种解法

- 高斯求和公式（使用乘除法）
- 迭代（使用for或while）
- 递归（递归终止条件需要使用if）

考虑限制条件的话，难点在于如何实现这个 1~n 的循环。

我们可以利用运算符的短路效应实现该循环

先看下逻辑与 与 逻辑或 的短路效应

- if a && b； 若 a 为 false，对 b 的判断不会执行（即短路），直接返回 false
- if  a || b； 若 a 为 true，对 b的判断不会执行，直接返回 true

我们可以利用这种短路机制，通过递归实现 1~n 的相加

```go
func sumNums(n int) int {
   res := 0
   var f func(*int,int) bool
   f = func(res *int,n int) bool{
      *res += n
      return n > 0 && f(res,n-1)
   }
   f(&res,n)
   return res
}
```

### 68-I.二叉搜索树的最近公共祖先

> 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
>
> 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
>

BST 的中序遍历是有序数组，root 节点的左子树中所有节点 val 都比 root.val 小，右子树中所有节点的 val 比root.val 大。根据这个特点，我们对 root、p、q 的节点值进行分情况讨论：

1. p.val > root.val 且 q.val < root.val，说明 p 和 q 分别存在于 root 节点的右左子树中，root 节点为最近公共祖先（p.val < q.val 同理）；

2. p.Val < root.Val 且 q.Val > root.Val，与第一种情况同理，p、q 分别在root 的左右子树中，root 节点为最近公共祖先；

3. root.Val == p.Val 或者 root.Val == q.Val，说明另一个节点在 root 的左子树或者右子树中， root 已经是两节点的最近公共祖先，因为我们是从整棵树的根节点开始往下遍历的

4. p.Val < root.Val 且 q.Val < root.Val，p 和 q 均在 root 的左子树中，root 的左孩子节点会是 p 和 q 的公共祖先，对 root.Left 进行递归操作寻找 p 和 q 的最近公共祖先
5. p.Val > root.Val 且 q.Val > root.Val，与情况 4 同理，对 root.Right 进行递归操作寻找 p 和 q 的最近公共祖先。

因为 1 和 2同理，4 和 5 同理，精简一些，只有三种情况，令 p.Val 小于 q.Val

1. p.Val <= root.Val 或者 root.val < q.Val，此时 root 为 最近公共祖先
2. p.Val < root.Val < q.Val，p 和 q 分别在 root 的左右子树中，root 为 最近公共祖先
3. p.Val < root.Val 且 q.Val < root.Val，对 root.Left 递归操作；若均大于，对 root.Right 递归操作

虽然我分析过程中写的是递归，但实际还是有递归和迭代两种解法的

递归

```go
//Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   if root.Val<p.Val && root.Val<q.Val{
      return lowestCommonAncestor(root.Right,p,q)
   }
   if root.Val>p.Val && root.Val>q.Val{
      return lowestCommonAncestor(root.Left,p,q)
   }
   return root
}
```

迭代

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   for root != nil {
      if p.Val<root.Val && q.Val<root.Val{
         root = root.Left
         continue
      }
      if p.Val>root.Val && q.Val>root.Val{
         root = root.Right
         continue
      }
      return root
   }
   // 此种情况不会发生，因为题目说明p和q均存在于给定BST中
   // 只是为了保证代码的逻辑正确
   return root
}
```

### 68-II.二叉树的最近公共祖先

> 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
>
> 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
>
> **说明:**
>
> - 所有节点的值都是唯一的。
> - p、q 为不同节点且均存在于给定的二叉树中。

取消了第一版题目BST的条件，退化为普通的二叉树，我们就只能相对暴力地去解决这个问题。

解法参考自题解：[剑指 Offer 68 - II. 二叉树的最近公共祖先（DFS ，清晰图解） - 二叉树的最近公共祖先 - 力扣（LeetCode）](https://leetcode.cn/problems/er-cha-shu-de-zui-jin-gong-gong-zu-xian-lcof/solution/mian-shi-ti-68-ii-er-cha-shu-de-zui-jin-gong-gon-7/)

考虑通过递归对二叉树进行先序遍历，遇到 p 或 q 时返回，从底至顶回溯，当节点 p、q 分别在 root 的两侧 或 p、q 两者之一就是 root 节点时，root 为 p 和 q 的LCA，返回 root。

然后具体分析如何实现：

1. 终止条件：
	1. 当越过叶节点时，直接返回 nil；
	2. root 等于 p 或 q 时，返回 root
2. 递推工作：
	1. 递归 root.Left，记返回值为 left；
	2. 递归 root.Right，记返回值为 right；
3. 返回值：根据 left 和 right，分为以下四种情况：
	1. left 和 right 同时为空，说明 root 的左右子树均不包含 p 和 q，以 root 为根节点的子树不可能包含 p 和 q，返回 nil
	2. left 为 空，right 不为空，说明 p 和 q 都包含在 root.Right 中，返回 right（分两种情况，第一种：p 和 q 其中一个节点在 root 的右子树中，此时 right 为所包含的节点；第二种：p 和 q 均在 root 的右子树中，此时 right 指向 p 和 q 的 LCA）
	3. right 为 空，left 不为空，与情况2同理
	4. left 和 right 均不为空，说明 p、q 分别在 root 的左右子树中，root 为 LCA，返回 root 即可。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
   // 递归结束条件
   if root == nil{
      return nil
   }
   if root.Val==p.Val || root.Val==q.Val{
      return root
   }
   // 开始递归
   left := lowestCommonAncestor(root.Left,p,q)
   right := lowestCommonAncestor(root.Right,p,q)

   if left==nil && right==nil{
      return nil
   }
   if left == nil{
      return right
   }
   if right == nil{
      return left
   }
   return root
}
```

这里想看过的一段话：
学习一个算法，最要的是弄清楚这个算法要解决什么样的问题，它的已知量（input）是什么？待求的未知量（output）是什么？
如果这三个问题没找到答案就去学习算法，就会浪费大量时间在逻辑猜测与记忆上

函数 lowestCommonAncestor 输入为 root、p 和 q

输出分几种情况：

1. 若 root 为空节点，直接返回 root 即可
2. 以 root 为根节点的树包含 p 和 q，返回 p 和 q 的 LCA；
3. 以 root 为根节点的树只包含 p 或 q 其中一个，则返回 包含的节点；
4. 以 root 为根节点的树不包含 p 且 不包含 q，返回 nil

解决问题：在以 root 为根节点的子树中，寻找 p 和 q 的 LCA



## day20.分治算法（中等）

### 07.重建二叉树

> 输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
>
> 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
>
> **限制：**
>
> 0 <= 节点个数 <= 5000

解题思路：参考自题解[重建二叉树 - 重建二叉树 - 力扣（LeetCode）](https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/solution/mian-shi-ti-07-zhong-jian-er-cha-shu-by-leetcode-s/)

对于任意一颗树而言，前序遍历的形式总是

[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]

即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是

[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]



根据二叉树的前序遍历，我们可以构造出根节点，之后要考虑的是如何构造根节点的左孩子节点和右孩子节点。

只要我们在中序遍历中定位到根节点，我们就可以知道该二叉树左子树和右子树的节点数量。知道数量后，我们就可以在前序遍历和中序遍历中通过切片得到左子树和右子树的前序遍历和中序遍历。

之后递归构造根节点的左孩子节点和右孩子节点即可

在中序遍历中对根节点进行定位时，我们这里使用遍历的方式

```go
//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
   // 递归退出条件，遍历序列长度为0，节点为空
   if len(preorder) == 0{
      return nil
   }
   rootVal := preorder[0]
   // 构造根节点
   root := &TreeNode{rootVal,nil,nil}
   // k初始化为0，寻找中序遍历中根节点的下标
   k := 0
   for ;k<len(preorder);k++{
      if inorder[k] == rootVal{
         break
      }
   }
   // 递归构造根节点的左孩子节点
   root.Left = buildTree(preorder[1:k+1],inorder[:k])
   // 递归构造根节点的右孩子节点
   root.Right = buildTree(preorder[k+1:],inorder[k+1:])
   return root
}
```

### 16.数值的整数次方

> 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
>
> 示例 1：
>
> 输入：x = 2.00000, n = 10
> 输出：1024.00000
> 示例 2：
>
> 输入：x = 2.10000, n = 3
> 输出：9.26100
> 示例 3：
>
> 输入：x = 2.00000, n = -2
> 输出：0.25000
> 解释：2-2 = 1/22 = 1/4 = 0.25
>
>
> 提示：
>
> -100.0 < x < 100.0
> -2^31 <= n <= 2^31-1
> -10^4 <= xn <= 10^4

这道题，首先要和面试官沟通关于数学上的定义，0的0次方可以是0，也可以是1，和面试官交流后再写代码，确定这种情况下返回0还是1 。

然后处理 n<0 的情况，n小于0时，x取1/x（为防止分母为0，提前对x等于0的情况进行处理），n取其相反数，将n化为正数，可统一化算法流程。

通俗解法就是循环 n 次，每次循环返回变量对 x 作乘积。这种做法提交后会超时！

考虑算法：快速幂

迭代解法的时间复杂度为 O(n)，快速幂可将时间复杂度降低至 O(logn)。

咱从二进制角度分析快速幂算法

设 n 的二进制长度为 m+1，则 n 的二进制表示和二进制转十进制表示如下：
$$
n=(0b)b_mb_{m-1}...b_1b_0=b_m*2^m+b_{m-1}*2^{m-1}+...+b_1*2^{1}+b_0*2^0
$$
则 x 的 n 次方可展开为：
$$
x^n=x^{b_m*2^m+b_{m-1}*2^{m-1}+...+b_1*2^{1}+b_0*2^0}=x^{b_m*2^m}*...*x^{b_0*2^0}
$$
那我们求 x 的 n 次方就可以转化为这 m+1 项的乘积，对这 m+1 项，从后向前，依次计算

```go
func myPow(x float64, n int) float64 {
   if x == 0{
      return x
   }
   var res float64 = 1
   if n < 0{
      x = 1/x
      n = -n
   }
   factor := x
   for n != 0{
      if n & 1 == 1{
         res *= factor
      }
      factor *= factor
      n >>= 1
   }
   return res
}
```

### 33.二叉树的后序遍历序列

> 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
>
> 参考以下这颗二叉搜索树：
>
>  5
> / \
>
>    2   6
>   / \
>  1   3
> 示例 ：
>
> 输入: [1,6,3,2,5]
> 输出: false
> 示例 2：
>
> 输入: [1,3,2,6,5]
> 输出: true
>
>
> 提示：
>
> 数组长度 <= 1000

一句话描述：后序遍历的根节点在最后一位，前面可以分为两部分，一部分为左子树的后序遍历，另一部分为右子树的后序遍历，右子树所有节点值均大于根节点，找到右子树节点的起始节点，判断该部分是否所有节点值都大于根节点的值，若否，说明该序列无法构成BST的后序遍历序列，之后再递归判断左子树部分和右子树部分的序列是否为BST的后序遍历序列。

下面是详细描述：

二叉树后序遍历的结果是：[ [左子树的后序遍历结果], [右子树的后序遍历结果]，根节点]

我们可以把后序遍历数组分为三部分：左子树的后序遍历、右子树的后序遍历和该二叉树的根节点。二叉树的根节点一定是后序遍历数组的最后一个元素。

BST左子树所有元素节点值均小于根节点，右子树所有元素值均大于根节点。所以，我们可以从后序遍历数组的第一个元素遍历到倒数第二个元素，进行判断，是否大于等于根节点，若大于等于，说明该元素之前所有元素为左子树的后序遍历结果，从该元素至倒数第二个元素为右子树的后序遍历结果。从该元素起，遍历到倒数第二个元素，若存在元素值小于根节点的值，说明该数组不是BST的后序遍历序列。 然后递归地判断左子树序列与右子树序列是否为某BST的后序遍历序列。

有一点需要注意的是，right的初始化应该是序列长度-1，我刚开始初始化为0，结果就出错了。

如果将 k 初始化为 0，且右子树区间大小为 0，最后 k 值为 0,递归验证右子树的区间会是 postorder[0:n-1]，而实际上该区间大小为 0,所以我们初始化 k 时，需要其值 大于等于 n-1。

这是容易出错的一个点，需要牢记

```Go
func verifyPostorder(postorder []int) bool {
	// 若数组大小小于等于2，则必定是某BST的后序遍历序列
	if len(postorder) <= 2{
		return true
	}
	n := len(postorder)
	// 后序遍历的根节点在数组最后一位
	root := postorder[n-1]
	// k 为右子树的后序遍历序列起始下标
	// 初始化为 n-1
	k := n-1
	// 寻找 k 值
	for i:=0;i < n-1;i++{
		if postorder[i] > root{
			k = i
			break
		}
	}
	// 验证根节点的右子树所有节点值均大于根节点值
	// 若存在小于根节点的节点，返回 false
	for i:=k+1;i<n;i++{
		if postorder[i] < root{
			return false
		}
	}
	// 验证左子树的后序遍历序列
	left := verifyPostorder(postorder[:k])
	// 验证右子树的后序遍历序列
	right := verifyPostorder(postorder[k:n-1])
	return left && right
}
```



## day21.位运算（简单）

### 15.二进制中 1 的个数

> 编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。
>
> 提示：
>
> 请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
> 在 Java 中，编译器使用 二进制补码 记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

方法1：逐位判断，循环检查每一个二进制位是否为 1

对 Golang，题目输入为一个 uint32 类型的数字，但它实际代表的是一个数的二进制表示形式。我们可以进行 32 次的循环，判断每一位是否为 1 。

```go
func hammingWeight(num uint32) int {
   res := 0
   for i:=0;i<32;i++{
      if num & 1 == 1{
         res ++
      }
      num >>= 1
   }
   return res
}
```

方法2：位运算优化，消除二进制末尾的 1

非常巧妙的做法，当 n 非零时，n=n&(n-1)可消除n的二进制中最后一个出现的 1.

因此，执行 n=n&(n-1)使得n变成 0 的操作次数，就是 n 的二进制中 1 的个数。

```go
func hammingWeight_2(num uint32) int {
   res := 0
   for num != 0{
      res ++
      num &= num-1
   }
   return res
}
```

第一种解法的时间复杂度为 O(logn)，第二种解法为 O(m)，m 为二进制串中 1 的个数；

两种解法的空间复杂度均为 O(1)。

个人感觉啊，位运算优化只是一种技巧，在效率方面对比逐位检查提升并不大。因为逐位检查已经是对数级别了，对数级别的时间复杂度已经是相当高效。

### 55.不用加减乘除做加法

> 写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
>
> 示例:
>
> 输入: a = 1, b = 1
> 输出: 2
>
>
> 提示：
>
> a, b 均可能是负数或 0
> 结果不会溢出 32 位整数

解题思路：位运算（参考自官方题解）

其实主要就是处理好进位

首先，考虑两个二进制位相加的四种情况如下：

```
0 + 0 = 0
0 + 1 = 1
1 + 0 = 1
1 + 1 = 0 (进位)
```

可以发现，对于整数 a 和 b：

在不考虑进位的情况下，其无进位加法结果为 a ^ b。

而所有需要进位的位为 a & b，进位后的进位结果为 a & b << 1。

so，我们可以将 a+b 的和拆解为 无进位的和 与 进位结果 的和。

在代码中，我们可以直接将 b 看做进位结果的和进行处理，之后不断进行求 无进位加法结果 和 进位结果 的过程，将进位结果赋值给 b，直至进位结果等于 0，说明加法结束，返回 a 即可。

```go
func add(a int, b int) int {
   for b != 0{
      carry := (a & b) << 1
      a ^= b
      b = carry
   }
   return a
}
```



## day22.位运算（中等）

### 56-I.数组中数字出现的次数 I

> 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
>
> 示例 1：
>
> 输入：nums = [4,1,4,6]
> 输出：[1,6] 或 [6,1]
> 示例 2：
>
> 输入：nums = [1,2,10,4,1,4,3,3]
> 输出：[2,10] 或 [10,2]
>
>
> 限制：
>
> 2 <= nums.length <= 10000

本题最通俗的解法是哈希表记录每个数字出现的次数，不符合题目要求空间复杂度，位运算解法可将空间复杂度从 O(n) 降低至 O(1)

解题思路：分组异或

先考虑一种简单的情况，一个整形数组 nums 中除一个数字出现一次之外，其余数字都出现了两次，找出这个数字。

很容易想到用 异或 来解决。两个相同的数字进行异或操作结果为 0。对 nums 进行一次遍历，xor 变量初始化为 0，与 nums 中每个数字进行异或操作，相同的数字会两两抵消，最后 xor 就与仅出现一次的数字值相同。

回到本题，有两个出现一次的数字，设为 a 和 b，如果我们使用相同的方法对 nums 进行异或，最终得到的结果 xor 会是 a 与 b 异或的结果。xor 必定不等于 0，因为若 xor 等于 0，说明 a==b，与条件相悖。

我们可以根据 xor 的二进制中任意不等于 0 的位进行分组，将 a 和 b 分到不同的组中，然后对两组数字分别进行异或，取出 a 和 b。

假设 xor 的第二位为 1，那我们就取 x = 10，对 nums 进行遍历，记当前数字为 num，根据 num&x==0和 num&x!=0 分组，a 和 b 在各自组中异或得到。

```go
func singleNumbers(nums []int) []int {
   xor := 0
   for _,num := range nums{
      xor ^= num
   }
   x := 1
   for {
      if xor & x == 0{
         x <<= 1
      } else {
         break
      }
   }
   a,b := 0,0
   for _,num := range nums{
      if num & x == 0{
         a ^= num
      } else {
         b ^= num
      }
   }
   return []int{a,b}
}
```

### 56-II.数组中数字出现的次数 II

> 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
>
> 示例 1：
>
> 输入：nums = [3,4,3,3]
> 输出：4
> 示例 2：
>
> 输入：nums = [9,1,7,9,7,9,7]
> 输出：1
>
>
> 限制：
>
> 1 <= nums.length <= 10000
> 1 <= nums[i] < 2^31

本题最通俗的解法是哈希表记录每个数字出现的次数，我们还是用位运算进行优化，将空间复杂度从 O(n) 降低至 O(1)

解题思路：依次确定二进制位

记数组中出现一次的数字为 res，题目给定数组中整数范围为 1~2^31-1，我们可以从最低位开始，依次确定 res 当前位 是 0 还是 1 。

在每一位中，我们遍历数组所有元素，统计当前位为 1 的个数 num，若 num 对 3 取余 为 0，说明 res 当前位为 0，否则为 1 。

```go
func singleNumber(nums []int) int {
   n := len(nums)
   res := 0
   for i:=0;i<32;i++{
      num := 0
      bit := 1<<i
      for j:=0;j<n;j++{
         if nums[j] & bit != 0{
            num ++
         }
      }
      if num % 3 == 1{
         res |= bit
      }
   }
   return res
}
```



## day23.数学（简单）

### 39.数组中出现次数超过一半的数字

> 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
>
> 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
>
> 示例 1:
>
> 输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
> 输出: 2
>
>
> 限制：
>
> 1 <= 数组长度 <= 50000

本题通俗解法为哈希表统计数组中每个元素出现次数，然后遍历哈希表键值对，找到值大于数组长度一般的键，返回即可。时间空间复杂度均为 O(n)

方法二：排序，将 nums 元素从小到大排序，那下标为 n/2 的元素为要寻找的元素

考虑一种更好的解法

方法三：摩尔投票算法（参考自官方题解）

算法步骤如下：

1. 维护一个候选主要元素 candidate 和候选主要元素出现的次数 count，初始时 candidate 为任意值，count 为 0；
2. 遍历数组 nums 中的元素，设元素值为 x，依次对元素执行如下操作：
	如果 count = 0，则将 x 赋给 candidate，否则不更新 candidate 的值；
	如果 x = candidate，count 自增 1，否则将 count 减 1 。
3. 遍历结束后，如果 nums 中存在主要元素，则 candidate 为主要元素，否则 candidate 可 能为数组中的任意一个元素。

由于数组不一定存在主要元素，因此需要第二次遍历数组，验证 candidate 是否为主要元素。这次遍历统计 candidate 在数组中出现的次数，若大于数组长度的一半，则 candidate 是主要元素，否则，不是。

为什么当数组中存在主要元素时，Boyer-Moore 投票算法可以确保得到主要元素？
在 Boyer-Moore 投票算法中，遇到相同的数则将 count 加 1，遇到不同的数则将 count 减 1。根据主要元素的定义，主要元素的出现次数大于其他元素的出现次数之和，因此在遍历过程中，主要元素和其他元素两两抵消，最后一定剩下至少一个主要元素，此时 candidate 为主要元素，且 count≥1。

```go
func majorityElement(nums []int) int {
   candidate := 0
   count := 0
   for _,num := range nums{
      if count == 0{
         candidate = num
         count = 1
         continue
      }
      if num == candidate{
         count ++
      } else {
         count --
      }
   }
   return candidate
}
```

### 66.构建乘积数组

> 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
>
> 示例:
>
> 输入: [1,2,3,4,5]
> 输出: [120,60,40,30,24]
>
>
> 提示：
>
> 所有元素乘积之和不会溢出 32 位整数
> a.length <= 100000

若可以使用除法的话，我们可以先一次遍历，得到所有数字的乘积 x，然后创建新的数组，一次遍历，每次 x 除以当前遍历到的元素，添加到新数组即可。

不使用乘法，我的想法是，构建前缀乘积数组 prefix 和后缀乘积数组 suffix，长度均与给定数组长度相同，设 n 为给定数组长度

- prefix[i] = a[0] * a[1] * ... * a[i-1]
- suffix[i] = a[i+1] * a[i+2] * ... * a[n-1]

然后创建乘积数组 res, res[i] = prefix[i] * suffix[i]

```go
func constructArr(a []int) []int {
   n := len(a)
   //当输入数组长度为 0 时，直接返回空数组，
   //若不进行处理，会在下面代码 suffix[n-1]报错超出索引
   if n == 0{
      return []int{}
   }
   // 构建前缀乘积数组和后缀乘积数组
   prefix,suffix := make([]int,n),make([]int,n)
   res := make([]int,n)
   prefix[0],suffix[n-1] = 1,1
   for i:=1;i<n;i++{
      prefix[i] = prefix[i-1] * a[i-1]
   }
   for i:=n-2;i>=0;i--{
      suffix[i] = suffix[i+1] * a[i+1]
   }
   for i:=0;i<n;i++{
      res[i] = prefix[i] * suffix[i]
   }
   return res
}
```

时间空间复杂度均为 O(n)

思考如何进行优化，时间复杂度的 O(n) 是没有优化空间的，只能从空间复杂度下手

通过两轮循环，分别计算 prefix 和 suffix 的乘积。

先从上到下计算前缀乘积数组，再从下到上计算后缀乘积数组，因为不需要额外空间存储，只用 res 数组即可，将空间复杂度从 O(n) 降低至 O(1) （此空间复杂度不包含返回数组）。

```go
func constructArr_2(a []int) []int {
   n := len(a)
   if n == 0{
      return []int{}
   }
   res := make([]int,n)
   res[0] = 1
   // 计算前缀树组乘积
   for i:=1;i<n;i++{
      res[i] = res[i-1] * a[i-1]
   }
   //计算后缀数组乘积
   temp := 1
   for i:=n-2;i>=0;i--{
      temp *= a[i+1]
      res[i] *= temp
   }
   return res
}
```



## day24.数学（中等）

### 14-I.剪绳子

> 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
>
> 示例 1：
>
> 输入: 2
> 输出: 1
> 解释: 2 = 1 + 1, 1 × 1 = 1
> 示例 2:
>
> 输入: 10
> 输出: 36
> 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
> 提示：
>
> 2 <= n <= 58
> 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/ （整数拆分）

先上传统的动态规划解法

1. 确定dp数组大小及下标含义：dp[i] 代表长度为 i 的绳子切割后的最大乘积，长度为 n+1
2. dp 数组初始化：dp[1]=1,dp[2]=2,dp[3]=3,需要注意的是，该初始值只针对 n>=4 的情况，n<4的情况，我们会单独处理
3. 状态转移方程：dp[i] = max(dp[i] * dp[i-j])，j 从 1 遍历到 i-1，这里我们可以使用一个小技巧，j 从 1 遍历至 n/2 即可，因为 1*6 = 6 * 1

再说下 n<4 的情况，题目给定 n>=2，当 n<4 时，返回 n-1 即可。

```go
func cuttingRope(n int) int {
   if n < 4{
      return n-1
   }
   dp := make([]int,n+1)
   dp[1],dp[2],dp[3] = 1,2,3
   for i:=4;i<=n;i++{
      for j:=1;j<=i/2;j++{
         dp[i] = max(dp[i],dp[j]*dp[i-j])
      }
   }
   return dp[n]
}

func max(x,y int) int{
   if x > y{
      return x
   }
   return y
}
```

这道题还有一种数学的解法，时间和空间复杂度都超过dp

公式我在纸上推导，拍照放在了同一文件夹中，剪绳子.jpg

```go
func cuttingRope_2(n int) int {
   if n <= 3{
      return n-1
   }
   a,b := n/3,n%3
   if b == 0{
      return int(math.Pow(3,float64(a)))
   } else if b == 1{
      return int(math.Pow(3,float64(a-1)) * 4)
   } else {
      return int(math.Pow(3,float64(a)) * 2)
   }
}
```

### 57-II.和为 s 的连续正数序列

> 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
>
> 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
>
> 示例 1：
>
> 输入：target = 9
> 输出：[[2,3,4],[4,5]]
> 示例 2：
>
> 输入：target = 15
> 输出：[[1,2,3,4,5],[4,5,6],[7,8]]
>
>
> 限制：
>
> 1 <= target <= 10^5

解题思路：滑动窗口

设连续正整数序列的左边界与右边界分别为 left 和 right，构建滑动窗口向右滑动。循环中，每轮判断滑动窗口内元素和与 target 的大小关系

- 若相等则记录结果，然后移动左指针
- 若 大于 target 则移动左指针
- 若小于 target 则移动右指针

每次移动指针的同时，更新滑动窗口内元素和 sum。

因为序列至少由两个数字组成，所以左指针边界为 [1,target/2]

```go
func findContinuousSequence(target int) [][]int {
   // left,right 为滑动窗口左右指针
   // sum 动态记录窗口元素和
   // 窗口至少含有两个数
   left, right, sum := 1, 2,3
   res := make([][]int, 0)
   // 序列至少右两个元素组成，所以左边界只需遍历到 target/2
   for left <= target>>1 {
      if sum < target {
         // move right cursor and increase sum
         right++
         sum += right
      } else if sum > target {
         // move left cursor and reduce sum
         sum -= left
         left++
      } else {
         nums := make([]int, 0)
         for i := left; i <= right; i++ {
            nums = append(nums, i)
         }
         res = append(res,nums)
         sum -= left
         left++
      }
   }
   return res
}
```

### 62.圆圈中最后剩下的数字

> 0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
>
> 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
>
> 示例 1：
>
> 输入: n = 5, m = 3
> 输出: 3
> 示例 2：
>
> 输入: n = 10, m = 17
> 输出: 2
>
>
> 限制：
>
> 1 <= n <= 10^5
> 1 <= m <= 10^6

约瑟夫环问题，模拟整个删除过程最直观，这里我用数组进行模拟，首先构建长度为 n 的模拟数组 nums，元素值分别为 0 ~ n-1，start 为每次循环中数组第一个元素下标，初始化为 0，之后开始模拟

先获取当前数组的长度 clen，删除第 m 个数字，m 可能大于等于 n，我们将 m-1 对 n 取余，直接得到要删除的元素下标，考虑到初始元素下标不一定为 0，最终待删除的元素下表 loc_del = (m-1+start) % clen，将删除 该元素后的数组重新赋值给 nums，start 更新为 loc_del，开启下一次循环，直至数组长度为 1，可得到圆圈剩下的最后一个数组，返回该元素即可。

```go
func lastRemaining(n int, m int) int {
   // 每次删除的第一个元素下标
   start := 0
   nums := []int{}
   for i:= 0;i <n;i++{
      nums = append(nums,i)
   }
   for len(nums) > 1{
      clen := len(nums)
      loc_del := (m-1+start) % clen
      nums = append(nums[:loc_del],nums[loc_del+1:]...)
      start = loc_del
   }
   return nums[0]
}
```

提交后会超时，个人感觉时间是浪费在数组的拼接上。

下述解法参考自官方题解

现在我们建模 n 个数字删除 第 m 个元素的问题为 f(n,m)，f(n,m) 的值为最后剩余元素值（同元素下标）

那 f(1,m) 的答案是一定的，数组只存在一个元素，f(1,m) = 0，我们可以试想下，可不可以通过 f(1,n) 推出 f(n,m) 的值。

进一步来说，如果能通过 f(n-1,m) 推出 f(n,m)，那就一定能从 f(1,n) 推出 f(n,m)

因为 f(n,m) 要删除第 m%n 个元素后，长度就变成了 n-1，那自然而然就与 f(n-1,m) 扯上了关联，它们的区别在哪里呢?

对应元素下标不同！ 只要我们把它们的下标相对应起来，那我们就可以求解改题目。

下标不同在哪里呢？起始位置，f(n-1,m) 起始元素下标为 0，而 f(n,m) 删除一个元素后的起始元素下标为 m%n，我们把它们对应起来，就可以通过 f(n-1,m) 求得 f(n,m)

同理，由于 f(1,m) 值是固定的，我们可以从 f(1,m) 递归到 f(n,m)

```go
func lastRemaining_2(n int, m int) int {
	// n=1时，剩余元素下标为0
	res := 0
	// 从前向后推，求n=2,3,...,n时，剩余元素下标
	// i 为数组长度
	for i:= 2;i <= n;i++{
		// 将i-1数组元素下标与长度为i的数组下标对应起来
		// 求得长度为 i 的数组剩余元素
		res = (res + m) % i
	}
	return res
}
```



## day25.模拟（中等）

### 29.顺时针打印矩阵

> 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
>
> 示例：
>
> 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
> 输出：[1,2,3,6,9,8,7,4,5]
>
>
> 限制：
>
> 0 <= matrix.length <= 100
> 0 <= matrix[i].length <= 100
> 注意：本题与主站 54 题（螺旋矩阵）相同

简单模拟，这种题目，个人感觉比较看经验，做过类似题目的话，很快可以模拟出来，没做过或者没思路的话，就有点无从下手或者不知道如何用代码实现自己的想法。

需要考虑三个问题：移动方向、边界 和 结束条件。

1. 移动方向：很明显为 右、下、左、上 这样的循环，我们可以用一个二维数组代表四个方向，每移动到边界，更改方向；
2. 边界：边界问题是本题的重点，因为边界是随着遍历在变化的，打印矩阵的过程中，边界逐渐变小。规则是：如果当前行（列）遍历结束后，将次行（列）的边界向内移动一格；
3. 结束条件：当矩阵所有位置都被打印过，即遍历数组的长度等于矩阵元素个数时，结束遍历。

代码中，x和y代表当前遍历到的元素所在位置索引，dirs代表移动的四个方向（上、右、下、左是一个循环），dir[cur_d]代表下一次移动的方向，结束条件是 res 数组元素个数等于矩阵元素个数。

```go
func spiralOrder(matrix [][]int) []int {
   if len(matrix) == 0 || len(matrix[0]) == 0{
      return []int{}
   }
   m,n := len(matrix),len(matrix[0])
   directions := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
   cur_dir := 0
   top,right,bottom,left := 0,n-1,m-1,0
   res := make([]int,0)
   x,y := 0,0
   for len(res) < m*n{
      res = append(res,matrix[x][y])
      if cur_dir == 0 && y == right{
         cur_dir ++
         top ++
      } else if cur_dir == 1 && x == bottom{
         cur_dir ++
         right --
      } else if cur_dir == 2 && y == left{
         cur_dir ++
         bottom --
      } else if cur_dir == 3 && x == top{
         cur_dir ++
         left ++
      }
      cur_dir %= 4
      x += directions[cur_dir][0]
      y += directions[cur_dir][1]
   }
   return res
}
```

也可以按层模拟，将矩阵分为若干层，先打印最外层元素，再打印次外层的元素，直到输出完最内层的元素。

对于每层，从左上方开始以顺时针的顺序遍历当前层所有元素。还是要着重考虑边界的问题，假设当前层左上角元素索引为(top,left)，右下角元素索引为(bottom,right). 

- 从左到右遍历上侧元素，依次为 (top,left) 到 (top,right)。
- 从上到下遍历右侧元素，依次为 (top+1,right) 到 (bottom,right)。
- 如果 left<right 且 top<bottom，则从右到左遍历下侧元素，依次为 (bottom,right−1) 到 (bottom,left+1)，以及从下到上遍历左侧元素，依次为 (bottom,left) 到 (top+1,left)。

遍历完当前层的元素之后，将 left 和 top 分别增加 1，将 right 和 bottom 分别减少 1，进入下一层继续遍历，直到遍历完所有元素为止。

```go
func spiralOrder(matrix [][]int) []int {
   if len(matrix) == 0 || len(matrix[0]) == 0{
      return []int{}
   }
   m,n := len(matrix),len(matrix[0])
   top,right,bottom,left := 0,n-1,m-1,0
   res := []int{}
   for left <= right && top <= bottom{
      for i := left;i<=right;i++{
         res = append(res,matrix[top][i])
      }
      for i := top+1;i<=bottom;i++{
         res = append(res,matrix[i][right])
      }
      // 这里一定要进行判断，否则会重复收录元素
      if left < right && top < bottom{
         for i := right-1;i>=left;i--{
            res = append(res,matrix[bottom][i])
         }
         for i := bottom-1;i>top;i--{
            res = append(res,matrix[i][left])
         }
      }
      top ++
      right --
      bottom --
      left ++
   }
   return res
}
```

注意：如果循环这么写

```go
for left<=right && top<=bottom{
   for i:=left;i<=right;i++{
      res = append(res,matrix[top][i])
   }
   for i:=top+1;i<=bottom;i++{
      res = append(res,matrix[i][right])
   }
   for i:=right-1;i>left;i--{
      res = append(res,matrix[bottom][i])
   }
   for i:=bottom;i>top;i--{
      res = append(res,matrix[i][left])
   }
   left ++
   right --
   top ++
   bottom --
}
```

输入为 []int{1,2,3} ，那输出会为 []int{1,2,3,2}，所以一定要判断边界的关系。

### 31.栈的压入、弹出序列

> 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
>
> 示例 1：
>
> 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
> 输出：true
> 解释：我们可以按以下顺序执行：
> push(1), push(2), push(3), push(4), pop() -> 4,
> push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
> 示例 2：
>
> 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
> 输出：false
> 解释：1 不能在 2 之前弹出。
>
>
> 提示：
>
> 0 <= pushed.length == popped.length <= 1000
> 0 <= pushed[i], popped[i] < 1000
> pushed 是 popped 的排列。
> 注意：本题与主站 946 题（验证栈序列）相同

简单题，模拟该过程，声明一个栈 stack 和 变量 i（i 符合条件的出栈元素数量），遍历压入序列 pushed，依次入栈，每次入栈结束后，判断栈顶元素是否和出栈序列 popped[i] 相等，若相等，出栈，i++，再次判断栈顶元素和 popped[i]是否相等...直至栈长度为0或者栈顶元素和popped[i]不相等。

遍历完成后，若 i = len(popped)，说明所有元素出栈成功，返回 true。判断条件改为 len(stack)=0 也可以，道理相同。

```go
func validateStackSequences(pushed []int, popped []int) bool {
   stack := []int{}
   i := 0
   for _,num := range(pushed){
      stack = append(stack,num)
      for len(stack) > 0 && stack[len(stack)-1] == popped[i]{
         stack = stack[:len(stack)-1]
         i++
      }
   }
   return len(stack) == 0
}
```



## day26.字符串（中等）

今天这两道题非常考察细节的处理，一定要足够细心

### 20.表示数值的字符串

> 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
>
> 数值（按顺序）可以分成以下几个部分：
>
> 若干空格
> 一个 小数 或者 整数
> （可选）一个 'e' 或 'E' ，后面跟着一个 整数
> 若干空格
> 小数（按顺序）可以分成以下几个部分：
>
> （可选）一个符号字符（'+' 或 '-'）
> 下述格式之一：
> 至少一位数字，后面跟着一个点 '.'
> 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
> 一个点 '.' ，后面跟着至少一位数字
> 整数（按顺序）可以分成以下几个部分：
>
> （可选）一个符号字符（'+' 或 '-'）
> 至少一位数字
> 部分数值列举如下：
>
> ["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
> 部分非数值列举如下：
>
> ["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
>
>
> 示例 1：
>
> 输入：s = "0"
> 输出：true
> 示例 2：
>
> 输入：s = "e"
> 输出：false
> 示例 3：
>
> 输入：s = "."
> 输出：false
> 示例 4：
>
> 输入：s = "    .1  "
> 输出：true
>
>
> 提示：
>
> 1 <= s.length <= 20
> s 仅含英文字母（大写和小写），数字（0-9），加号 '+' ，减号 '-' ，空格 ' ' 或者点 '.' 。

分段进行判断即可，注释写的很清楚，一切尽在代码中：

```go
func isNumber(s string) bool {
   n := len(s)
   // 若字符串长度为0，无法表示数值
   if n == 0{
      return false
   }
   // 从下标0开始遍历
   index := 0
   // 读取字符开头处的若干空格
   for index < n && s[index] == ' '{
      index ++
   }
   // 读取整数部分,，numeric代表读取到index是否为数值
   // index 持续向前推进
   numeric := ScanInteger(s,&index)
   // 若下一个字符为'.'，说明为小数，读取小数部分
   if index < n && s[index] == '.'{
      index ++
      // 这里用逻辑或的原因：
      // 1. '.'前若无数字，则'.'后至少需要一位数字
      // 对应左表达式为 true，右表达式为 false
      // 2.'.'前有数字，则后方可以有也可以没有数字
      // 对应左true右true 和 左false右true
      numeric = ScanUnsignedInteger(s,&index) || numeric
   }
   // 若出现e，后面需要跟一个整数，原理同上
   if index < n && (s[index] == 'e' || s[index] == 'E'){
      index ++
      numeric = numeric && ScanInteger(s,&index)
   }
   // 读取字符串结尾处的若干空格
   for index < n && s[index] == ' '{
      index ++
   }
   // 若 numeric 为 true，且读取到字符串末尾，说明为数值
   return numeric && index == len(s)
}

// 读取字符串s从下标index开始的有符号整数
func ScanInteger(s string,index *int) bool{
   // 先读取正负号，若无符号，默认正数
   if *index < len(s) && (s[*index] == '+' || s[*index] == '-'){
      *index ++
   }
   // 读取符号后，读取无符号整数
   return ScanUnsignedInteger(s,index)
}

// 读取字符串s从下标index开始的无符号整数
// 若未读取到无符号整数，返回false，否则，返回true
func ScanUnsignedInteger(s string,index *int) bool {
   start_of_integer := *index
   // 根据ASCII码判断是否为数字字符
   // 若读取到非数字字符，直接跳出即可
   for *index < len(s) && s[*index] >= '0' && s[*index] <= '9'{
      *index ++
   }
   // 若*index大于读取前的值，说明有整数被读取到
   return *index > start_of_integer
}
```

### 67.把字符串转换成整数

> 写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
>
> 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
>
> 当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
>
> 该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
>
> 注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
>
> 在任何情况下，若函数不能进行有效的转换时，请返回 0。
>
> 说明：
>
> 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
>
> 示例 1:
>
> 输入: "42"
> 输出: 42
> 示例 2:
>
> 输入: "   -42"
> 输出: -42
> 解释: 第一个非空白字符为 '-', 它是一个负号。
>      我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
> 示例 3:
>
> 输入: "4193 with words"
> 输出: 4193
> 解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
> 示例 4:
>
> 输入: "words and 987"
> 输出: 0
> 解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
>      因此无法执行有效的转换。
> 示例 5:
>
> 输入: "-91283472332"
> 输出: -2147483648
> 解释: 数字 "-91283472332" 超过 32 位有符号整数范围。 
>      因此返回 INT_MIN (−2^31) 。

这里我想先解释一下为什么负数范围比正数范围大 1 .

因为对有符号数，在计算机中表达时，最高位约定为符号位，符号位为 0 时代表正数，符号位为 1 时代表负数。

对 int32，真正表达数值的部分就是剩下的 31 位，正数范围很容易理解为 1~2^31-1（2147483647）

有一个特殊的值，就是 0 值，对符号位 1 和 0 的时候，都是 0，用两个形式表达同一个数无疑是浪费的，所以，我们就规定，符号位为 1 时的全 0，表达 -2147483648，这也是负数表示的范围比正数多1的原因。

综上，对于任意位的，无论是8位，16位，32位甚至64位的整数类型表示范围的计算公式为：

如总位数为n位，那么有符号数的范围为：-2^(n-1) ~ 2^(n-1)-1，无符号数的表示范围为：0~2^n-1



此题解参考自：[剑指 Offer 67. 把字符串转换成整数 - 力扣（LeetCode）](https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/submissions/)

首部空格： 删除之即可；
符号位： 三种情况，即 ''+'' , ''-'' , ''无符号" ；新建一个变量保存符号位，返回前判断正负即可。
非数字字符： 遇到首个非数字的字符时，应立即返回。
数字字符：
字符转数字： “此数字的 ASCII 码” 与 “ 00 的 ASCII 码” 相减即可；
数字拼接： 若从左向右遍历数字，设当前位字符为 c ，当前位数字为 x ，数字结果为 res，则数字拼接公式为：res = 10 * res + x

x = ascii(c) - ascii('0')

接下来是比较关键的数字越界处理：

题目要求返回的数值范围应在 [-2^{31}, 2^{31} - 1]，因此需要考虑数字越界问题。而由于题目指出 环境只能存储 32 位大小的有符号整数 ，因此判断数字越界时，要始终保持 res 在 int 类型的取值范围内。

在每轮数字拼接前，判断 res 在此轮拼接后是否超过 2147483647，若超过则加上符号位直接返回。

设数字拼接边界 bndry = 2147483647 // 10 = 214748364，则以下两种情况越界：

1. res > boundry， 拼接后越界
2. res = boundry && x > 7，拼接后越界

这种处理方式非常巧妙，> ‘7’ 这个操作看似只考虑到了最大值，其实也考虑了最小值，只要大于，直接返回即可。

```go
const (
   MAX_INT = 1 << 31 - 1
   MIN_INT = -1 * (1 << 31)
)

func strToInt(str string) int {
   // index 为 str 遍历到的位置
   index := 0
   // 为防止读取空格出现越界panic，需要提前处理str为空字符或全空格的情况
   if str == ""{
      return 0
   }
   for i:=0;i<len(str);i++{
      if str[i] != ' '{
         break
      }
      if i == len(str) - 1{
         return 0
      }
   }
   // 读取前置空格
   for str[index] == ' '{
      index ++
   }
   // 符号位初始化为 0
   sign := 1
   if str[index] == '+'{
      index ++
   } else if str[index] == '-' {
      sign = -1
      index ++
   }
   boundry := MAX_INT/10
   res := 0
   for _,s := range str[index:]{
      // 读取到非数字字符，直接跳出
      if s < '0' || s > '9'{
         break
      }
      // 预计出现越界情况，根据符号位返回最大或最小值
      if res > boundry || res == boundry && s > '7'{
         if sign == 1{
            return MAX_INT
         } else {
            return MIN_INT
         }
      }
      res = res * 10 + int(s - '0')

   }
   return sign * res
}
```



## day27.栈与队列（困难）

### 59-I.滑动窗口的最大值

> 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
>
> 示例:
>
> 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
> 输出: [3,3,5,5,6,7] 
> 解释: 
>
>   滑动窗口的位置                最大值
> ---------------               -----
> [1  3  -1] -3  5  3  6  7       3
>  1 [3  -1  -3] 5  3  6  7       3
>  1  3 [-1  -3  5] 3  6  7       5
>  1  3  -1 [-3  5  3] 6  7       5
>  1  3  -1  -3 [5  3  6] 7       6
>  1  3  -1  -3  5 [3  6  7]      7
>
>
> 提示：
>
> 你可以假设 k 总是有效的，在输入数组 不为空 的情况下，1 ≤ k ≤ nums.length。
>

最简单的，可以先去暴力解决，对每个滑动窗口，遍历所有元素得到最大值。对长度为 n 的数组 nums，窗口数量为 n-k+1，暴力解法时间复杂度为 O((n-k+1)*k) = O(nk)

```go
func maxSlidingWindow(nums []int, k int) []int {
   n := len(nums)
   if n == 0 || k == 0{
      return []int{}
   }
   res := []int{}
   for i := 0;i < n-k+1;i++{
      res = append(res,max(nums[i:i+k]))
   }
   return res
}


func max(sli []int) (res int){
   res = sli[0]
   for _,i := range sli{
      if i > res{
         res = i
      }
   }
   return
}
```

然后想一想，暴力没有用到题目中哪些条件，可以针对这些条件去做改进

相邻的两个窗口，有重叠元素 k-1 个，这 k-1 个元素中的最大值对两个窗口来说是相同的，这就是我们改进的方向。

设想一下，一个窗口中两个元素下标分别为 i 和 j，其中 i 在 j 的左侧（i<j)

- 若 nums[i] < nums[j]，**在滑动窗口右移的过程中，只要 nums[i] 在窗口中，则 nums[j] 也一定还在窗口中，这是 i<j 所保证的。因此，由于 nums[j] 的存在，nums[i] 一定不会是滑动窗口最大值，我们可以将 nums[i] 永久删除。**
- 若 nums[i] >= nums[j]，nums[i] 出队后，nums[j] 是有可能成为队列最大值的，nums[j] 需要保留

因此，我们可以使用一个队列来存储还没有被删除的元素。在队列中，这些元素的值是单调递减的。当窗口向右滑动时，我们需要把一个新的元素放入队列中，为了保持队列的性质，我们需要不断将新元素与队列队尾元素进行比较，若新元素较大，则队尾元素出队，不断进行此操作，直至队列为空 或 新的元素小于等于队尾的元素。

由于队列的元素是严格单调递减的，且队列中元素属于该窗口，所以队首元素就是该窗口的最大值。

此时，我们需要考虑最后一个问题，若当前窗口的最大值为窗口最左侧元素，那进入下一个窗口前队列中该元素应该出队，因为该元素并不属于下一个窗口。

该队列元素单调递减，满足这种单调性的队列一般称作 单调队列。

```go
func maxSlidingWindow_2(nums []int, k int) []int {
   n := len(nums)
   monoQ := []int{}
   res := make([]int,0,n-k+1)
   for i:=0;i<n;i++{
      for len(monoQ) > 0 && nums[i] > monoQ[len(monoQ)-1]{
         monoQ = monoQ[:len(monoQ)-1]
      }
      monoQ = append(monoQ,nums[i])
      // 若单调队列最大值非窗口元素
      // 将该元素出队
      if i >= k && monoQ[0] == nums[i-k]{
         monoQ = monoQ[1:]
      }
      // i=k-1时，窗口元素数量达到 k
      // 开始向 res 数组添加窗口最大值
      if i >= k-1{
         res = append(res,monoQ[0])
      }
   }
   return res
}
```

### 59-II.队列的最大值

> 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
>
> 若队列为空，pop_front 和 max_value 需要返回 -1
>
> 示例：
>
> 输入: 
> ["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
> [[],[1],[2],[],[],[]]
> 输出: [null,null,null,2,1,2]
>
> 1 <= push_back,pop_front,max_value的总操作数 <= 10000
> 1 <= value <= 10^5

有了上一题的基础后，这一题就简简单单了

和第一题的解析类似，这里我们也可以使用单调队列。把队列看作是第一题的窗口，只不过滑动规则有所差异，pop 出队时，窗口的左指针右移，并对我们实现的单调队列进行相应操作，求队列最大值时，只需返回单调队列队首元素即可。

```go
type MaxQueue struct {
   // 存储队列
   q     []int
   // 单调队列
   monoQ []int
}

// 构造函数
func Constructor() MaxQueue {
   return MaxQueue{[]int{},[]int{}}
}

// 取最大值，若队列长度为 0，返回-1，否则返回单调队列队首元素
func (this *MaxQueue) Max_value() int {
   if len(this.q) == 0{
      return -1
   }
   return this.monoQ[0]
}

// 入队，存储队列直接入队
// 单调队列需要判断队尾元素与新元素大小关系
// 若新元素大于队尾元素，出队
// 直至单调队列长度为 0或者 新元素值 小于等于 队尾元素
func (this *MaxQueue) Push_back(value int)  {
   this.q = append(this.q,value)
   for len(this.monoQ)>0 && value>this.monoQ[len(this.monoQ)-1]{
      this.monoQ = this.monoQ[:len(this.monoQ)-1]
   }
   this.monoQ = append(this.monoQ,value)
}

// 出队，先判断存储队列长度，若为 0，返回-1
// 获取存储队列队首元素后，存储队列队首元素出队
// 若该元素等于单调队列队首元素，则单调队列队首元素出队
func (this *MaxQueue) Pop_front() int {
   if len(this.q) == 0{
      return -1
   }
   res := this.q[0]
   this.q = this.q[1:]
   if res == this.monoQ[0]{
      this.monoQ = this.monoQ[1:]
   }
   return res
}
```



## day28.搜索与回溯算法（困难）

### 37.序列化二叉树

> 请实现两个函数，分别用来序列化和反序列化二叉树。
>
> 你需要设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
>
> **示例：**
>
> 输入：root = [1,2,3,null,null,4,5]
> 输出：[1,2,3,null,null,4,5]

本题接参考自官方题解：[二叉树的序列化与反序列化 - 二叉树的序列化与反序列化 - 力扣（LeetCode）](https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/solution/er-cha-shu-de-xu-lie-hua-yu-fan-xu-lie-hua-by-le-2/)

**二叉树的序列化本质上是对其值进行编码，更重要的是对其结构进行编码。可以遍历树来完成上述任务。众所周知，我们一般有两个策略：广度优先搜索和深度优先搜索。**

广度优先搜索可以按照层次的顺序从上到下遍历所有的节点

深度优先搜索可以从一个根开始，一直延伸到某个叶，然后回到根，到达另一个分支。根据根节点、左节点和右节点之间的相对顺序，可以进一步将深度优先搜索策略区分为：先序遍历、中序遍历 以及 后序遍历。

对题目所给例子进行举例，最终序列化字符串是`1,2,3,None,None,4,None,None,5,None,None,`。其中，None，None 用来标记缺少左右子节点，这是我们在序列化期间保存树结构的方式。

即我们可以先序遍历这颗二叉树，遇到空子树的时候序列化成 None，否则继续递归序列化。那么我们如何反序列化呢？首先我们需要根据 , 把原先的序列分割开来得到先序遍历的元素列表，然后从左向右遍历这个序列：

- 如果当前的元素为 None，则当前为空树
- 否则先解析这棵树的左子树，再解析它的右子树

```go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type Codec struct {

}

func Constructor() Codec {
   return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
   var sb strings.Builder
   var preorder func(root *TreeNode)
   preorder = func(root *TreeNode) {
      if root == nil{
         sb.WriteString("nil,")
         return
      }
      sb.WriteString(strconv.Itoa(root.Val))
      sb.WriteString(",")
      preorder(root.Left)
      preorder(root.Right)
   }
   preorder(root)
   return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
   s := strings.Split(data,",")
   var build func() *TreeNode
   build = func() *TreeNode {
      if s[0] == "nil"{
         s = s[1:]
         return nil
      }
      val,_ := strconv.Atoi(s[0])
      s = s[1:]
      node := &TreeNode{val,build(),build()}
      return node
   }
   return build()
}
```

### 38.字符串的排列

> 输入一个字符串，打印出该字符串中字符的所有排列。
>
> 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
>
> 示例:
>
> 输入：s = "abc"
> 输出：["abc","acb","bac","bca","cab","cba"]
>
>
> 限制：
>
> 1 <= s 的长度 <= 8

全排列问题，如果题目所给字符不包含重复字符，用下面简单回溯即可：

```go
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
```

本题的难点在于不能有重复元素，这也就意味着输入字符串中包含重复元素。上面的代码不再可行

用集合去重是可行的，但太过笨重，试想是否可以改善回溯过程，得到不含重复元素的排列。

去重，很容易想到要先对元素进行排序，这样才能判断相邻的节点是否重复使用了。

回溯的问题在纸上都可以画成一个树的结构，横向迭代+纵向递归。

针对本题，画出该树结构，可以发现同一树层上，若当前节点与上一个节点值相同，则对该节点回溯的结果与上一节点相同，意味着可以跳过此节点，迭代至下一节点。

那同一个树枝呢？同一树枝上，若上一个节点使用过，该节点还是可以使用的，因为这是一个正常的排列。

接下来就是代码实现，当前元素与上一个元素值相同的情况下，判断是在同一树枝，还是同一树层。

used[i-1]=true 说明上一节点已使用，说明是同一树枝，

used[i]=false，说明为同一树层，可跳过此次回溯。

```go
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
```

本题解题思路与 LeetCode 48：全排列 II 完全一致，有兴趣可以再去做做练习一下



## day29.动态规划（困难）

### 19.正则表达式匹配

> 请实现一个函数用来匹配包含'. '和'\*' 的正则表达式。模式中的字符'.'表示任意一个字符，而'\*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"abaca"匹配，但与"aa.a"和"ab*a"均不匹配。
>
> 示例 1:
>
> 输入:
> s = "aa"
> p = "a"
> 输出: false
> 解释: "a" 无法匹配 "aa" 整个字符串。
> 示例 2:
>
> 输入:
> s = "aa"
> p = "a\*"
> 输出: true
> 解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
> 示例 3:
>
> 输入:
> s = "ab"
> p = ".*"
> 输出: true
> 解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
> 示例 4:
>
> 输入:
> s = "aab"
> p = "c\*a\*b"
> 输出: true
> 解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
> 示例 5:
>
> 输入:
> s = "mississippi"
> p = "mis\*is\*p\*."
> 输出: false
> s 可能为空，且只包含从 a-z 的小写字母。
> p 可能为空，且只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'。

本题解参考自LeetCode官方题解。

题目中的匹配是一个「逐步匹配」的过程：我们每次从字符串 p 中取出一个字符或者「字符 + 星号」的组合，并在 s 中进行匹配。对于 p 中一个字符而言，它只能在 s 中匹配一个字符，匹配的方法具有唯一性；而对于 p 中字符 + 星号的组合而言，它可以在 s 中匹配任意自然数个字符，并不具有唯一性。因此我们可以考虑使用动态规划，对匹配的方案进行枚举。

我们用 dp\[i\]\[j\] 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配。在进行状态转移时，我们考虑 p 的第 j 个字符的匹配情况：

如果 p 的第 j 个字符是一个小写字母，那么我们必须在 s 中匹配一个相同的小写字母

如果 s 的第 i 个字符与 p 的第 j 个字符不相同，那么无法进行匹配；否则我们可以匹配两个字符串的最后一个字符，完整的匹配结果取决于两个字符串前面的部分。

dp\[i\]\[j\] =dp\[i-1\]\[j-1\]

若 p 的第 j 个字符为 ‘*’，我们可以对 p 的第 j-1 个字符匹配任意次，匹配 0 次的情况下，有 dp\[i\]\[j\] = dp\[i\]\[j-2\]，在匹配 s中字符 1、2 次的情况下，有 dp\[i\]\[j\] = dp\[i-1\]\[j-2\]、dp\[i\]\[j\]= dp\[i-2\]\[j-2\]，如果用这种方式进行转移，那么我们就需要枚举这个组合到底匹配了 s 中的几个字符，会增导致时间复杂度增加，并且代码编写起来十分麻烦。我们不妨换个角度考虑这个问题：字母 + 星号的组合在匹配的过程中，本质上只会有两种情况：

- 匹配 s 末尾的一个字符，将该字符扔掉，而该组合还可以继续进行匹配；
- 不匹配字符，将该组合扔掉，不再进行匹配。

最终的状态转移方程如下：

- if p[j]!='*' && match(i,j)，dp\[i\]\[j\] = dp\[i-1\]\[j-1\]
- if p[j]!='*' && !match(i,j)，dp\[i\]\[j\] = false
- if p[j] == '*' && match(i,j-1)，dp\[i\]\[j\] = dp\[i\]\[j-2\]
- if p[j] == '*' && !match(i,j-1)，dp\[i\]\[j\] = dp\[i-1\]\[j\] || dp\[i\]\[j-2\]

其中，match(i,j) 是判断 s[i-1] 与 p[j-1] 是否匹配的辅助函数.

字符串的字符下标是从 0 开始的，因此在实现上面的状态转移方程时，需要注意状态中每一维下标与实际字符下标的对应关系。

```go
func isMatch(s string, p string) bool {
   m,n := len(s),len(p)
   dp := make([][]bool,m+1)
   for i:=0;i<=m;i++{
      dp[i] = make([]bool,n+1)
   }
   // 状态初始化，模式和正则表达式均为空时，匹配成功
   dp[0][0] = true
   // match 用于匹配单个字符
   // s[i-1] 与 p[j-1] 是否匹配成功
   match := func(i,j int) bool {
      if i == 0{
         return false
      }
      if p[j-1] == '.'{
         return true
      }
      return s[i-1] == p[j-1]
   }
   // i 从 0，j 从 1 开始遍历
   // 因为当s不为空，而p为空的时候，匹配一定失败
   for i:=0;i<=m;i++{
      for j:=1;j<=n;j++{
         // 对应正则为 字符+* 组合
         if p[j-1] == '*'{
            dp[i][j] = dp[i][j-2]
            if match(i,j-1){
               dp[i][j] = dp[i][j] || dp[i-1][j]
            }
         } else {
            // 对应 正则为单个字符 的情况
            if match(i,j){
               dp[i][j] = dp[i-1][j-1]
            }
         }
      }
   }
   return dp[m][n]
}
```

### 49.丑数

> 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
>
>  示例:
>
> 输入: n = 10
> 输出: 12
> 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
> 说明:  
>
> 1 是丑数。
> n 不超过1690。

丑数的递推性质： 丑数只包含因子 2, 3, 5，因此有 “丑数 = 某较小丑数 * 某因子” （因子为 2、3、5）。

该性质是解题的关键，刚开始我们只拥有丑数 1，然后 经过 1 与因子 2、3、5 相乘，得到三个丑数，将该三个丑数再次与因子相乘，又得到部分丑数（不一定为 9个，因为可能有重复出现的数字），我们不断经过此步骤，可以得到很多丑数

那如何得到第 n 个呢？上面得到的丑数是无序的，我们可以使用动态规划得到排好序的丑数，丑数数字 dp 初始化时只有一个元素 1，三个因子的指针 p2、p3、p5 均指向 1，计算得到 三个因子指向的丑数 与 该因子的 乘积，取最小值，就是下一个丑数，然后找到对应的指针，将该指针自增1，指向下一个丑数，每次可以得到一个丑数，n-1次循环可得到第 n 个丑数。

动态规划三步：

1. 定义dp数组大小及下表含义：dp[i] 代表第 i 个丑数
2. dp 数组状态初始化：dp[1] = 1，三个指针 p2,p3,p5=1,1,1
3. 状态转移方程，dp[i] = min(dp[p2]*2,dp[p3]\*3,dp[p5]\*5)，之后找到对应的指针，将该指针自增1

```go
func nthUglyNumber(n int) int {
   dp := make([]int,n+1)
   // dp 数组初始化，只有 1 一个丑数
   dp[1] = 1
   // 三个指针初始化指向第一个丑数
   p2,p3,p5:= 1,1,1
   for i:=2;i<=n;i++{
      // 寻找三个指针指向元素与对应因子乘积的最小值
      num := min(dp[p2]*2,min(dp[p3]*3,dp[p5]*5))
      dp[i] = num
	  // 找到对应指针，该指针右移（即自增1）
	  // 可能对应不止一个指针
      if num == dp[p2] * 2{
         p2++
      }
      if num == dp[p3] * 3{
         p3++
      }
      if num == dp[p5] * 5{
         p5++
      }
   }
   return dp[n]
}
```

此题我们还可以用小根堆来解决

初始时堆为中只有第一个丑数 1 。

每次取出堆顶元素 x，则 x 是堆中最小的丑数，由于 2x, 3x, 5x 也是丑数，因此将 2x, 3x, 5x 加入堆。

上述做法会导致堆中出现重复元素的情况。为了避免重复元素，可以使用哈希集合去重，避免相同元素多次加入堆。

在排除重复元素的情况下，第 n 次从最小堆中取出的元素即为第 n 个丑数。

```go
type maxH []int

func (this maxH) Len() int{ return len(this)}
func (this maxH) Less (i,j int) bool {
   return this[i] < this[j]
}
func (this maxH) Swap (i,j int){
   this[i],this[j] = this[j],this[i]
}
func (this *maxH) Push(v interface{}){
   *this = append(*this,v.(int))
}
func (this *maxH) Pop() interface{}{
   old := *this
   n := len(old)
   res := old[n-1]
   *this = old[:n-1]
   return res
}


func nthUglyNumber(n int) int {
   uglys := &maxH{1}
   factors := []int{1,3,5}
   record := map[int]struct{}{}
   for i:=1;;i++{
      num := heap.Pop(uglys).(int)
      if i == n{
         return num
      }
      for _,f := range factors{
         next := num * f
         if _,has := record[next];!has{
            heap.Push(uglys,next)
            record[next] = struct{}{}
         }
      }
   }
}
```

### 60.n个骰子的点数

> 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
>
> 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
>
> 示例 1:
>
> 输入: 1
> 输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
> 示例 2:
>
> 输入: 2
> 输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]
>
>
> 限制：
>
> 1 <= n <= 11

这道题我的第一反应是做排列组合问题。看题解才知道是 dp 问题。

1 个骰子共有 6 中可能的点数，两个骰子点数范围为 [2,12]，共有 11 种可能的点数，对 n 个骰子，点数范围为 [n,6*n]，共有 5 \* n + 1 种可能的点数。

如果我们知道了 x-1 个骰子所有的点数及其概率，那 x 个骰子的点数及其概率可以通过其推出。对 x 个骰子，也就是在 x-1 的基础上添加一个骰子，该骰子的点数可为 1-6，每个点数的概率均为 1/6，那我们可以对 x-1 个骰子的点数进行遍历，再对 1-6 开启第二层遍历，设当前遍历到 x-1 个骰子的点数为 m，第 x 个骰子遍历到的点数为 n，那么 dp\[x\]\[m+n\] += dp\[x-1\]\[m\]/6

1. 确定dp数组及下标含义：dp\[i\]\[j\] 代表 i 个骰子 点数为 j 的概率，则 len(dp) = n+1, len(dp[0]) = 6\*n+1
2. 数组初始化：一个骰子的点数及其概率我们是已知的，dp\[1\[1\] - dp\[1\]\[6\] 的值均为 1/6，dp 数组中其余值初始化为 0 即可，其余有效信息将由 1 个骰子的点数及其概率得出。
3. 状态转移方程

$$
dp[n][x]=\sum_{i=1}^6dp[n-1][x-i]*1/6
$$

可以看出，该状态转移方程是一个理论上合理的方案，但用此公式求解的话，需要考虑越界的问题，是一种逆向的递推公式，我们将其改为正向即可，通过 x-1 个骰子的点数推导 x 个骰子的点数。

```Go
func dicesProbability(n int) []float64 {
   // dp数组初始化
   dp := make([][]float64,n+1)
   for i:=0;i<n+1;i++{
      dp[i] = make([]float64,6*n+1)
   }
   for i:=1;i<=6;i++{
      dp[1][i] = float64(1)/float64(6)
   }
   // 已知1个骰子的点数及其概率，向后推导至n个骰子的点数
   for i:=2;i<=n;i++{
      // 通过 i-1 个骰子的点数，求第 i 个骰子的点数
      for j:=i-1;j<=6*(i-1);j++{
         // 每个骰子的点数为1-6，概率均为1/6
         for k:=1;k<=6;k++{
            dp[i][j+k] += (dp[i-1][j])/6
         }
      }
   }
   return dp[n][n:]
}
```

另外，感觉这道题是 dp 思路很好的练习题目

一个阶段有个多个状态，且当前阶段的每个状态会影响到下一个阶段的多个状态（至多6个）



## day30.分治算法（困难）

### 17.打印从1到最大的n位数

> 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
>
> 示例 1:
>
> 输入: n = 1
> 输出: [1,2,3,4,5,6,7,8,9]
>
>
> 说明：
>
> 用返回一个整数列表来代替打印
> n 为正整数

题目描述为返回整数列表，那通过不断迭代，从1，不断加至最大的n位数，依次添加至结果列表，返回即可。

```go
func printNumbers(n int) []int {
   res := []int{}
   limit := math.Pow(10,float64(n))
   x := 1
   for x < int(limit){
      res = append(res,x)
      x += 1
   }
   return res
}
```

在《剑指offer》中，需要打印字符串数组，因为题目没有规定 n 的范围，当 n 非常大的时候，最大的 n 位数会超出整型的范围，所以说，一定要考虑到**大数问题**，这才是本题的考点。大数问题的解决办法就是用字符串表达大数。

最直观的解法是用字符串模拟数字加法，另外，也容易想到，返回的字符串数组为 n 个 0-9 的全排列。全排列就要用到递归。

基于分治的思想，先确定高位，依次向低位移动。例如 n=2 时，先确定十位，再确定个位。

最后需要处理每个字符串的前置零，数组的第一个元素为全零字符串，也需要删除，代码注释很详细了，如下：

```go
func printNumbers(n int) []int {
   // 先按照返回字符串数组的思路进行解题
   // strs 存储回溯得到的字符串数组（未处理前置零）
   strs := []string{}
   var backtrace func(path string)
   backtrace = func(path string){
      // 字符串长度达到n时，将字符串添加至strs
      // 结束递归
      if len(path) == n {
         strs = append(strs,path)
         return
      }
      // 从 0-9 迭代
      for i:=0;i<10;i++{
         // 将当前数字的字符添加至 path
         path += strconv.Itoa(i)
         backtrace(path)
         // 回溯，删除最新添加的字符
         path = path[:len(path)-1]
      }
   }
   backtrace("")
   // 删除数组第一个全零字符串
   strs = strs[1:]
   k := len(strs)
   // delZeroStrs 用来保存去除前置零的字符串数组
   delZeroStrs := make([]string,k)
   for i:=0;i<k;i++{
      delZeroStrs[i] = delZero(strs[i])
   }
   // 最后，为符合本题要求，转化为整数数组
   res := make([]int,k)
   for i:=0;i<k;i++{
      num,_ := strconv.Atoi(delZeroStrs[i])
      res[i] = num
   }
   return res

}

// delZero 用于删除输入字符串 s 前置零
func delZero (s string) string{
   for i:=0;i<len(s);i++{
      if s[i] != '0'{
         return s[i:]
      }
   }
   return s
}
```

下述内容是关于回溯写法相关的

对 printNumbers 函数中的 backtrace 函数，我们输入参数为 path，调用时传入空字符串，然后回看代码的时候发现，提前声明一个 path 变量，为空字符串，然后回溯函数好像就不需要参数列表了

做了尝试，发现完全没问题

```go
func printNumbers_2(n int) []int {
   strs := []string{}
   path := ""
   var backtrace func()
   backtrace = func(){
      if len(path) == n {
         strs = append(strs,path)
         return
      }
      for i:=0;i<10;i++{
         path += strconv.Itoa(i)
         backtrace()
         path = path[:len(path)-1]
      }
   }
   backtrace()
   strs = strs[1:]
   k := len(strs)
   delZeroStrs := make([]string,k)
   for i:=0;i<k;i++{
      delZeroStrs[i] = delZero(strs[i])
   }
   res := make([]int,k)
   for i:=0;i<k;i++{
      num,_ := strconv.Atoi(delZeroStrs[i])
      res[i] = num
   }
   return res
}
```

另外，如果我们的 path 是一个 string 切片的话，我们就不再需要回溯操作，每次迭代操作会将上一次迭代的字符覆盖掉，变成一个 dfs 。

这段代码是我挺久之前写的，注释就不修改了，可以当参考看一下

```go
func printNumbers(n int) []int {
   // 先按照返回字符串数字思路进行解题
   ans := []string{}
   var dfs func(x int)
   // temp存储当前字符串的字节切片（代码用字符串切片存储）
   temp := make([]string,n)
   dfs = func(x int) {
      if x == n{
         // 达到n位时，递归结束，
         // 组合temp添加至ans切片末尾
         ans = append(ans,strings.Join(temp,""))
         return
      }
      // 每一位从0-9进行递归
      for i:=0;i<10;i++{
         temp[x] = strconv.Itoa(i)
         // 当前位处理完成后，进行下一位递归处理
         dfs(x+1)
      }
   }
   // 从第0位开始递归
   dfs(0)
   // 去除每个字符串的前置零
   for i:=0;i<len(ans);i++{
      ans[i] = delzero(ans[i])
   }
   // 去除第一个字符串，该字符串全零
   ans = ans[1:]
   // 最后，为符合本题要求，转化为整数切片
   res := make([]int,len(ans))
   for i:=0;i<len(ans);i++{
      res[i],_ = strconv.Atoi(ans[i])
   }
   return res
}
// 去除字符串高位 0
func delzero(s string) string {
   start := 0
   for i:=0;i<len(s);i++{
      if s[i]=='0'{
         start++
      } else {
         break
      }
   }
   return s[start:]
}
```

我还应该学一下字符串模拟加法[在字符串上模拟数字加法, 解决大数字溢出的问题 - 打印从1到最大的n位数 - 力扣（LeetCode）](https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/solution/zai-zi-fu-chuan-shang-mo-ni-shu-zi-jia-f-kmv1/)

### 51.数组中的逆序对

> 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
>
> 示例 1:
>
> 输入: [7,5,6,4]
> 输出: 5
>
>
> 限制：
>
> 0 <= 数组长度 <= 50000

这道题很容易想到用双层循环，暴力地去得到结果

设数组长度为 n，

第一层的循环次数为 n-1，分别用于求数组对应下标元素的“贡献度”，因为最后一个元素值无论是多少，其后方已经没有元素了，贡献度一定为 0，

第二次循环从第一次循环的下一个元素开始，遍历到数组最后一个元素。

这样，计算出前 n-1 个元素的贡献度，相加即可得到逆序对的总数。

```go
func reversePairs(nums []int) int {
   n := len(nums)
   res := 0
   for i:=0;i<n-1;i++{
      for j:=i+1;j<n;j++{
         if nums[i] > nums[j]{
            res ++
         }
      }
   }
   return res
}
```

提交后发现超时，需要想其他的方法

参考自LeetCode官方题解

相比第一题，我认为这道题更体现出分治的思想。解题的前提是理解归并排序，归并排序的核心在于合并两个有序序列时，我们只需要 O(n) 的时间复杂度，当序列长度为 1 时，自然有序，然后向上回溯。

「归并排序」是分治思想的典型应用，它包含这样三个步骤：

- 分解： 待排序的区间为 [l, r]，令 m=(l+r)/2 ，我们把 [l, r] 分成[l,m] 和[m+1,r]
- 解决： 使用归并排序递归地排序两个子序列
- 合并： 把两个已经排好序的子序列 [l,m] 和 [m+1,r] 合并起来

在待排序序列长度为 1 的时候，递归开始「回升」，因为我们默认长度为 1 的序列是排好序的。

那么求逆序对和归并排序又有什么关系呢？关键就在于「归并」当中「并」的过程。我们通过一个实例来看看。假设我们有两个已排序的序列等待合并，分别是 L= { 8,12,16,22,100 } 和 R = \{ 9, 26, 55, 64, 91 \}。一开始我们用指针 lPtr = 0 指向 L 的首部，rPtr = 0 指向 R 的头部。记已经合并好的部分为 M。

L = [8, 12, 16, 22, 100]   R = [9, 26, 55, 64, 91]  M = []
我们发现 lPtr 指向的元素 8 小于 rPtr 指向的元素 9，于是把 lPtr 指向的元素放入 M，并把 lPtr 后移一位。
这个时候我们把左边的 8 加入了 M，我们发现右边没有数比 8 小，所以 8 对逆序对总数的「贡献」为 0。

接着我们继续合并，把 9 加入了答案，此时 lPtr 指向 12，rPtr 指向 26。
此时 lPtr 比 rPtr 小，把 lPtr 对应的数加入答案，并考虑它对逆序对总数的贡献为 rPtr 相对 R 首位置的偏移 1（即右边只有一个数比 12 小，所以只有它和 12 构成逆序对），以此类推。

我们发现用这种「算贡献」的思想在合并的过程中计算逆序对的数量的时候，只在 lPtr 右移的时候计算，是基于这样的事实：当前 lPtr 指向的数字比 rPtr 小，但是比 R 中 [0 ... rPtr - 1] 的其他数字大，所以这里就贡献了 rPtr 个逆序对。

以上就是官方题解的内容。

我想再补充一下自己的理解：

设数组 nums 长度 为 n，左半序列 left=nums[:n/2]，右半序列 right=nums[n/2:]，两者都是有序的，

那在当前数组中，右半序列的所有元素贡献度均为 0，因为其右方数组均大于等于自身

在左半序列中，每个元素相对于左半序列来说贡献度为 0，但我们要求的是针对数组 nums 每个元素的贡献度，左半序列每个元素大于右半序列元素的个数就是其贡献度，就可以合并两个序列，每次左指针移动的时候，得到左半序列某个元素的贡献度。

很巧妙的分治思想

```go
func reversePairs(nums []int) int {
   return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
   if start >= end {
      return 0
   }
   mid := start + (end - start)/2
   cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid + 1, end)
   tmp := []int{}
   i, j := start, mid + 1
   for i <= mid && j <= end {
      if nums[i] <= nums[j] {
         tmp = append(tmp, nums[i])
         cnt += j - (mid + 1)
         i++
      } else {
         tmp = append(tmp, nums[j])
         j++
      }
   }
   for ; i <= mid; i++ {
      tmp = append(tmp, nums[i])
      cnt += end - (mid + 1) + 1
   }
   for ; j <= end; j++ {
      tmp = append(tmp, nums[j])
   }
   for i := start; i <= end; i++ {
      nums[i] = tmp[i - start]
   }
   return cnt
}
```

然后，我想着，能够在在常见的归并排序写法中，通过加入一个整型指针，归并时修改该指针的值，来达到统计逆序对的目的呢

尝试后，发现是可以的，占位符 _ 代表排序后的数组，因为用不到，所以使用了占位符。

需要特别注意注意一种情况，就是左序列的指针指向的元素值和右序列指向的元素值相等时，如何处理，在常规的归并排序中，在这种情况下，无论我们先处理左指针指向的元素，还是右指针指向的元素，都是可以的。

但在这里不行，我们必须优先移动左指针，因为相同的情况下，左指针指向的元素已经和右指针指向的元素无法形成逆序对。这点需要特别注意

```go
func reversePairs_3(nums []int) int {
   res := 0
   _ = merge_sort(nums,&res)
   return res
}

func merge_sort(nums []int,cnt *int) []int {
   n := len(nums)
   if n < 2{
      return nums
   }
   mid := n/2
   left_part := merge_sort(nums[:mid],cnt)
   right_part := merge_sort(nums[mid:],cnt)
   result := merge(left_part,right_part,cnt)
   return result
}

func merge(left_part,right_part []int,cnt *int) []int {
   result := make([]int,0)
   i,j := 0,0
   for i<len(left_part) && j<len(right_part){
      // 重点在这里 需要=
      // 不然 [1,3,2,3,1]过不了
      if left_part[i] <= right_part[j]{
         result = append(result,left_part[i])
         *cnt = *cnt + j
         i++
      } else {
         result = append(result,right_part[j])
         j++
      }
   }
   for k:=0;k<len(left_part)-i;k++{
      *cnt += len(right_part)
   }
   result = append(result,left_part[i:]...)
   result = append(result,right_part[j:]...)
   return result
}
```

## day31.数学（困难）

### 14-II.剪绳子 II

> 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
>
> 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
>
> 示例 1：
>
> 输入: 2
> 输出: 1
> 解释: 2 = 1 + 1, 1 × 1 = 1
> 示例 2:
>
> 输入: 10
> 输出: 36
> 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
>
>
> 提示：
>
> 2 <= n <= 1000

本题与 day24-剪绳子 题目区别仅在于答案需要取模，解题思路不做复述了

我刚开始的想法是，对运行结果进行一次取余就可以，代码如下

```go
import "math"

func cuttingRope(n int) int {
   mod := math.Pow(10,9) + 7
   if n <= 3{
      return n-1
   }
   a,b := n/3,n%3
   res := float64(0)
   if b == 0{
      res = math.Pow(3,float64(a))
   } else if b == 1{
      res = math.Pow(3,float64(a-1)) * 4
   } else {
      res = math.Pow(3,float64(a)) * 2
   }
   return int(res) % int(mod)
}
```

提交之后发现结果错误，我忽视了指数运算的增长速度

题解参考自链接：[面试题14- II. 剪绳子 II（数学推导 / 贪心思想 + 快速幂求余，清晰图解） - 剪绳子 II - 力扣（LeetCode）](https://leetcode.cn/problems/jian-sheng-zi-ii-lcof/solution/mian-shi-ti-14-ii-jian-sheng-zi-iitan-xin-er-fen-f/)

原来在求 3 ^ a 时已经溢出，现在需要考虑的是取模的处理。本题需要取模的原因是 n 值的取值范围扩大了，导致计算过程中求 3 的次方时，数值指数扩大，需考虑溢出情况。我们的取余应该在做指数运算后进行

两种解决方案：循环求余 和 快速幂求余，两种方法均基于以下运算规则

(x∗y)%p=[(x%p)(y%p)]%p

一种比较通俗的解法是 循环取余：

x^a % p=[(x^(a−1)%p) \* (x%p)] % p=[(x^(a−1)%p) \* x]%p

本题中 a=3，p=10^9+7，利用此公式，可通过循环操作依次取x、x^2、x^3、...、x^a 对 p 的余数，保证每轮中间值在取值范围内。

```go
func cuttingRope(n int) int {
   x := int(math.Pow(10,9) + 7)
   if n<=3{
      return n-1
   }
   a,b := n/3,n%3
   if b == 0{
      return remainder(3,a,x)
   }
   if b == 1{
      return remainder(3,a-1,x) * 4 %x
   }
   return remainder(3,a,x)%x *2 %x
}
// 循环求余法
func remainder(a,b,p int) int {
   rem := 1
   for i:=0;i<b;i++{
      rem = (rem*a)%p
   }
   return rem
}
```

第二种解决方案是快速幂求余

我们在 day20 的 求数值的整数次方 中已经讲解过了快速幂，这里就不再赘述了，再将其每次操作进行取余即可。

```go
// 快速幂求余法
func fastmul_remainder(a,b,p int) int {
   rem := 1
   for b>0{
      if b&1==1{
         rem = (rem*a)%p
      }
      a = a*a % p
      b /= 2
   }
   return rem
}
```

对这道题的取余我想再多说几句，关于为什么可以这样做，让我困惑了挺长一段时间，下面假设很不合理，但这确实是我的困惑所在。
我做了一个这样的假设，a=3，b=6，p=90，范围最大值为100，3^6=486，远大于100，产生溢出情况，用循环求余解决溢出，3^4=81，对90取余还是81,3^5=243，会直接溢出，还是没有避免溢出的情况啊。。然后，我是这么想的，题目中 p=10^9+7，MaxInt32=2147483647，题目没有设定数据范围，那先假设整数为最大为 int64，p要小于 MaxInt32，那一个小于p的数乘以一个小于p的数，结果一定在 int64 的取值范围内！这下就反应过来，我之前的假设是非常不合理的。
**另外，10^9+7这个数好像经常用用于取模，为什么是这么数字，我查了一下大概是这么说的，一方面，该数字要足够大，另一方面，需要是一个大的质数来减少冲突。而10^9+7这个数字，相加不超过 int32，相乘不超过 int64. 一般来说x的选取只要10^x＋7保证比初始输入数据的范围大就可以。比如有些数据范围小的题为了避免用long long而把模数设定为10007。**

**这道题中主要用到 10^9+7 相乘不超过 int64 的性质。**

### 43.1~n整数中 1 出现的次数

> 输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
>
> 例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
>
> 示例 1：
>
> 输入：n = 12
> 输出：5
> 示例 2：
>
> 输入：n = 13
> 输出：6
>
>
> 限制：
>
> 1 <= n < 2^31

这道题看的官方题解，数学题么，主要还是找规律和做总结。

解题思路为统计各位上1出现的次数。设输入数字 x 为 n 位数，描述为 x=xn_xn−1_xn−2...x2_x1
若当前统计为 x3 位，记高位为 xn_xn−1_xn−2...x4 ，记为 high，低位为 x2_x1 ，记为 low
若数字 x 在 x3 位上的数字为 m，当前位上的因子 factor 为 10^(3-1) ,根据 m 的大小分情况讨论：

- m=0,则该位上1出现的次数为 high * factor
- m=1,则该位上1出现的次数为 high * factor + low + 1
- m>1,则该位上 1 出现的次数为 (high+1) * factor

问题来了，这种情况讨论是如何总结出来的呢，我觉得，就是单纯凭感觉找规律吧，为编程方便，从个位，也就是最低位开始统计1出现的次数。

### 44.数字序列中某一位的数字

> 数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
>
> 请写一个函数，求任意第n位对应的数字。
>
> 示例 1：
>
> 输入：n = 3
> 输出：3
> 示例 2：
>
> 输入：n = 11
> 输出：0
>
>
> 限制：
>
> 0 <= n < 2^31

这是非常考验边界问题处理的一道题目，解题思路不难想到，但要想ac，非常有难度。

题解参考自：[面试题44. 数字序列中某一位的数字（迭代 + 求整 / 求余，清晰图解） - 数字序列中某一位的数字 - 力扣（LeetCode）](https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/)

解题分为三步：

1. 确定 第n位 所在数字的数字长度digit，比如，个位数 3 的数字长度为1，百位数 202 的数字长度为3
2. 确定 第n位 所在的数字 num，比如三位数的第二个数字为 101
3. 确定 第n位 在数字 num 的第几位，比如三位数 101 的第二位数字为 0

一步一步来

对每个数字长度范围的数字规律总结

| 数字范围  | 单个数字长度 | 数字数量  | 该范围所有数字长度之和 |
| --------- | ------------ | --------- | ---------------------- |
| 1~9       | 1            | 9         | 9                      |
| 10~99     | 2            | 90        | 180                    |
| 100~999   | 3            | 900       | 2700                   |
|           |              |           |                        |
| start~end | digit        | 9 * start | 9 * start * digit      |

由上表，可得

数字长度递归公式：digit = digit + 1

起始数字递归公式：start = start * 10

该范围内所有数字长度之和：count = 9 * start * digit

一、确定 第n位 所在的数字长度

循环执行 n 减去一位数、二位数、...、的数字长度之和，直至 n<=count 跳出

二、确定 第n位 所在的数字

这里建议看一下题解开头附上的链接，有讲解的图片，我这里直接敷上公式

所求数位 在从数字 start 开始的第 [(n - 1) / digit] 个 数字 中（ start 为第 0 个数字）。

三、确定 第n位 在该数字的哪一位，很明显这里是用取余操作

第n位在 num 的第 (n-1)%2 位

```go
import "strconv"

func findNthDigit(n int) int{
   // digit为数字长度，start为该长度数字的第一个数字
   // count为该长度数字占据的位数
   digit,start,count := 1,1,9
   // 1.确定n所在数字的数字长度
   for n > count{
      n -= count
      start *= 10
      digit ++
      count = 9 * start * digit
   }
   // 确定n所在的数字num
   num := start + (n-1)/digit
   str := strconv.Itoa(num)
   // 确定n所在数字num的哪一位
   res,_ := strconv.Atoi(string(str[(n-1)%digit]))
   return res
}
```
