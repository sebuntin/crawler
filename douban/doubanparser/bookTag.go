package doubanparser

import (
	"engine"
	"regexp"
)

var BookTagReg = regexp.MustCompile(`<td><a href="(/tag/[^"]*)">([^<]*)</a><b>\([\d]+\)</b></td>`)

func ParseBookTag(content []byte, prefix string) engine.ParserResult {
	result := engine.ParserResult{}
	submatchs := BookTagReg.FindAllSubmatch(content, -1)
	//limit := 2
	for _, item := range submatchs {
		result.Items = append(result.Items, string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:        prefix + string(item[1]),
			ParserFunc: ParseBookList,
		})
		//limit --
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
