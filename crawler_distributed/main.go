package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
	"awesomeProject/crawler_distributed/config"
	itemsaver"awesomeProject/crawler_distributed/persist/client"
	"awesomeProject/crawler_distributed/rpcsupport"
	worker "awesomeProject/crawler_distributed/worker/client"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	workerHosts	= flag.String("worker_hosts","","worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", *itemSaverHost))
	if nil != err {
		panic(err)
	}

	pool:=createClientPool(strings.Split(*workerHosts,","))

	processor:= worker.CreateProcessor(pool)
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

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _,h:=range hosts{
		client,err:=rpcsupport.NewClient(h)
		if err==nil{
			clients=append(clients, client)
		}else{
			log.Printf("error connecting to %s: %v\n",h,err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients{
				out <- client
			}
		}
	}()

	return out
}