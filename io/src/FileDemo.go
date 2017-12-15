package main

import (
	"os"
	"fmt"
	"bytes"
	"io"
)

type Student struct {
	name string
}


func createFile() {
	file, e := os.Create("E://hello.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	file.WriteString("hello,world")
	fmt.Println(file.Name())

}

func readFile() {

	file, err := os.Open("e://hello.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	//方案一
	//by,_:=ioutil.ReadAll(file)
	//fmt.Println(string(by))

	//方案二
	buffer := bytes.NewBufferString("")
	var b = make([]byte, 100)
	for {
		n, e := file.Read(b)
		fmt.Println(n)

		if e != nil && e != io.EOF {
			panic(e)
		}

		if n > 0 && n <= len(b) {
			buffer.Write(b[:n])
		} else {
			break
		}
	}
	fmt.Println(buffer.String())
}

func writeFile() {
	file, err := os.Open("e://hello.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("nihao")
}

func main() {
	fmt.Println(os.Getpid())

	//createFile()

	readFile()
}
