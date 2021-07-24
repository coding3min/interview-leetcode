本操作是高危操作，会导致所有数据丢掉，用来恢复无法恢复的mysql，重建以后再导入备份好的数据。

1、停止MySql数据库的运行
2、删除`mysql/var`里面的所有数据库，这里是数据目录，具体请查看`/etc/my.cnf`的配置
3、进入mysql/bin目录下，执行`./mysql_install_db`命令
此时会在mysql/var目录下创建两个目录文件mysql、test
4、修改mysql、test两个目录及目录下所有文件的权限：
``` BASH
chown mysql:mysql -R mysql test
```
注意这里一定要加上-R参数，否则启动会报错
5、启动数据库
``` BASH
./mysqld_safe --user=mysql &
```
6、修改root密码
``` BASH
mysql/bin/mysqladmin -u root password "yourpasswd"
```
这样，就完成了MySql数据库的重建了。关于停止MySql的运行，直接用启动MYSQL服务命令也行，也可以用停止进程的方法。启动数据库也可以直接用启动MYSQL服务的命令来启动。