#!/bin/sh
logname=$(date +%Y%m%d%k%M%S)
if [ ! -d "./cache/log/" ];then
	mkdir -p "./cache/log"
fi
echo $logname
./the3party/redis-server ./config/redis.conf > ./cache/log/$logname 2>&1 &

