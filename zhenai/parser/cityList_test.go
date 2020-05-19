package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//testurl := "https://zhenai.com/zhenghun"
	//body, err := fetcher.Fetch(testurl)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s", body)
	//err = ioutil.WriteFile("./bodyfile.txt", body, 0644)
	//if err != nil {
	//	panic(err)
	//}
	const citySize = 470
	body, err := ioutil.ReadFile("./bodyfile.html")
	if err != nil {
		panic(err)
	}
	cityList := ParserCityList(body)
	if len(cityList.Items) != citySize || len(cityList.Requests) != citySize {
		t.Fail()
	}
}
