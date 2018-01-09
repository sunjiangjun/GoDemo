package main

import (
	."./pool"
	"runtime"
	"time"
)

func addTask()  {
	for i:=0;i<100 ;i++ {
		t:=NewTask(i,"hello,world")
		AddTask(t)
	}
}

func main() {
	runtime.GOMAXPROCS(1)

	go Loop()
	 addTask()

	time.Sleep(1000000000)

}
