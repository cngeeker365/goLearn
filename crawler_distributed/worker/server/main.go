package main

import (
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"awesomeProject/crawler_distributed/worker"
	"fmt"
	"log"
)

func main()  {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
