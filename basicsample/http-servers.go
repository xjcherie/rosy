package basicsample

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) { //todo: req参数没有使用，为什么不可用删除？
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello) //todo: http有哪些方法
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
}
