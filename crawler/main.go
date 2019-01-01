package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: persist.ItemSaver(),
	}
	//从首页进行爬取
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})

	//从上海市开始爬取
	//e.Run(engine.Request{
	//	Url: 		"http://m.zhenai.com/zhenghun/bozhou",
	//	ParserFunc: parser.ParseCity,
	//})
}