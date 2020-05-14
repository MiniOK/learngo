package parser

import (
	"my-crawler-code/engine"
	"regexp"
)

// 解析城市
const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)">[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	// 声明一个 result
	result := engine.ParserResult{}
	for _, m := range matches {
		// 返回城市的名字
		result.Items = append(result.Items, "User"+string(m[2]))
		// 返回城市的 URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			PsrserFunc: engine.NilParser,
		})

	}

	return result
}
