package parser

import (
	"io/ioutil"
	"log"
	"testing"

)

func TestParserCityList(t *testing.T) {
	//contents, err := fetcher.FetcherURL("citylist_test_data.html")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserCityList(contents)
	//fmt.Printf("%s",contents)
	const resultSize = 470
	if len(result.Items) != resultSize {
		log.Printf("result should have %d requsts;but had %d",resultSize,len(result.Items))
	}
	expectedURL := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedItems := []string{
		"City 阿坝","City 阿克苏","City 阿拉善盟",
		//"","","",
	}

	for i,urls := range expectedURL{
		if urls != result.Requests[i].URL {
			t.Errorf("expected URL is %s ;but got %s",urls,result.Requests[i].URL)
		}
	}
	for i,item := range expectedItems{
		if item != result.Items[i].(string) {  //.(string)将item类型转换为string
			 t.Errorf("expected item is %s;but got %s",item,result.Items[i].(string))
		}
	}
}
