package engine

import (
	"fetcher"
	"fmt"
	"log"
)

// Request 定义解析器的输出类
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 定义解析器的输出结果
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fecting url: %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		return ParserResult{}, fmt.Errorf("Fetcher: error fetching url %s: %v", r.URL, err)
	}
	parserResult := r.ParserFunc(body)
	return parserResult, nil
}
