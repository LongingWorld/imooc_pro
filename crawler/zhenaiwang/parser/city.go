package parser

import (
	"regexp"

	"GitHub/imooc_pro/crawler/engine"
)

const CityRequest = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParserResult  {
	reg := regexp.MustCompile(CityRequest)
	matchs := reg.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	for _,m :=range matchs  {
		name := string(m[2])
		result.Requests = append(result.Requests,
			engine.Request{URL: string(m[1]),
			ParserFunc:func(bytes []byte) engine.ParserResult {  //匿名函数  Closure闭包
				return ParserProfile(bytes,name)},
		})
		result.Items = append(result.Items, "User" + string(m[2]))
	}
	return result
}
