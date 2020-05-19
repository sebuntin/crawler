package engine

import (
	"log"
	"time"
)

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkReady(chan Request)
	Run()
}

type ConcurrentEngine struct {
	Scheduler  Scheduler
	NumWorkers int
}

func (e ConcurrentEngine) Run(seed ...Request) {
	//in := make(chan Request, 10)
	out := make(chan ParserResult)
	//e.Scheduler.ConfigureMasterWorkChan(in)
	e.Scheduler.Run()
	// 将种子请求放入RequestChan中
	for _, r := range seed {
		e.Scheduler.Submit(r)
	}
	// 创建worker
	for i := 0; i < e.NumWorkers; i++ {
		creatWorker(e.Scheduler, out)
	}
	var count int
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("item: %v, count: %d\n", item, count)
			count ++
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func creatWorker(scheduler Scheduler, out chan ParserResult) {
	in := make(chan Request)
	go func() {
		for {
			scheduler.WorkReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			time.Sleep(time.Second * 1)
			out <- result
		}
	}()
}
