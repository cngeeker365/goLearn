package scheduler

import "awesomeProject/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	go func() {s.workerChan <- r}()
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request)  {
	s.workerChan = c
}


