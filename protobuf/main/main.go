package main

import (
	"encoding/json"
	"extreme-go/protobuf/helloworld"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	req := helloworld.HelloRequest{
		Name:    "hello",
		Age:     99,
		Courses: []string{"math", "english", "go", "gin"},
	}
	jsonStruct := Hello{Name: "erkelost", Age: 99, Courses: []string{"math", "english", "go", "gin"}}
	res, _ := json.Marshal(jsonStruct)
	rsp, _ := proto.Marshal(&req)
	fmt.Println(rsp)
	fmt.Println(res)
	fmt.Println(string(res))
	fmt.Println(string(rsp))
}
