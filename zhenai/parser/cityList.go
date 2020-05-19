package parser

import (
	"engine"
	"regexp"
	"strings"
)

const cityListPattern = `{"linkContent":"([^"]*)","linkURL":"(http:\\u002F\\u002Fwww.zhenai.com\\u002Fzhenghun\\u002F[a-z0-9]+[^"]*)"}`

// ParserCityList 解析城市列表信息
func ParserCityList(contents []byte) engine.ParserResult {
	result := engine.ParserResult{}
	reg := regexp.MustCompile(cityListPattern)
	rets := reg.FindAllStringSubmatch(string(contents), -1)
	limit := 10
	for _, ret := range rets {
		result.Items = append(result.Items, "City "+ret[1])
		result.Requests = append(result.Requests, engine.Request{
			URL:        strings.Replace(ret[2], "\\u002F", "/", -1),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}

// ParserPersonList 解析城市中的人物列表信息
//func ParserPersonList(contents []byte) engine.ParserResult {
//
//}
