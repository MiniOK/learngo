package main

import (
	"fmt"
	"io"
	"time"

	"hello.go/retriever/mock"
	real2 "hello.go/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "junchuan",
			"age":  "ershi",
		})
}

// 接口组合 也是使用者定义接口的一种方法
type RetrieverPoster interface {
	Retriever
	Poster
}

// 定义一个 session 方法
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	mockretriever := mock.Retriever{"This is a fake imooc.com"}
	inspect(r)
	r = &real2.Retriever{UserAgent: "Maliza 5.0", TimeOut: time.Second}
	inspect(r)

	fmt.Println("Try a session")
	fmt.Println(session(&mockretriever))
	io.WriteCloser()

}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents：", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
