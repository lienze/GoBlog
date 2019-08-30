package zversion

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// get the version info from the version file
	var filever string
	if contents, err := ioutil.ReadFile("../../version"); err == nil {
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("content:", string(result))
		filever = result
	} else {
		fmt.Println(err)
		t.Log(err)
		t.FailNow()
	}

	if ver != filever {
		t.Log("ver:", ver, ",filever:", filever)
		t.FailNow()
	}
}
