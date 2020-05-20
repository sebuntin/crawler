package doubanparser

import (
	"engine"
	"regexp"
)

var BookTagReg = regexp.MustCompile(`<td><a href="(/tag/[^"]*)">([^<]*)</a><b>\([\d]+\)</b></td>`)

func ParseBookTag(content []byte, prefix, url string) engine.ParserResult {
	result := engine.ParserResult{}
	submatchs := BookTagReg.FindAllSubmatch(content, -1)
	//limit := 2
	for _, m:= range submatchs {
		item := string(m[2])
		suffix := string(m[1])
		result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			CurURL:     prefix + suffix,
			RefURL:     url,
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseBookList(c, url)
			},
		})
		//limit --
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
