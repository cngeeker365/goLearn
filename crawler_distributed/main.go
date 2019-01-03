package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/persist/client"
	"fmt"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if nil != err {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: 		&scheduler.QueuedScheduler{},
		WorkerCount: 	10,
		ItemChan: 		itemChan,
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