package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	//const cityurl = "http://www.zhenai.com/zhenghun/xianning"
	//body, err := fetcher.Fetch(cityurl)
	//if err != nil {
	//	panic(err)
	//}
	//err = ioutil.WriteFile("./cityfile.html", body, 0644)
	//if err != nil {
	//	panic(err)
	//}
	body, err := ioutil.ReadFile("./cityfile.html")
	if err != nil {
		panic(err)
	}
	personList := ParseCity(body)
	for _, p := range personList.Items {
		fmt.Println(p)
	}
}
