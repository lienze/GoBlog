#!/bin/bash
cd ../src
find ./ -name "*.go" > ./cscope.files
cscope -Rbq
ctags -R

