package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount	int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	//多个worker共用一个in和out
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigMasterWorkerChan(in)

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for i:=0;i<e.WorkerCount;i++{
		createWorker(in, out)
	}

	itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items{
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <- in
			result, err := worker(request)
			if err !=nil{
				continue
			}
			out <- result
		}
	}()
}