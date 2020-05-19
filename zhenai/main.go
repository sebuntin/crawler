package main

import (
	"engine"
	"math/rand"
	"parser"
	"time"
)

// <a href="/c101010100-p100116/" ka="sel-city-101010100">北京</a>
const SEEDURL = "https://movie.douban.com/"
//const PREFIX = "https://www.zhipin.com/"




func main() {
	// 开启爬虫任务，先传入种子任务
	engine.Run(engine.Request{URL:SEEDURL, ParserFunc:parser.ParserCityList})
}


