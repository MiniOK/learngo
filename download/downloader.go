package main

import (
	"fmt"
	"hello.go/download/infra"
)

func getReriever() infra.Retriever {
	return infra.Retriever{}
}

func main() {
	fmt.Println(getReriever().Get("https://www.imooc.com"))
}
