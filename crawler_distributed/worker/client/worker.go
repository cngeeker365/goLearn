package client

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor,error) {
	client,err:= rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil{
		return nil, err
	}
	return func(req engine.Request) (result engine.ParseResult, e error) {
		sReq:=worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err:=client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult),nil
	}, nil
}
