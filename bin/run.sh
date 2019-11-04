#!/bin/sh

if [ ! -d "./log/"  ];then
	mkdir -p "./log"
fi

#echo $1
if [ $1x = "--release"x ];then
	logname=$(date +%Y%m%d%k%M%S)
	echo "release mode"
	./GoBlog > ./log/$logname 2>&1 &
else
	echo "debug mode"
	go run ../main.go
fi

