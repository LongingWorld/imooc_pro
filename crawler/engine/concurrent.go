package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Summit(Request)
	Run()
	WorkReady(chan Request)
	ConfigureMasterWorkChan(chan Request)
}

//负责通过建立goroutine将Engine分发的Request输送给worker进行工作
func (e *ConcurrentEngine)Run(seeds ...Request)  {

	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkChan(in)

	for i:=0;i<e.WorkerCount;i++  {
		createWorker(in,out)   //创建处理请求的goroutine：接收Request(请求) 输出ParserResult(解析结果)
	}

	for _,r := range seeds  {
		e.Scheduler.Summit(r)   //将Seeds Request分发给channel in(接收者)
	}

	countItems :=0
	for{
		result := <-out    //等待接收输出
		for _,item :=range result.Items{
			log.Printf("Got item:#%d %v",countItems,item)
			countItems++
		}
		for _,request := range result.Requests{
			e.Scheduler.Summit(request)  //将第二层请求分发给channel in(接收者)
		}
	}
}

func createWorker(in chan Request,out chan ParserResult)  {  //创建goroutine
	go func() {
		for   {
			request := <- in   //等待接收
			result, err := Worker(request)
			if err !=nil {
				continue
			}
			out <- result   //等待输出
		}
	}()
}
