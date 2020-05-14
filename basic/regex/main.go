package main

import (
	"fmt"
	"regexp"
)

const text = `<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/zibo"class= "">淄博</a>`

func main() {
	// re := regexp.MustCompile(`<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-z]"[^>]*>[^<]+</a>`)
	re := regexp.MustCompile(`<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	matches := re.FindAllString(text, -1)
	fmt.Println(matches)
}
