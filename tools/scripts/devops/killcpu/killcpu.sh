#! /bin/sh 
# filename killcpu.sh
if [ $# != 1 ] ; then
  echo "USAGE: $0 <CPUs>"
  exit 1;
fi
for i in `seq $1`
do
  echo -ne " 
i=0; 
while true
do
i=i+1; 
done" | /bin/sh &
  pid_array[$i]=$! ;
done
 
for i in "${pid_array[@]}"; do
  echo 'kill ' $i ';';
done