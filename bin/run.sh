#!/bin/sh

logname=$(date +%Y%m%d%k%M%S)
if [ ! -d "./log/"  ];then
	mkdir -p "./log"
fi

#echo $1
if [ $1x = "--debug"x ];then
	echo "debug mode"
	go run ../main.go
else
	echo "release mode"
	./main > ./log/$logname 2>&1 &
fi

