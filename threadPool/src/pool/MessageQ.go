package pool


import (
	"fmt"
	"time"
)


var (
	Queue =make( chan  Task,10)
)


type Task struct {
	uuid int
	data interface{}
}

func NewTask(uuid int,data string) Task {
	return Task{uuid:uuid,data:data}
}


func AddTask(task Task)  {
	fmt.Println("add task:",task.uuid)
	Queue<-task
}


func printer(t Task)  {
	fmt.Println("printer...pre :",t.uuid)
}

func Loop()  {

	for {
		t:=<-Queue
		go printer(t)
		time.Sleep(100)
	}
}
