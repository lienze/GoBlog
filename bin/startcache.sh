#!/bin/sh
logname=$(date +%Y%m%d%k%M%S)
if [ ! -d "./cache/log/" ];then
	mkdir -p "./cache/log"
fi
echo $logname
if [ `uname` = "Darwin" ];then
	./the3party/redis/mac/redis-server ./config/redis_master.conf > ./cache/log/$logname 2>&1 &
else
	./the3party/redis/linux/redis-server ./config/redis_master.conf > ./cache/log/$logname 2>&1 &
fi

