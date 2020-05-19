package scheduler

import "engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.RequestChan <- r
}

func (q *QueuedScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	panic("not implement")
}

func (q *QueuedScheduler) WorkReady(c chan engine.Request) {
	q.WorkerChan <- c
}

func (q *QueuedScheduler) Run() {
	q.RequestChan = make(chan engine.Request)
	q.WorkerChan = make(chan chan engine.Request)
	go func() {
		requestQ := make([]engine.Request, 0, 1000)
		workerQ := make([]chan engine.Request, 0, 1000)
		for {
			var activeRequeset engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequeset = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-q.WorkerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequeset:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
