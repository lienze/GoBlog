#!/bin/bash

# update the js and css resource
# rm ./public/js/*.min.js
java -jar ./the3party/yuicompressor-2.4.8.jar --type js --charset utf-8 -o '.js$:.min.js' ./html/js/*.js
mv ./html/js/*.min.js ./public/js/

go build -gcflags "-N -l" -i -o ./main ../main.go

if [ -f "./main" ];then
	echo "build succeed!"
fi
