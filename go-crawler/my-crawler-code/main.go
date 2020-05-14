package main

import (
	"my-crawler-code/engine"
	"my-crawler-code/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		PsrserFunc: parser.ParseCityList,
	})
}
