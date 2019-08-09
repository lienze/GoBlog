#!/bin/bash
if [ -d ./public/result ];then
	rm -rf ./public/result
fi
mkdir ./public/result
gitstats ../../GoBlog ./public/result
