# 消息队列

## 如何保证队列数据一致性，防止重复消费，幂等性

所有的消息中间件在实际的业务场景中都逃脱不了保证消息的一致性的问题

`kafka`实际上有个`offset`的概念，就是每个消息写进去，都有一个`offset`，可以理解成一个自增持的序号，一个个排下来

然后消费者`consumer`，每隔一段时间，就把自己消费过的消息提交一下，如果说出现宕机或者重启，则会继续从上次消费的序号接着往下排，继续消费

但是有的时候，消费者`consumer`消费的消息，由于各种原因，比如网络、宕机、停电。。。都没来得及写`offset`,这个时候少数消息会再次消费一次

这个时候，我们可以**用一个唯一的id标识来区分**，这不是消息中间件做的事，而是开发者要做的，比如你消费一个就往数据库插入一条记录，然后下次再去消费的时候，你去查一下，看看这个消息是否被消费了，消费了那就不要重复消费了。

（补充一下：确认一条数据在百万级别海量数据里是否存在？--可以用**布隆过滤器**）

* 根据主键查一下，如果这数据都有了，就别插入了，update一下（虽然重复插入会因为唯一键约束而报错，我觉得我们还是应该避免报错）
* 如果是写redis，反正每次都是set，天然幂等性
* 如果不是上面两个场景，那做的稍微复杂一点，你需要让生产者发送每条数据的时候，里面加一个全局唯一的id，类似订单id之类的东西，然后你这里消费到了之后，先根据这个id去比如redis里查一下，之前消费过吗？如果没有消费过，你就处理，然后这个id写redis。如果消费过了，就别处理了，保证别重复处理相同的消息即可。

引用:

* [rabbitMq-kafka消息高可用，一致性](https://blog.csdn.net/lanshen110119/article/details/89399084)
* [如何保证消息队列的高可用和幂等性以及数据丢失，顺序一致性](https://www.bilibili.com/read/cv1923046/)

## 如何预防rabbitmq消息的丢失

**情况1、生产者自己丢失了消息\(网络故障/发送失败\)**

解决方案：

rabbitmq：一般这种都是采用回调接口的方案（confirm模式），就是说你扔一个消息过去了，对方给你一个回调接口，告诉你成功了或者失败了，失败了你可以选择继续扔消息， \(重试机制等\)，来保证消息一定送达

开启confirm模式（异步的）之后，你每次写的消息都会分配一个唯一的id，然后如果写入了rabbitmq中，rabbitmq会给你回传一个ack消息，告诉你说这个消息ok了。如果rabbitmq没能处理这个消息，会回调你一个nack接口，告诉你这个消息接收失败，你可以重试。而且你可以结合这个机制自己在内存里维护每个消息id的状态，如果超过一定时间还没接收到这个消息的回调，那么你可以重发。

引用： [如何保证消息队列的高可用和幂等性以及数据丢失，顺序一致性](https://www.bilibili.com/read/cv1923046/)

**情况2、消息中间件弄丢了消息**

解决方案：以rabbitmq来说，开启**持久化**就好了，当发生宕机的时候，queue会自动从磁盘恢复数据，除非极其罕见的是，rabbitmq还没持久化，自己就挂了，可能导致少量数据会丢失的，但是这个概率较小

**情况3、消费者弄丢了消息**

rabbitmq如果丢失了数据，主要是因为你消费的时候，刚消费到，还没处理，结果进程挂了，比如重启了，rabbitmq认为你都消费了，这数据就丢了。

关闭rabbitmq自动ack机制，可以通过一个api来调用，每次代码里确保处理完的时候，再程序里ack。这样的话，如果还没处理完，就没有ack，那rabbitmq就认为你还没处理完，这个时候rabbitmq会把这个消费分配给别的consumer去处理，消息是不会丢的

## 如何预防kafak消息丢失

[TODO](https://github.com/minibear2333/interview-leetcode/issues/35)

## 如何处理消息积压

出现场景：消费能力被阻塞（消费者挂掉或者处理速度慢），生产者还不停的往队列里扔消息

解决方法：

* 快速恢复`consumer`服务，慢慢消费，如果积压的数据量太大的话恢复较慢
* 临时写脚本快速的把这批消息给消费掉，或者增加消费者数量/消费速度，要避免负载把其他服务打挂
* 扩容：提高相同消息的队列数量，出现问题时写脚本分发到不同队列里，再给每个队列指定消费者，消费结束后再恢复

预防：

* 提前准备多个队列在投递时随机投递，存储同类型无顺序要求的消息
* 使用多个消费者

## 消息过期或者队列满了怎么办

消息队列TTL超时或者队列满了数据会丢失，这个时候可以自己再去找消息，然后临时写个代码，自己再手动的去把这些消息重新推送到队列里去。

另一种解决方案

* 可以在 rabbitmq 中声明死信队列，死信队列为处理过期或不能正确路由的消息提供了驻留场所，可以防止消息丢失，便于分析无法消费的原因
* 写程序处理死信队列里的数据，并接入告警分析

如果投递不成功，需要把数据暂存内存或者暂存redis之类的数据库中，等待恢复时重试

## 怎么实现一个延时队列？

rabbitmq 可以配置[延时队列插件](https://github.com/rabbitmq/rabbitmq-delayed-message-exchange)

消费者是无感知的，可以正常消费，生产者设置延迟时间，到了时间后队列里才会出现消息

如果一定要手动实现，可以维护不同的队列指代不同的延迟时间，程序根据相应队列来保证最新消息的时间戳与当前时间延迟转发到实际的消费队列

引用：[过期时间，死信队列，延迟队列，优先队列，持久化](http://www.gxitsky.com/article/1604455229805099)

## 什么是优先队列，怎么实现

优先级高的消息先消费

在rabbitmq中，手动开启设置队列支持的最大优先级（建议不要设置太大，会消耗资源），在投递消息的时间设置优先级数值，队列会自动排序，但是如果消费速度太快，没有任何数据积压的时候，不存在排序也不存在优先级的问题

## 如何保证消息队列的顺序性

同一订单不同消息投递到不同位置，不同消费者消费了同一订单的不同消息，一般出现在

* 程序设计问题投递到不同队列
* 同一队列多个消费者并发消费，无法保证消费顺序

解决方法：

* 队列都是有顺序性保证的，在投递时，创建多个队列，hash投递，hash相同订单号投递同一个队列
* 消费时，保证同一个队列只允许一个消费者消费
* 多线程并发处理时，避免多进程，而是增加线程数，维护多个内存队列把消息归类

引用: [如何保证消息的顺序性](https://xie.infoq.cn/article/c84491a814f99c7b9965732b1)

