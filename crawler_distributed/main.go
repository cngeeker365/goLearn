package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	itemsaver"awesomeProject/crawler_distributed/persist/client"
	worker "awesomeProject/crawler_distributed/worker/client"
	"fmt"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if nil != err {
		panic(err)
	}

	processor, err:= worker.CreateProcessor()
	if err!=nil{
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: 		&scheduler.QueuedScheduler{},
		WorkerCount: 	10,
		ItemChan: 		itemChan,
		RequestProcessor:processor,
	}
	//从首页进行爬取
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),
	})

	//从上海市开始爬取
	//e.Run(engine.Request{
	//	Url: 		"http://m.zhenai.com/zhenghun/bozhou",
	//	ParserFunc: parser.ParseCity,
	//})
}