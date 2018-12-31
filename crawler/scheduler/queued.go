package scheduler

import "awesomeProject/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//func (s *QueuedScheduler) ConfigMasterWorkerChan(chan engine.Request) {
//	panic("implement me")
//}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			// r&w 各自独立发生，因此要用select来处理
			// r 收到，但无法发送给 w，因为不知道w是谁，因此要将收到的 r 放入队列
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
