#!/bin/bash
this_dir=`pwd`
dist_dir="$this_dir/dist"
# echo $this_dir
if [ -d $dist_dir ];then
	rm -rf $dist_dir
fi
mkdir $dist_dir
cp -rf "$this_dir/html" $dist_dir

