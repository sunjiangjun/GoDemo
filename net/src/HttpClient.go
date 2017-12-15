package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func get() {
	resp, e := http.Get("https://www.baidu.com")
	if e != nil {
		panic(e)
	}
	defer resp.Body.Close()
	message, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(message))
}

func post()  {

	resp,e:=http.Post("http://www.baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))

	if e!=nil {
		panic(e)
	}

	defer resp.Body.Close()

	b,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func newRequest()  {

	client:=&http.Client{}
	request,_:=http.NewRequest("POST","http://www.baidu.com",strings.NewReader("name=cjb"))

	resp,e:=client.Do(request)

	if e!=nil {
		panic(e)
	}

	defer resp.Body.Close()

	b,_:=ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}


func main() {
	//get()

	//post()

	newRequest()
}
