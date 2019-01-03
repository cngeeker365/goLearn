package main

import (
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T){
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil{
		panic(err)
	}

	req := worker.Request{
		Url: "http://m.zhenai.com/u/1320662004",
		Parser: worker.SerializedParser{
			FuncName: config.ParseProfile,
			Args: "微微一笑",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	}else {
		fmt.Println(result)
	}
}