package scheduler

import "GitHub/imooc_pro/crawler/engine"

type SimpleScheduler struct {  //实现Scheduler接口
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workChan =c
}

func (s *SimpleScheduler) Summit(r engine.Request) {
	//send request down to worker chan
	s.workChan <- r
}



