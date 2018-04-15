package main

import (
	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/zhenaiwang/parser"
)

const requesturl = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		URL : requesturl,
		ParserFunc: parser.ParserCityList,
	})
}

