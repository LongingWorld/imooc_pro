package scheduler

import "GitHub/imooc_pro/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan chan  chan engine.Request
}

func (s *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//TODO: scheduler create a goroutine to send request to worker
func (s *QueueScheduler) Run() {
	s.workChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ)>0 && len(workerQ)>0{
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			//select
			select{
			case r := <-s.requestChan :
				requestQ = append(requestQ,r)
			case w := <-s.workChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (s *QueueScheduler) WorkReady(r chan engine.Request) {
	s.workChan <- r
}

func (s *QueueScheduler) Summit(r engine.Request) {
	s.requestChan <- r
}




