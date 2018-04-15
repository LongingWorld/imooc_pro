package engine

import (
	"GitHub/imooc_pro/crawler/fetcher"
	"log"
)

func Run(Seeds ...Request)  {
	var requests []Request
	for _,r := range Seeds{
		requests = append(requests,r)
	}
	for len(requests)>0{
		req := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s",req.URL)
		body, err := fetcher.FetcherURL(req.URL)
		if err !=nil {
			log.Printf("Fetch error : fetch RUL %s： %v",req.URL,err)
			continue
		}
		parserResult := req.ParserFunc(body)
		requests = append(requests,parserResult.Requests...)
		for _,item := range parserResult.Items{
			log.Printf("Got items %v",item)
		}
	}
}
