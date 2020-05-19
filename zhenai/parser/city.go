package parser

import (
	"engine"
	"regexp"
)

const CityPattern = `<th><a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]*)</a></th>`

func ParseCity(contents []byte) engine.ParserResult {
	result := engine.ParserResult{}
	reg := regexp.MustCompile(CityPattern)
	submatch := reg.FindAllStringSubmatch(string(contents), -1)
	for i := range submatch {
		result.Items = append(result.Items, "User "+submatch[i][2])
		result.Requests = append(result.Requests, engine.Request{
			URL: submatch[i][1],
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, submatch[i][2])
			},
		})
	}
	return result
}
