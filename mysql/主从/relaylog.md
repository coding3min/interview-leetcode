 前言：MySQL进行主主复制或主从复制的时候会在home目录下面产生相应的relay log，本文档总结这些相关参数的定义及解释.

### 1、什么是relay log

> The relay log, like the binary log, consists of a set of numbered files containing events that describe database changes, and an index file that contains the names of all used relay log files.
> The term "relay log file" generally denotes an individual numbered file containing database events. The term"relay log" collectively denotes the set of numbered relay log files plus the index file

来源： [官网文档](http://dev.mysql.com/doc/refman/5.5/en/slave-logs-relaylog.html)

理解：relay log很多方面都跟binary log差不多。

区别是：从服务器I/O线程将主服务器的二进制日志读取过来记录到从服务器本地文件，然后SQL线程会读取relay-log日志的内容并应用到从服务器，从而使从服务器和主服务器的数据保持一致

### 2、relay log的相关参数说明

通过语句：`show variables like '%relay%'``，查看查询到`relay`的所有相关参数
``` sql
mysql> show variables like '%relay%';

+-----------------------+----------------+

| Variable_name | Value |

+-----------------------+----------------+

| max_relay_log_size        

| relay_log                 

| relay_log_basename       

| relay_log_index          

| relay_log_info_file        

| relay_log_info_repository  

| relay_log_purge            

| relay_log_recovery         

| relay_log_space_limit     

| sync_relay_log               

| sync_relay_log_info  

      

+-----------------------+----------------+
```
参数详细解释：



 **max_relay_log_size**

标记relay log 允许的最大值，如果该值为0，则默认值为max_binlog_size(1G)；如果不为0，则max_relay_log_size则为最大的relay_log文件大小；

**relay_log**

定义relay_log的位置和名称，如果值为空，则默认位置在数据文件的目录（datadir），文件名为host_name-relay-bin.nnnnnn（By default, relay log file names have the form host_name-relay-bin.nnnnnn in the data directory）；

**relay_log_index**

同relay_log，定义relay_log的位置和名称；一般和relay-log在同一目录

**relay_log_info_file**

设置relay-log.info的位置和名称（relay-log.info记录MASTER的binary_log的恢复位置和relay_log的位置）

**relay_log_purge**

是否自动清空不再需要中继日志时。默认值为1(启用)。

**relay_log_recovery**

当slave从库宕机后，假如relay-log损坏了，导致一部分中继日志没有处理，则自动放弃所有未执行的relay-log，并且重新从master上获取日志，这样就保证了relay-log的完整性。默认情况下该功能是关闭的，将relay_log_recovery的值设置为 1时，可在slave从库上开启该功能，建议开启。

**relay_log_space_limit**

防止中继日志写满磁盘，这里设置中继日志最大限额。但此设置存在主库崩溃，从库中继日志不全的情况，不到万不得已，不推荐使用；

**sync_relay_log**

这个参数和sync_binlog是一样的，

当设置为1时，slave的I/O线程每次接收到master发送过来的binlog日志都要写入系统缓冲区，然后刷入relay log中继日志里，这样是最安全的，因为在崩溃的时候，你最多会丢失一个事务，但会造成磁盘的大量I/O。

当设置为0时，并不是马上就刷入中继日志里，而是由操作系统决定何时来写入，虽然安全性降低了，但减少了大量的磁盘I/O操作。这个值默认是0，可动态修改，建议采用默认值。

**sync_relay_log_info**

这个参数和sync_relay_log参数一样，当设置为1时，slave的I/O线程每次接收到master发送过来的binlog日志都要写入系统缓冲区，然后刷入relay-log.info里，这样是最安全的，因为在崩溃的时候，你最多会丢失一个事务，但会造成磁盘的大量I/O。当设置为0时，并不是马上就刷入relay-log.info里，而是由操作系统决定何时来写入，虽然安全性降低了，但减少了大量的磁盘I/O操作。这个值默认是0，可动态修改，建议采用默认值。



### 总结

以上只是简单的介绍了每个参数的作用，这些参数具体的设置还是需要根据每个用户的实际系统情况进行设置的；

推荐从库线上环境使用以下配置
``` sql
#slave
replicate-do-db=cloud
log-slave-updates
replicate-wild-do-table=cloud.%
binlog-ignore-db=mysql
slave-skip-errors=1032,1062,1053,1146,2003
#relay log
max_relay_log_size = 0；
relay_log=$datadir/relay-bin
relay_log_purge = 1;
relay_log_recovery = 1;
sync_relay_log =0；
sync_relay_log_info = 0；
```

- `cloud`是库名，用于跨库同步，如果有多个用逗号隔开
- 如果是mha环境，则| relay_log_purge 不要开启，设置为0，可以使用 purge_relay_logs 来定期清除


本站整理自：https://blog.51cto.com/douya/1788753