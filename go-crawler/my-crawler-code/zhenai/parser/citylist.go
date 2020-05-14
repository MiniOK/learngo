package parser

import (
	"my-crawler-code/engine"
	"regexp"
)

// 各个城市的解析
const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	// 声明一个 result
	result := engine.ParserResult{}
	for _, m := range matches {
		// 返回城市的名字
		result.Items = append(result.Items, "City"+string(m[2]))
		// 返回城市的 URL
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			PsrserFunc: ParseCity,
		})
	}
	return result
}
