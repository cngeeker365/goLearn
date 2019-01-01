package engine

type ConcurrentEngine struct {
	Scheduler 		Scheduler
	WorkerCount		int
	ItemChan 		chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request //每个scheduler应该知道给出来的chan是共用，还是独立的（即每个worker自己的）
	//ConfigMasterWorkerChan(chan Request)
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	//多个worker共用一个in和out
	//in := make(chan Request)
	out := make(chan ParseResult)
	//e.Scheduler.ConfigMasterWorkerChan(in)
	e.Scheduler.Run()

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for i:=0;i<e.WorkerCount;i++{
		//createWorker(in, out)
		//createWorkerQ(out, e.Scheduler)
		createWorkerComm(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items{
			//log.Printf("Got item #%d: %v\n", itemCount, item)
			//itemCount++
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorkerComm(in chan Request, out chan ParseResult, ready ReadyNotifier){
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)

			request := <- in
			result, err := worker(request)
			if err !=nil{
				continue
			}
			out <- result
		}
	}()
}

func createWorkerQ(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			s.WorkerReady(in)

			request := <- in
			result, err := worker(request)
			if err !=nil{
				continue
			}
			out <- result
		}
	}()
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// tell scheduler i'm ready
			request := <- in
			result, err := worker(request)
			if err !=nil{
				continue
			}
			out <- result
		}
	}()
}