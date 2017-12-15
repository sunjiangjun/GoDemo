package main

import (
	"bufio"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
)

/**

bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，
创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。

 */
func testBufio()  {
	//r:=strings.NewReader("hello,world")
	//rd:=bufio.NewReader(r)
	//
	//p:=make([]byte,30)
	//rd.Read(p)
	//fmt.Println(string(p))

	w:=bufio.NewWriter(bytes.NewBuffer(make([]byte,0)))
	w.WriteString(" nihao")
	w.Flush()
	fmt.Println(w.Buffered())
}

/**
   bytes 包 下的 reader  、 buffer
   buffer实现了 bufio包中 reader 、writer 接口
 */
func testBytes()  {

	r:=bytes.NewReader([]byte("hello"))
	b:=make([]byte,2)
	r.Read(b)
	fmt.Println(r.Size())
	fmt.Println(r.Len())

	buffer:=bytes.NewBufferString("hello,world")
	fmt.Println(buffer.Len())
	buffer.WriteString(" sunhongtao")
	fmt.Println(buffer.String())
}


/**

io包提供了对I/O原语的基本接口。本包的基本任务是包装这些原语已有的实现（如os包里的原语），
使之成为共享的公共接口，这些公共接口抽象出了泛用的函数并附加了一些相关的原语的操作。

因为这些接口和原语是对底层实现完全不同的低水平操作的包装，
除非得到其它方面的通知，客户端不应假设它们是并发执行安全的。


 */
func testIO()  {
	w:=bufio.NewWriter(bytes.NewBufferString("hello"))
	io.WriteString(w,"world")
	w.Flush()

	b,_:=ioutil.ReadAll(bytes.NewBufferString("nihao"))
	fmt.Println(string(b))

	s,_:=ioutil.ReadFile("E:\\hello.txt")
	fmt.Println(string(s))

	ioutil.WriteFile("E:\\hello.txt" ,[]byte("welcome"),7777)
}




func main() {

	//testBufio()

	//testBytes()

	testIO()

}
