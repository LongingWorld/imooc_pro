package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Summit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request)  {

	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkChan(in)

	for i:=0;i<e.WorkerCount;i++  {
		createWorkder(in,out)
	}

	for _,r := range seeds  {
		e.Scheduler.Summit(r)
	}

	for{
		result := <-out
		for _,item :=range result.Items{
			log.Printf("Got item:%v",item)
		}
		for _,request := range result.Requests{
			e.Scheduler.Summit(request)
		}
	}
}

func createWorkder(in chan Request,out chan ParserResult)  {  //创建goroutine
	go func() {
		for   {
			request := <- in
			result, err := Worker(request)
			if err !=nil {
				continue
			}
			out <- result
		}
	}()
}
