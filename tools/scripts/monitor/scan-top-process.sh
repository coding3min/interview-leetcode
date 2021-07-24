#/bin/bash


if [ "$1" ] 
then
    limitline=$1
else
    echo "usage: ./scan-top-process.sh {limit line}"
    echo "ep: ./scan-top-process.sh 80"
    exit -1
fi

top10res=`timeout 5 ps caux|sort -k3nr|head -10|awk '{print $2}'`
cnt=0

echo `date`
for i in $top10res;do
    let cnt=$cnt+1
    echo top$cnt
    echo cmdline: `cat /proc/$i/cmdline`
    echo "%CPU   RSS %MEM   PID"
    res=`timeout 5 ps -o %cpu,rss,%mem,pid -mp $i | grep -v "-" | grep -v CPU` 
    cpu_usage=`echo $res | awk '{print $1}'`
    if [[ `expr $cpu_usage \> $1` -eq 1 ]];then
        echo -e "\033[31m $res \033[0m"
    else
        echo -e "\033[32m $res \033[0m"
    fi
done