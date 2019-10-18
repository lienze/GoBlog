#!/bin/sh
DetailInfo=`ps ux | grep "GoBlog" | grep -v "grep"`
IDList=`ps ux | grep "GoBlog" | grep -v "grep" | awk {'print $2'}`
echo "[Detail Process Info]"$DetailInfo
for ID in $IDList
do
	if [ ! -z $ID ];then
		echo "Kill process ID:"$ID" are you sure? (Y/n)\c"
		read user_answer
		# echo $user_answer
		if [ ${user_answer} = 'y' ] || [ ${user_answer} = 'Y' ]; then
			kill $ID
		fi
	else
		echo "no process found!"
	fi
done

