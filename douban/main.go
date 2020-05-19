package main

import (
	"engine"
	"doubanparser"
	"scheduler"
)

const SEEDURL = "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		NumWorkers: 10,
	}
	e.Run(engine.Request{
		URL: SEEDURL,
		ParserFunc: func(content []byte) engine.ParserResult {
			return doubanparser.ParseBookTag(content, "https://book.douban.com")
		},
	})
}
