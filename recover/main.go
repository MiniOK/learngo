package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error code：", err)
		} else {
			//  Sprintf() 只是进行格式化， 不会打印
			panic(fmt.Sprintf("I don`t kown what to do: %v", r))
		}
	}()
	// panic(errors.New("this is an error"))
	// b := 0
	// a := 5 / b
	// fmt.Println(a)
	panic(123)
}

func main() {
	tryRecover()
}
