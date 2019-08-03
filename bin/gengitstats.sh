#!/bin/bash
if [ -d ./html/result ];then
	rm -rf ./html/result
fi
mkdir ./html/result
gitstats ../../GoBlog ./html/result
