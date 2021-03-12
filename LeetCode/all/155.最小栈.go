/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (56.38%)
 * Likes:    834
 * Dislikes: 0
 * Total Accepted:    210.4K
 * Total Submissions: 373.1K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 * 
 * 
 * 
 * 
 * 示例:
 * 
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 * 
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 * 
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 * 
 * 
 */

// @lc code=start
type MinStack struct {
    // 一个存储push的元素的栈1
    value []int
    // 另一个存储最小值栈2
    minStack []int
    // 栈2有一个性质，每个栈2内元素位置为栈顶时，始终表示当前栈1最小的元素
    // 比如
    // 栈1 [1 2 3 4 0]
    // 栈2 [1 1 1 1 0]
    // 这样出栈同同步出栈，始终最小值就是栈2的栈顶，使用了一个当前最优解的思路
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        []int{},
        []int{},
    }
}

func (this *MinStack) Push(x int)  {
    this.value = append(this.value,x)
    var min int
    // 入栈时刷新最小栈，始终栈顶保存最小值
    if len(this.minStack)==0|| this.minStack[len(this.minStack)-1] > x{
        min = x
    }else{
        min = this.minStack[len(this.minStack)-1]
    }
    this.minStack = append(this.minStack,min)
}

func (this *MinStack) Pop()  {
    // 这里不判断也是可以的，go语言里不会报错
    if len(this.value) == 0{
        return
    }
    // 同步出栈即可
    this.value = this.value[:len(this.value)-1]
    this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
    if len(this.value) == 0{
        return -1
    }
    return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
    if len(this.value) == 0{
        return -1
    }
    return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

