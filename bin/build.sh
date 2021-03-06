#!/bin/bash

# copy the version file from the root path
#cp ../version ./
#sed -i "1i# this file was generated by build.sh, don't modify." ./version

# update the js and css resource
# rm ./public/js/*.min.js
java -jar ./the3party/yuicompressor-2.4.8.jar --type js --charset utf-8 -o '.js$:.min.js' ./html/js/*.js
mv ./html/js/*.min.js ./public/js/

java -jar ./the3party/yuicompressor-2.4.8.jar --type css --charset utf-8 -o '.css$:.min.css' ./html/css/*.css
mv ./html/css/*.min.css ./public/css/

if [ $1x = "--release"x ];then
	# start making main binary
	go build -i -o ./GoBlog ../main.go
else
	go build -gcflags "-N -l" -i -o ./GoBlog ../main.go
fi

if [ -f "./GoBlog" ];then
	echo "build succeed!"
fi
