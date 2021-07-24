 我在测试主从方案的时候发现状态丢失了，同步用的`binlog`也不见了(binlog doesn't exist)，非常奇怪，回顾解决以后写在这里供大家参考。

### 报错与原因

发现错误信息类似于

``` error
Slave: Table 'XXX' doesn't exist
Error running query, ......，We stopped at log 'mysql-bin.000036' position 154.
```

赶快去找，本想看看主数据的这个文件的`position 154`是什么语句，使用语句

```sql
show binlog events in 'mysql-bin.000036'
```
居然返回了502。

又到主库的服务器查看了下`binlog`的存储情况，发现`binlog`的编号是从`36`开始的，前面的不见了！难道设置了参数定期删除`binlog`？

于是又来到了`my.cnf`文件，查看文件之后找到了一个`expire_logs_days`。经搜索，确定了这个参数就是删除以前`binlog`文件的“罪魁祸首”。

到这来，大概明白了为啥主从同步没有成功，因为这是基于`binlog`的（逐行扫描`sql`语句进行同步写入），如果`binlog`文件不全，就无法正确的进行主从同步。

### 解决办法

这种情况从删除那一天起，至今所有的同步语句全部都丢失了，所以除非可以精确的知道执行了哪些语句，或者那些语句都不重要可以忽略，不然都必须要清理数据库，备份主库，重新手动更新从库来解决。可以参考我的[备份数据库](https://coding3min.com/845.html)这篇文章的。

如果你精确的知识执行了哪些语句，需要先停止从库，执行丢失的语句，再进行从库同步设置。

``` sql
mysql> stop slave;
Query OK, 0 rows affected, 1 warning (0.00 sec)
mysql> reset slave;
Query OK, 0 rows affected (0.00 sec)
.....执行你的语句
mysql> change master to master_host='192.168.1.51', master_user='replslave', master_password='replslave', master_log_file='mysql-bin-000002',master_log_pos=168;
Query OK, 0 rows affected (0.11 sec)
```

如果不知道怎进行从库同步设置，请参考[mariadb/mysql建立主从](https://coding3min.com/844.html)

### 小结
1.根据情况设置`expire_logs_days`，位于`mariadb`的配置文件中，意思是超时天数，超过这个数值就清理掉过期的`binlog`
2.还有一个参数叫`max_binlog_size`，默认是`1G`，如果设置的太小可能导致大事物被截断，保持默认就好。


参考：[mysql的expire_logs_days参数引发的问题](https://blog.csdn.net/a7442358/article/details/102667545)
