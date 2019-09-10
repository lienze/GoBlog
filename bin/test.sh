#!/bin/bash
# go test ../... -cover -v 
go test ../src/cache -cover -v
go test ../src/zversion -cover -v
go test ../src/ztime -cover -v

