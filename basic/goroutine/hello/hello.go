package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

}

/*
	声明管道的格式
var 变量 chan 元素类型
var sh1 chan int
var ch2 chan string
var ch3 chan []int
	创建channel
申明channel 之后使用make进行初始化
ch1 := make(chan int)

channel 有 send receive 和 close 三种操作
发送和接收使用 <- 符号
x := <-ch 用 x 接收值
<- ch 接收值 忽略结果
close(ch) 关闭通道

无缓冲通道又称为阻塞通道

deadlock 怎么解决
使用一个goroutine 去接收

使用无缓冲通道将导致发送和接收的goroutine 同步化，因此，无缓冲通道也称为同步通道
判断 通道是否关闭，统称使用的 for range
单向通道
chan<- int 只能发送
<-chan int 只能接收

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个 goroutine 可以访问共享资源
使用 sync 包的 Mutex 类型来实现，

读写互斥锁


*/
