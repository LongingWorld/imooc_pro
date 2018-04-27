package main

import (
	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/persist"
	"GitHub/imooc_pro/crawler/scheduler"
	"GitHub/imooc_pro/crawler/zhenaiwang/parser"
)

const requesturl = "http://www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueueScheduler{},
		WorkerCount:100,
		ItemChan:persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	URL : requesturl,
	//	ParserFunc: parser.ParserCityList,
	//})
	e.Run(engine.Request{
		URL:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc:parser.ParseCity,
	})
}

