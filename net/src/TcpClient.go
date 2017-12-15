package main

import (
	"net"
	"fmt"
)

func main() {
	conn,e:=net.Dial("tcp","127.0.0.1:8888")

	if e!=nil {
		panic(e)
	}
	defer conn.Close()

	conn.Write([]byte("i am client"))

	b:=make([]byte,1024)
	conn.Read(b)
	fmt.Println(string(b))
}
