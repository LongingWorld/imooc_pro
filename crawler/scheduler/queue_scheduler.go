package scheduler

import "GitHub/imooc_pro/crawler/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan chan  chan engine.Request
}

func (s *QueueScheduler) Run() {
	go func() {
	}()
}

func (s *QueueScheduler) WorkReady(r chan engine.Request) {
	s.workChan <- r
}

func (s *QueueScheduler) Summit(r engine.Request) {
	s.requestChan <- r
}

func (*QueueScheduler) ConfigureMasterWorkChan(chan engine.Request) {
}


