package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Summit(Request)
	WorkerChan() chan Request //每一个worker对应一个channel
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

//负责通过建立goroutine将Engine分发的Request输送给worker进行工作
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) //创建处理请求的goroutine：接收Request(请求) 输出ParserResult(解析结果)
	}

	for _, r := range seeds {
		if isUrlDuplicate(r.URL) {
			log.Printf("Duplicate request: %s", r.URL)
			continue
		}
		e.Scheduler.Summit(r) //将Seeds Request分发给channel in(接收者)
	}

	//countItems :=0
	//profileCount := 0
	for {
		result := <-out //等待接收输出

		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
			//if _,ok := item.(model.Profile);ok {
			//	log.Printf("Got item:#%d %v",profileCount,item)
			//	profileCount++
			//}
		}
		for _, request := range result.Requests {
			if isUrlDuplicate(request.URL) {
				log.Printf("Duplicate request: %s", request.URL)
				continue
			}
			e.Scheduler.Summit(request) //将第二层请求分发给channel in(接收者)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) { //创建goroutine
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in //等待接收
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result //等待输出
		}
	}()
}

var urlStore = make(map[string]bool)

func isUrlDuplicate(url string) bool {
	if urlStore[url] {
		return true
	}
	urlStore[url] = true
	return false
}
