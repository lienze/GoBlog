#!/bin/bash
# ----------------------------------------------vars define
this_dir=`pwd`
dist_dir="$this_dir/dist"

# ------------------------------------------function define
function build_err()
{
	rm -rf $dist_dir
	exit
}

# -----------------------------------------------main logic
if [ -d $dist_dir ];then
	rm -rf $dist_dir
fi
if [ -f "$this_dir/main" ];then
	rm "$this_dir/main"
	echo "delete main success"
fi
mkdir $dist_dir
cp -rf "$this_dir/html" $dist_dir
source "$this_dir/build.sh"
if [ -f "$this_dir/main" ];then
	cp -f "$this_dir/main" $dist_dir
else
	echo "ERROR:do not exist main"
	build_err
fi
echo "publish success!"

