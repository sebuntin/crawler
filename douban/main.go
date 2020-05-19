package main

import (
	"doubanparser"
	"engine"
	"scheduler"
)

const SEEDURL = "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.SimpleScheduler{},
		NumWorkers: 10,
	}
	e.Run(engine.Request{
		URL: SEEDURL,
		ParserFunc: func(content []byte) engine.ParserResult {
			return doubanparser.ParseBookTag(content, "https://book.douban.com")
		},
	})
}
