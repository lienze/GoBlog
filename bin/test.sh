#!/bin/bash
# go test ../... -cover -v 

# go test ../src/cache -cover -v
# go test ../src/zversion -cover -v
# go test ../src/ztime -cover -v

package_list=(
	../src/cache
	../src/zversion
	../src/ztime
)
for i in ${package_list[@]};
do
	echo $i
	go test $i -cover -v
done

