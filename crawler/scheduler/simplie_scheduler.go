package scheduler

import "GitHub/imooc_pro/crawler/engine"

type SimpleScheduler struct {  //实现Scheduler接口
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
}


func (s *SimpleScheduler) Summit(r engine.Request) {
	//send request down to worker chan
	go func() {s.workChan <- r}()
}



