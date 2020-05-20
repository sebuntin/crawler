package engine

import (
	"fetcher"
	"fmt"
	"log"
)

// Request 定义解析器的输出类
type Request struct {
	CurURL     string
	RefURL     string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 定义解析器的输出结果
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

type Item struct {
	Url     string
	Id      string
	PayLoad interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fecting url: %s", r.CurURL, r.RefURL)
	body, err := fetcher.Fetch(r.CurURL, r.RefURL)
	if err != nil {
		return ParserResult{}, fmt.Errorf("Fetcher: error fetching url %s: %v", r.CurURL, err)
	}
	parserResult := r.ParserFunc(body)
	return parserResult, nil
}
