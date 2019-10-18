#!/bin/bash
# ----------------------------------------------vars define
this_dir=`pwd`
dist_dir="$this_dir/dist"
cd $this_dir/../
src_dir=`pwd`
cd $this_dir
# ------------------------------------------function define
function delete_dir()
{
	if [ -d $1 ];then
		rm -rf $1
		echo "delete $1 succeed"
	fi
}

function delete_file()
{
	if [ -f $1 ];then
		rm $1
		echo "delete $1 succeed"
	fi
}

# -----------------------------------------------main logic
delete_dir $dist_dir
delete_file $this_dir/GoBlog
delete_file $this_dir/dist.tar.gz
echo "Finished clean!"
