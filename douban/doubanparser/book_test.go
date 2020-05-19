package doubanparser

import (
	"fetcher"
	"fmt"
	"testing"
	"time"
)

func TestParseBookInfo(t *testing.T) {
	URLList := []string{"https://book.douban.com/subject/1770782/",
		"https://book.douban.com/subject/10554308/",
		"https://book.douban.com/subject/25862578/"}
	for _, doubanURL := range URLList {
		body, err := fetcher.Fetch(doubanURL)
		if err != nil {
			t.Fail()
		}
		info := ParseBookInfo(body, "")
		fmt.Printf("book info: %#v\n", info.Items[0])
		time.Sleep(time.Second * 2)
	}
}
