/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [284] 顶端迭代器
 *
 * https://leetcode-cn.com/problems/peeking-iterator/description/
 *
 * algorithms
 * Medium (72.26%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    7.4K
 * Total Submissions: 10.2K
 * Testcase Example:  '["PeekingIterator","next","peek","next","next","hasNext"]\n' +
  '[[[1,2,3]],[],[],[],[],[]]'
 *
 * 给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek() 操作的顶端迭代器 --
 * 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。
 * 
 * 示例:
 * 
 * 假设迭代器被初始化为列表 [1,2,3]。
 * 
 * 调用 next() 返回 1，得到列表中的第一个元素。
 * 现在调用 peek() 返回 2，下一个元素。在此之后调用 next() 仍然返回 2。
 * 最后一次调用 next() 返回 3，末尾元素。在此之后调用 hasNext() 应该返回 false。
 * 
 * 
 * 进阶：你将如何拓展你的设计？使之变得通用化，从而适应所有的类型，而不只是整数型？
 * 
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
    nextCache []int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iter: iter,
    }
}

func (this *PeekingIterator) hasNext() bool {
    return len(this.nextCache)!=0 || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
    if len(this.nextCache)>0{
        num := this.nextCache[0]
        this.nextCache = this.nextCache[1:]
        return num
    }
    return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if len(this.nextCache)!=0{
        return this.nextCache[0]
    }
    if this.iter.hasNext(){
        num := this.iter.next()
        this.nextCache = append(this.nextCache,num)
        return num
    }
    return 0
}
// @lc code=end

