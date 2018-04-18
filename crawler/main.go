package main

import (
	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/scheduler"
	"GitHub/imooc_pro/crawler/zhenaiwang/parser"
)

const requesturl = "http://www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		URL : requesturl,
		ParserFunc: parser.ParserCityList,
	})
}

