#### 前言

MySQL 中的 binlog 日志记录了数据库中数据的变动，便于对数据的基于时间点和基于位置的恢复，但是 binlog 也会日渐增大，占用很大的磁盘空间，因此，要对 binlog 使用正确安全的方法清理掉一部分没用的日志

#### 方法一、手动清理 binlog

清理前的准备：

查看主库和从库正在使用的 binlog 是哪个文件

```bash
show master status \G
show slave status \G
```

在删除 binlog 日志之前，首先对 binlog 日志备份，以防万一

开始动手删除 binlog：
删除指定日期以前的日志索引中 binlog 日志文件

```bash
purge master logs before'2016-09-01 17:20:00';
```

或 删除指定日志文件的日志索引中 binlog 日志文件

```bash
purge master logs to'mysql-bin.000022';
```

> 注意：<font color="red" >**时间和文件名一定不可以写错，尤其是时间中的年和文件名中的序号，以防不小心将正在使用的 binlog 删除！！！**</font> > <font color="red" >切勿删除正在使用的 binlog！！！</font>

使用该语法，会将对应的文件和 mysql-bin.index 中的对应路径删除。

#### 方法二、通过设置 binlog 过期的时间，使系统自动删除 binlog 文件

临时生效

```sql
mysql> show variables like 'expire_logs_days';
+------------------+-------+
| Variable_name  | Value |
+------------------+-------+
| expire_logs_days |   0  |
+------------------+-------+
mysql> set global expire_logs_days = 30;    #设置binlog多少天过期
```

长期生效需要修改配置文件

```bash
vim /etc/my.cnf
内容
expire_logs_days=3 #代表3天
max_binlog_size = 1073741824 #默认是1G
```

注意：

过期时间设置的要适当，对于主从复制，要看从库的延迟决定过期时间，避免主库 binlog 还未传到从库便因过期而删除，导致主从不一致！！！

本站整理自：[mysql 正确清理 binlog 日志的两种方法](https://blog.csdn.net/tengxing007/article/details/79881519)
