package engine

import (
	"log"
	"my-crawler-code/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetcher url: %s", r.Url)
		// 提取 页面text
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			//	 为了爬虫不挂掉,我们把错误给忽略掉
			log.Printf("Fetcher error fetching url %s %v", r.Url, err)
			continue
		}
		// 解析页面
		parserResult := r.PsrserFunc(body)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}
