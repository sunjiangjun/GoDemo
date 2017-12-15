package main

import (
	_ "github.com/garyburd/redigo/redis"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var cl redis.Conn
var e error


/**
 建立连接
 */
func conn()  {
	cl,e=redis.Dial("tcp","127.0.0.1:8889")
	if e!=nil {
		panic(e)
	}
}


/**
  执行指令
 */

func send()  {
	//案例一
	//rec,err:=cl.Do("HGET","student","name")
	//if err !=nil {
	//	panic(err)
	//}
	//fmt.Println(redis.String(rec,err))


	//案例二

	//rec,err:=cl.Do("HKEYS","student")
	//if err !=nil {
	//	panic(err)
	//}
	//ss,_:=redis.Strings(rec,err)
	//for _,body:= range ss {
	//	fmt.Println(body)
	//}
	//fmt.Println(redis.Strings(rec,err))


	//案例三
	rec,err:=cl.Do("HKEYS","student")
	if err !=nil {
		panic(err)
	}
	ss,_:=redis.Values(rec,err)

	for _,body:= range ss {
		fmt.Println(string(body.([]byte)))
	}

	var v1  string
	redis.Scan(ss,&v1)
	fmt.Println(v1)
}


/**
  批量处理{多个指令一并处理，并FIFO的顺序接受和处理相应}
 */
func sendBuffer()  {

	err:=cl.Send("HKEYS","student")
	if err !=nil {
		panic(err)
	}

	cl.Send("HGET","student","name")
	cl.Flush()

	rec,e:=cl.Receive()
	fmt.Println(redis.Strings(rec,e))

	rec1,e:=cl.Receive()
	fmt.Println(redis.String(rec1,e))

}

/**
  订阅与发布
 */
func sub()  {

	rec,err:=cl.Do("PUBLISH","qycloud","hello,world")
	if err !=nil {
		panic(err)
	}
	ss,_:=redis.Values(rec,err)

	for _,body:= range ss {
		fmt.Println(string(body.([]byte)))
	}

	var v1  string
	redis.Scan(ss,&v1)
	fmt.Println(v1)
}



func main() {

	conn()
	defer  cl.Close()

	//send()

	//sendBuffer()

	sub()

}