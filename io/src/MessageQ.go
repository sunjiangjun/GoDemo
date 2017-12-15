package main

import (
	"fmt"
	"runtime"
)

var change = make(chan int)

type Message struct {
	id   int
	data string
	do   func(pre string, id int)
}

func addMessage(req string, m Message) {
	m.do(req+m.data, m.id)
}

func main() {

	//查询当前运行使用CPU数量
	fmt.Println("设定本机可用CPU的数量：",runtime.NumCPU())

	//设定当前CPU 可使用的数量
	runtime.GOMAXPROCS(4)

	var n = 1

	task := Message{id: n, data: "i am message"}
	task.do = func(pre string, id int) {
		fmt.Println(id, pre)
		change <- 1
	}

	go addMessage("welcome", task)

	<-change
}
