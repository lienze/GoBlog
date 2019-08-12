#!/bin/bash
if [ -d ./public/stats ];then
	rm -rf ./public/stats
fi
mkdir ./public/stats
git_stats generate -l zh_tw -p ../../GoBlog -o ./public/stats
