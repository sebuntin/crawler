package doubanparser

import (
	"engine"
	"regexp"
)

//<a href="https://book.douban.com/subject/1770782/" title="追风筝的人" onclick="moreurl(this,{i:'1',query:'',subject_id:'1770782',from:'book_subject_search'})">
//
//追风筝的人
//
//
//
//
//</a>

var BookListRe = regexp.MustCompile(`<a href="(https://book.douban.com/subject/[\d]+/)" title="([^"]*)"`)

// 获取当前标签下的所有图书
func ParseBookList(content []byte, url string) engine.ParserResult {
	result := engine.ParserResult{}
	submatchs := BookListRe.FindAllSubmatch(content, -1)
	//limit := r
	for i := range submatchs {
		// 保证闭包数据的正确性，需要使用局部变量逃逸
		itemURL := string(submatchs[i][1])
		itemName := string(submatchs[i][2])
		result.Items = append(result.Items, string(submatchs[i][2]))
		result.Requests = append(result.Requests, engine.Request{
			CurURL: itemURL,
			RefURL: url,
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseBookInfo(c, itemName, itemURL)
			},
		})
		//limit --
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
