package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

// rpc 远程调用 如何想本地调用一样

func Add(a, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get("http://127.0.0.1:8000/add?a=888&b=555")
	body, _ := res.Body()

	fmt.Println(string(body))
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	fmt.Println(rspData)
	return rspData.Data
}

type ResponseData struct {
	Data int `json:"data"`
}

func main() {
	res := Add(500, 600)
	fmt.Println(res)
}
