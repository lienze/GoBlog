#!/bin/sh
DetailInfo=`ps ux | grep "redis-server" | grep -v "grep"`
ID=`ps ux | grep "redis-server" | grep -v "grep" | awk {'print $2'}`
if [ ! -z $ID ];then
	echo "[Detail Process Info]"$DetailInfo
	echo "Kill process ID:"$ID" are you sure? (Y/n)\c"
	read user_answer
	# echo $user_answer
	if [ ${user_answer} = 'y' ] || [ ${user_answer} = 'Y' ]; then
		kill $ID
	fi
else
	echo "no process found!"
fi

