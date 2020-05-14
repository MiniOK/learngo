package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i) for 循环中的 i 没有泄露成全局变量
}
