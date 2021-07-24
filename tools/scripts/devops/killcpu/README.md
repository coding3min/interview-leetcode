使用方法很简单，参数3表示消耗3颗CPU的资源，运行后，会有一堆 kill 命令，方便 kill 进程：

``` BASH
[root@test02 ~]# ./killcpu.sh 3
kill  30104 ;
kill  30106 ;
kill  30108 ;
[root@test02 ~]# top 
top - 15:27:31 up 264 days, 23:39,  4 users,  load average: 0.86, 0.25, 0.19
Tasks: 185 total,   5 running, 180 sleeping,   0 stopped,   0 zombie
Cpu0  : 100.0% us,  0.0% sy,  0.0% ni,  0.0% id,  0.0% wa,  0.0% hi,  0.0% si
Cpu1  :  0.0% us,  0.0% sy,  0.0% ni, 100.0% id,  0.0% wa,  0.0% hi,  0.0% si
Cpu2  : 100.0% us,  0.0% sy,  0.0% ni,  0.0% id,  0.0% wa,  0.0% hi,  0.0% si
Cpu3  : 100.0% us,  0.0% sy,  0.0% ni,  0.0% id,  0.0% wa,  0.0% hi,  0.0% si
Mem:   8165004k total,  8095880k used,    69124k free,    53672k buffers
Swap:  2031608k total,   103548k used,  1928060k free,  6801364k cached
```
