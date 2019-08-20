#!/bin/bash

# update the js and css resource
# rm ./public/js/*.min.js
java -jar ./the3party/yuicompressor-2.4.8.jar --type js --charset utf-8 -o '.js$:.min.js' ./html/js/*.js
mv ./html/js/*.min.js ./public/js/

java -jar ./the3party/yuicompressor-2.4.8.jar --type css --charset utf-8 -o '.css$:.min.css' ./html/css/*.css
mv ./html/css/*.min.css ./public/css/

# start making main binary
go build -gcflags "-N -l" -i -o ./main ../main.go

if [ -f "./main" ];then
	echo "build succeed!"
fi
