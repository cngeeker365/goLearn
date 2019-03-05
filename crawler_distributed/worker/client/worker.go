package client

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{

	return func(req engine.Request) (result engine.ParseResult, e error) {
		sReq:=worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <- clientChan
		err:=c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{},err
		}
		return worker.DeserializeResult(sResult),nil
	}
}
