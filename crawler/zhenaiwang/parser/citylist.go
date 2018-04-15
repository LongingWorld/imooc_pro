package parser

import (
	"GitHub/imooc_pro/crawler/engine"
	"regexp"
)

const CityListRequest = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult {
	reg := regexp.MustCompile(CityListRequest)
	matchs := reg.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matchs {
		//fmt.Printf("City:%s , CityURL:%s \n",m[2],m[1])
		//fmt.Println()
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{URL: string(m[1]), ParserFunc: engine.NilParserFunc})
	}
	//fmt.Printf("Match number is :%d\n",len(matchs))
	return result
}
