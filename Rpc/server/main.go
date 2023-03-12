package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// call id 的问题 我们可以使用 url的path
	// 数据传输协议 http 的协议 url 参数传递协议
	// 网络传输协议 http
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		fmt.Println("path:", r.URL.Path)
		// http://127.0.0.1:8000/add?a=1&b=2
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		data, _ := w.Write(res)

		fmt.Println(data)
	})

	_ = http.ListenAndServe(":8000", nil)
}
