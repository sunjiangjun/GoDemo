package main

import (
	"net"
	"fmt"
	"time"
)

/**

conn.Read的行为特点

1.1、Socket中无数据
连接建立后，如果对方未发送数据到socket，接收方(Server)会阻塞在Read操作上，这和前面提到的“模型”原理是一致的。执行该Read操作的goroutine也会被挂起。runtime会监视该socket，直到其有数据才会重新
调度该socket对应的Goroutine完成read。由于篇幅原因，这里就不列代码了，例子对应的代码文件：go-tcpsock/read_write下的client1.go和server1.go。

1.2、Socket中有部分数据
如果socket中有部分数据，且长度小于一次Read操作所期望读出的数据长度，那么Read将会成功读出这部分数据并返回，而不是等待所有期望数据全部读取后再返回。

1.3、Socket中有足够数据
如果socket中有数据，且长度大于等于一次Read操作所期望读出的数据长度，那么Read将会成功读出这部分数据并返回。这个情景是最符合我们对Read的期待的了：Read将用Socket中的数据将我们传入的slice填满后返回：n = 10, err = nil

1.4、Socket关闭
如果client端主动关闭了socket，那么Server的Read将会读到什么呢？
这里分为“有数据关闭”和“无数据关闭”。

有数据关闭是指在client关闭时，socket中还有server端未读取的数据。当client端close socket退出后，server依旧没有开始Read，10s后第一次Read成功读出了所有的数据，当第二次Read时，由于client端 socket关闭，Read返回EOF error

无数据关闭情形下的结果，那就是Read直接返回EOF error

1.5、读取操作超时
有些场合对Read的阻塞时间有严格限制，在这种情况下，Read的行为到底是什么样的呢？在返回超时错误时，是否也同时Read了一部分数据了呢？
不会出现“读出部分数据且返回超时错误”的情况

 */


 /**
  conn.Write的行为特点

2.1、成功写
前面例子着重于Read，client端在Write时并未判断Write的返回值。所谓“成功写”指的就是Write调用返回的n与预期要写入的数据长度相等，且error = nil。这是我们在调用Write时遇到的最常见的情形，这里不再举例了

2.2、写阻塞
TCP连接通信两端的OS都会为该连接保留数据缓冲，一端调用Write后，实际上数据是写入到OS的协议栈的数据缓冲的。TCP是全双工通信，因此每个方向都有独立的数据缓冲。当发送方将对方的接收缓冲区以及自身的发送缓冲区写满后，Write就会阻塞

2.3、写入部分数据
Write操作存在写入部分数据的情况。没有按照预期的写入所有数据。这时候循环写入便是
  */


func handleConn(conn net.Conn)  {

	fmt.Println("handler conn....")
	defer  conn.Close()
	b:=make([]byte,1024)
	conn.Read(b)
	fmt.Println(string(b))
	conn.Write([]byte("hello,I have receiver message"))
}

/**

Goroutine safe

基于goroutine的网络架构模型，存在在不同goroutine间共享conn的情况，那么conn的读写是否是goroutine safe的呢？在深入这个问题之前，我们先从应用意义上来看read操作和write操作的goroutine-safe必要性。

对于read操作而言，由于TCP是面向字节流，conn.Read无法正确区分数据的业务边界，因此多个goroutine对同一个conn进行read的意义不大，goroutine读到不完整的业务包反倒是增加了业务处理的难度。对与Write操作而言，倒是有多个goroutine并发写的情况。

每次Write操作都是受lock保护，直到此次数据全部write完。因此在应用层面，要想保证多个goroutine在一个conn上write操作的Safe，需要一次write完整写入一个“业务包”；一旦将业务包的写入拆分为多次write，那就无法保证某个Goroutine的某“业务包”数据在conn发送的连续性。

同时也可以看出即便是Read操作，也是lock保护的。多个Goroutine对同一conn的并发读不会出现读出内容重叠的情况，但内容断点是依
runtime调度来随机确定的。存在一个业务包数据，1/3内容被goroutine-1读走，另外2/3被另外一个goroutine-2读 走的情况。比如一个完整包：world，当goroutine的read slice size < 5时，存在可能：一个goroutine读到 “worl”,另外一个goroutine读出”d”

 */


func main() {

	l,e:=net.Listen("tcp",":8888")

	if e!=nil {
		panic(e)
	}

	defer l.Close()

	for  {
		conn,err:=l.Accept()
		if err !=nil{
			fmt.Println("error")
			break
		}

		time.Sleep(time.Second * 10)
		go handleConn(conn)
	}

}
