package main

import (
	"fmt"
	"time"

	glog "github.com/cihub/seelog"
	"github.com/jsonkey/example/src/utils"
)

type PersonInfo struct {
	ID   string
	Name string
	Age  int16
}

func main() {

	defer glog.Flush()
	glog.Info("Prepare to repel boarders")

	glog.Debug("Test Debug")

	fmt.Println("file not exists:", utils.LocalPathExists("/anc"))

	var person1 map[string]PersonInfo

	person1 = make(map[string]PersonInfo, 2)

	person1["123"] = PersonInfo{"123", "tommy", 30}
	fmt.Println("你好", person1["123"])

	person, ok := person1["23"]

	// ok return bool
	if ok {
		fmt.Println("Found person", person)
	} else {
		fmt.Println("Not Found")
	}

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.23456
	MyPrintf(v1, v2, v3, v4)

	// chan
	ch := make(chan int, 1024)
	defer close(ch)

	timeout := make(chan bool, 1)
	defer close(timeout)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	ch <- 1

	select {
	case <-ch:
		fmt.Println("chan get:")
	case <-timeout:
		fmt.Println("no chan get:")
	}

}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int type")
		case string:
			fmt.Println(arg, "is string type")
		case int64:
			fmt.Println(arg, "is int64 type")
		default:
			fmt.Println(arg, "is default type")
		}
	}
}
