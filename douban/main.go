package main

import (
	"doubanparser"
	"engine"
	"persist"
	"scheduler"
)

const SEEDURL = "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
const HeadURL = "https://book.douban.com"
const Prefix = "https://book.douban.com"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		NumWorkers: 1,
		ItemChan:persist.ItemSaver(),
	}
	e.Run(engine.Request{
		CurURL: SEEDURL,
		ParserFunc: func(content []byte) engine.ParserResult {
			return doubanparser.ParseBookTag(content, Prefix, HeadURL)
		},
	})
}
