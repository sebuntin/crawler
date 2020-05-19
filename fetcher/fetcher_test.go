package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	doubanurl := "https://movie.douban.com/"
	body, err := Fetch(doubanurl)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(body))
}
