# 操作系统

## 并发和并行的理解？

* 并发：在一个时间段中多个程序都启动运行在同一个处理机中，比如线程
* 并行：假设目前A，B两个进程，两个进程分别由不同的 CPU  管理执行，两个进程不抢占 CPU  资源且可以同时运行，这叫做并行。单核CPU是伪并行

## 线程与进程的优缺点

多线程的优点：

* 更加高效的内存共享。多进程下内存共享不便；
* 较轻的上下文切换。因为不用切换地址空间，CR3寄存器和清空TLB。

多进程的优点：

* 各个进程有自己内存空间，所以具有更强的容错性，不至于一个集成crash导致系统崩溃
* 具有更好的伸缩性，因为进程将地址空间，页表等进行了隔离，可以在不同的服务器上进行水平伸缩

## 如何提升多线程的效率？

答者：彬

1. 使用线程池，减少线程的创建销毁带来的开销
2. 根据不同的线程类型确定线程数量，例如cpu繁琐的，核线比1:2，I/O繁琐的1:1
3. 压测，看具体项目能吃多少线程
4. 在线程数无法减少的情况下，根据物理内存调整jvm堆大小，为线程提供足够内存空间
5. 升级物理机的cpu和内存
6. 单台机极限的情况下，使用集群
7. 最基本的，优化代码，减低复杂度

小熊补充

* 尽量使用线程池，从而不用频繁的创建，销毁线程；
* 减少线程之间的同步和通信；
* 通过Huge Page的方式避免产生大量的缺页异常；
* 避免需要频繁共享写的数据。

## 进程的三个状态

![](.gitbook/assets/2021-03-17-19-58-57.png)

* 就绪 —&gt; 执行：准备就绪，调度器满足了的需求
* 执行 —&gt; 阻塞：正在执行的进程由于发生某事件（例如请求I/O而等待I/O完成等）而暂时无法继续执行时，便放弃处理机而处于暂停状态，这时即使把处理机分配给进程也无法运行，故称该进程处于阻塞状态。
* 阻塞 —&gt; 就绪：处于阻塞状态的进程，在其等待的事件已经发生，如输入/输出完成，资源得到满足或错误处理完毕时，处于等待状态的进程并不马上转入执行状态，而是先转入就绪状态，然后再由系统进程调度程序在适当的时候将该进程转为执行状态；
* 执行 —&gt; 就绪：正在执行的进程，因时间片用完而被暂停执行，或在采用抢先式优先级调度算法的系统中, 当有更高优先级的进程要运行而被迫让出处理机时，该进程便由执行状态转变为就绪状态。

## 进程与线程有什么区别？

* 进程是资源分配的最小单位，线程是进程的一个实体，也是 CPU 调度和分派的基本单位，它是比进程更小的能独立运行的基本单位，有时又被称为轻量级进程。
* 创建进程或撤销进程，系统都要为之分配或回收资源，操作系统开销远大于创建或撤销线程时的开销；
* 不同进程地址空间相互独立，同一进程内的线程共享同一地址空间。一个进程的线程在另一个进程内是不可见的；
* 进程间不会相互影响，而一个线程挂掉将可能导致整个进程挂掉；

所以

* 如果以多进程的方式运行，那么允许多个任务同时运行；
* 如果以多线程的方式运行，那么允许将单个任务分成不同的部分运行；
* 为了防止进程/线程之间产生冲突和允许进程之间的共享资源，需要提供一种协调机制。

## 多线程一定就比单线程效率高吗？

* 多线程对 CPU，内存等都要求比较高，如果存在的上下文切换过于耗时，互斥时间太久，效率反而会低。
* 也就是说线程大于资源数容易出现抢占，少了又会出现浪费

## 进程间怎么通信？

* linux中的管道，优点是简单，缺点是不适合频繁的通信。大小受限且只能承载无格式字节流的方式
* 消息队列，按照独立单元（消息体）进行发送，双方约定好约定好消息类型或者正文的格式，发送到消息队列的数据太大，需要拷贝的时间也就越多
* 共享内存，使用ipcs创建共享内存，不需要拷贝，可立即感知，但容易出现冲突和覆盖

## 怎么保证多写入不出现冲突？

**信号量（计数器）**

为了防止冲突，我们得有个约束或者说一种保护机制。使得同一份共享的资源只能一个进程使用

p操作为申请资源，会将数值减去M，表示这部分被他使用了，其它进程暂时不能用。v操作是归还资源操作，告知归还了资源可以用这部分。

**信号**

一旦进程发送某一个信号给另一个进程，另一进程将执行相应的函数进行处理。也就是说把可能出现的异常等问题准备好，一旦信号产生就执行相应的逻辑即可。

