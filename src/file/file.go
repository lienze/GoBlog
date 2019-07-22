package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func InitFiles() error {
	//fmt.Println("InitFiles...")
	postPath := "../src/post/"
	files, _ := ioutil.ReadDir(postPath)
	for _, f := range files {
		ReadFile(postPath + f.Name())
	}
	var bInit bool = true
	if bInit == true {
		return nil
	} else {
		return errors.New("InitFiles error")
	}
}

func ReadFile(name string) {
	fmt.Println("Start ReadFile", name)
	if contents, err := ioutil.ReadFile(name); err == nil {
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("content:", string(result))
	} else {
		fmt.Println(err)
	}
}
