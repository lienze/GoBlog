package main

import (
	"GoBlog/src/server"
	"fmt"
)

func main() {
	err := server.NewServer()
	if err != nil {
		fmt.Println(err)
	}
}