**套接字（远程通信）**

可以跨主机通信 acept socket

## 同步、异步、阻塞、非阻塞的概念

* 同步：当一个同步调用发出后，调用者要一直等待返回结果。通知后，才能进行后续的执行。
* 异步：当一个异步过程调用发出后，调用者不能立刻得到返回结果。实际处理这个调用的部件在完成后，通过状态、通知和回调来通知调用者。
* 阻塞：是指调用结果返回前，当前线程会被挂起，即阻塞。
* 非阻塞：是指即使调用结果没返回，也不会阻塞当前线程。

## 进程调度算法有哪些？

* 先来先服务调度算法

每一次的调度都从队列中选择最先进入队列的投入运行。

* 时间片轮转调度算法

  按照进程到达的时间排序, 每个进程时间片大小相等，时间片不够用时等待。

* 短作业优先调度算法

缺点是长作业等待时间进程饥饿

* 最短剩余时间优先调度算法

同样的问题

* 高响应比优先调度算法

短作业优先+等待时间越长优先级越高

* 优先级调度算法

每次从后备作业队列中选择优先级最髙的一个或几个作业，将它们调入内存，分配必要的资源，创建进程并放入就绪队列。

## 死锁是什么？

彼此拥有对方正在申请的互斥量，解决方法：

* 按照**一定的先后顺序**申请这些互斥量。

## 产生死锁的必要条件

* 资源互斥；如果是共享不存在死锁等待
* 非抢占：如果可以抢占直接抢占
* 循环等待：必要不充分条件，同时等待时可以顺序执行，除非出现环路等待

## 解决死锁的基本方法

* 破坏互斥：使用类似CAS无锁结构
* 破坏不抢占：锁获取失败时自动放弃持有的锁，并重试
* 破坏环路等待：按顺序请求锁，给锁排序、给线程（安全序列）

## 怎么预防死锁?

* 破坏请求条件：一次性分配所有资源，这样就不会再有请求了；
* 破坏请保持条件：只要有一个资源得不到分配，也不给这个进程分配其他的资源；
* 破坏不可剥夺条件：当某进程获得了部分资源，但得不到其它资源，则释放已占有的资源；
* 破坏环路等待条件：系统给每类资源赋予一个编号，每一个进程按编号递增的顺序请求资源，释放则相反。

## 避免死锁的算法

**银行家算法**

* 当进程首次申请资源时，要测试该进程对资源的最大需求量，如果系统现存的资源可以满足它的最大需求量则按当前的申请量分配资源，否则就推迟分配。
* 当进程在执行中继续申请资源时，先测试该进程**已占用的资源数与本次申请资源数之和**是否超过了该进程对资源的最大需求量。
* 若超过则拒绝分配资源。若没超过则再测试系统**现存的资源能否满足**该进程尚需的最大资源量，若满足则按当前的申请量分配资源，否则也要推迟分配。

## 解决死锁

* 资源剥夺：挂起某些死锁进程，并抢占它的资源，将这些资源分配给其它死锁进程（但应该防止被挂起的进程长时间得不到资源）；
* 撤销进程：强制撤销部分、甚至全部死锁进程并剥夺这些进程的资源（撤销的原则可以按进程优先级和撤销进程代价的高低进行）；
* 进程回退：让一个或多个进程回退到足以避免死锁的地步。进程回退时自愿释放资源而不是被剥夺。要求系统保持进程的历史信息，设置还原点。

## 什么是缓冲区溢出？有什么危害？

缓冲区溢出是指当计算机向缓冲区内填充数据时超过了缓冲区本身的容量，溢出的数据覆盖在合法数据上。

## 你怎么理解操作系统里的内存碎片，有什么解决办法？

内存碎片通常分为内部碎片和外部碎片：

1. 内部碎片是由于采用固定大小的内存分区，当一个进程不能完全使用分给它的固定内存区域时就会产生内部碎片，通常内部碎片难以完全避免；
2. 外部碎片是由于某些未分配的连续内存区域太小，以至于不能满足任意进程的内存分配请求，从而不能被进程利用的内存区域。

现在普遍采取的内存分配方式是段页式内存分配。将内存分为不同的段，再将每一段分成固定大小的页。通过页表机制，使段内的页可以不必连续处于同一内存区域。

## 最后

如果文中有误，欢迎提pr或者issue，**一旦合并或采纳作为贡献奖励可以联系我直接无门槛**加入[技术交流群](https://mp.weixin.qq.com/s/ErQFjJbIsMVGjIRWbQCD1Q)

我是小熊，关注我，知道更多不知道的技术

![](.gitbook/assets/2021-03-17-19-57-33.png)
