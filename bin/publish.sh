#!/bin/bash
this_dir=`pwd`
echo $this_dir
if [ ! -d "$this_dir/dist" ];then
	mkdir "$this_dir/dist"
else
	rm -rf "$this_dir/dist"
	mkdir "$this_dir/dist"
fi
cp -rf "$this_dir/html" "$this_dir/dist"
