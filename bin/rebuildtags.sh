#!/bin/bash
cd ../
if [ -f cscope* ];then
	rm cscope*
fi
if [ -f tags ];then
	rm ctags
fi
find ./ -name "*.go" > ./cscope.files
cscope -Rbq
ctags -R

