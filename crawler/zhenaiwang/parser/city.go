package parser

import (
	"regexp"

	"GitHub/imooc_pro/crawler/engine"
)

var(
	CityRequest = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	CityUrlRequest = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(contents []byte) engine.ParserResult  {
	//reg := regexp.MustCompile(CityRequest)
	matchs := CityRequest.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m :=range matchs  {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests,
			engine.Request{URL: string(m[1]),
				ParserFunc:func(bytes []byte) engine.ParserResult {  //匿名函数  Closure闭包
					return ParseProfile(bytes,url,name)},
			})
		//result.Items = append(result.Items, "User" + string(m[2]))
	}

	//
	matchs = CityUrlRequest.FindAllSubmatch(contents,-1)
	for _,m :=range matchs {
		//result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{URL: string(m[1]),
				ParserFunc:ParseCity,
			})
	}
	return result
}
