package main

import (
	"fmt"
	"sync"
)


/**

sync 包中实现了两个关于锁的数据类型，sync.Mutex 和 sync.RWMutex。[ 互斥锁 mutex 是独占型，只能 lock 一次， unlock 一次，
然后才能继续 lock 否则阻塞。
读写互斥锁 reader-writer mutex 是所有的 reader 共享一把锁或是一个 writer 独占一个锁，
如果一个 reader lock 到锁了， 其他的 reader 还可以 lock 但是 writer 不能 lock 。 ]

对于 sync.Mutex 或是 sync.RWMutex 类型的变量 mutex 来说，假定 n < m，
对于 mutex.Unlock() 的第 n 次调用在 mutex.Lock() 的第 m 次调用返回之前发生。
[ 对于一个 mutex 来说，lock 一下，第二次 lock 会阻塞，只有 unlock 一下才可以继续 lock，就是这个意思。
然而 unlock 一个没有 lock 的 mutex 会怎么样呢？error ! ]

 */

var l sync.Mutex

func foo()  {

	fmt.Println("hello,world")

	l.Unlock()
}

func main() {
	fmt.Println("1")
	l.Lock()
	go foo()

	fmt.Println(3)
	l.Lock()
	fmt.Println("2")
}
