package engine

import (
	"GitHub/imooc_pro/crawler/fetcher"
	"log"
)

// SimpleEngine
type SimpleEngine struct {}

// Run
func (e SimpleEngine)Run(Seeds ...Request)  {
	var requests []Request
	for _,r := range Seeds{
		requests = append(requests,r)
	}
	for len(requests)>0{
		req := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s",req.URL)

		parserResult, e := Worker(req)
		if e != nil {
			continue
		}
		requests = append(requests,parserResult.Requests...)
		for _,item := range parserResult.Items{
			log.Printf("Got items %v",item)
		}
	}
}

func Worker(r Request) (ParserResult,error)  {   //fetcher and parser  reconsitution
	log.Printf("Fetching %s",r.URL)
	body, err := fetcher.FetcherURL(r.URL)
	if err !=nil {
		log.Printf("Fetch error : fetch RUL %s： %v",r.URL,err)
		return ParserResult{},err
	}
	return r.ParserFunc(body),nil
}