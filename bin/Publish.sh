#!/bin/bash
# ----------------------------------------------vars define
this_dir=`pwd`
dist_dir="$this_dir/dist"
cd $this_dir/../
src_dir=`pwd`
cd $this_dir
# ------------------------------------------function define
function build_err()
{
	rm -rf $dist_dir
	exit
}

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
delete_file $this_dir/main
delete_file $this_dir/dist.tar.gz
mkdir $dist_dir
cp -rf "$this_dir/html" $dist_dir
cp -rf "$this_dir/post" $dist_dir
cp -rf "$this_dir/config" $dist_dir
cp -rf "$this_dir/public" $dist_dir
cp -rf "$this_dir/cache" $dist_dir
source "$this_dir/build.sh"
cp -rf "$src_dir/version" $dist_dir
if [ -f "$this_dir/main" ];then
	cp -f "$this_dir/main" $dist_dir
else
	echo "ERROR:do not exist main"
	build_err
fi
echo "publish success!"
delete_dir $dist_dir/html/js
delete_dir $dist_dir/html/css
delete_file $dist_dir/cache/dump.rdb

# ---------------------------------------------third party
mkdir -p $dist_dir/the3party/
cp -rf "$this_dir/the3party/redis" "$dist_dir/the3party/"

# ----------------------------------------------copy script
cp "$this_dir/startcache.sh" $dist_dir

# -----------------------------------------------pack dist
tar -czf dist.tar.gz ./dist
if [ -f "$this_dir/dist.tar.gz" ];then
	echo "pack success!"
fi
