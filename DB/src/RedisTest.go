package main

import (
	"github.com/go-redis/redis"
	"fmt"

)


var (
	client *redis.Client
)

func testString()  {

	cmd :=client.Get("name")

	client.Set("age",11,0)

	m,_:=cmd.Result()

	fmt.Println(m)
}

func testHash()  {

	mp:=client.HGetAll("sunhongtao")
	m,_:=mp.Result()

	for index,value :=range m {

		fmt.Println(index,value)
	}

	for i:=0;i<len(m) ;i++  {

		fmt.Println(m["name"])
	}

	var p map[string] interface{}

	p=make(map[string]interface{})
	p["name"] = "jiangjun"
	p["age"]="12"


	client.HMSet("jiangjun",p)


	fmt.Println(m)
}

func testList()  {

	//cmd:=client.LPush("demo",11,2,15)


	l,_:=client.LRange("demo",0,10).Result()


	for index,v := range l{
		fmt.Println(index,v)
	}


	for i:=0;i<len(l);i++{
		fmt.Println(l[i])

	}

}

func testTran()  {



}

func testPS()  {

	//i,_:=client.Publish("messagequene","nihao").Result()
	//fmt.Println(i)

	for   {
		m,_:=client.Subscribe("demolist").ReceiveMessage()
		fmt.Println(m)
	}


}


func main() {

	ch :=make(chan int)

	client = redis.NewClient(&redis.Options{Addr:"127.0.0.1:8889"})

	r,e:=client.Ping().Result()
	if e!=nil {
		panic(e)
	}
	fmt.Println(r)


	// redis string

	//testString()


	//redis Hash

	//testHash()

	//redis list

	//testList()


	//redis S/P
	go testPS()

	<-ch
}
