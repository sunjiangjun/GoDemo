package main

import (
	"net/http"
	"fmt"
)

type MyHandler struct {
}

//请求处理器 必须实现的方法
func (handler *MyHandler) ServeHTTP(writer http.ResponseWriter,req *http.Request)  {
	writer.Write([]byte("hello,world"))
}


func handle()  {
	mux:=http.NewServeMux()

	//重定向
	rh:=http.RedirectHandler("http://www.baidu.com", 307)

	//注册响应处理器， 根据指定的路径
	mux.Handle("/foo",rh)

	//根据地址，建立连接
	http.ListenAndServe(":3000",mux)

}

func handle2()  {

	//注册匿名的处理器
	http.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello,world")
		writer.Write([]byte("hello,world"))

	})

	//注册自定义的处理器
	http.Handle("/boo",&MyHandler{})

	//建立连接 、开启监听
	http.ListenAndServe(":3000",nil)
}

func main() {
	//handle()
	handle2()
}
