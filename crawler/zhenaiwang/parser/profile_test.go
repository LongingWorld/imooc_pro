package parser

import (
	"io/ioutil"
	"testing"

	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/model"
)

func TestParseProfile(t *testing.T) {
	bytes, e := ioutil.ReadFile("profile_test_data.html")
	if e != nil{
		panic(e)
	}
	result := ParseProfile(bytes,"http://album.zhenai.com/u/110409917","妈妈催我找对象")
	if len(result.Items) != 1 {
		t.Errorf("Items  should contain 1 element;but was %v",result.Items)
	}
	profile := result.Items[0]

	expectedProfile :=engine.Item{
		URL:"http://album.zhenai.com/u/110409917",
		Type:"zhenai",
		Id:"110409917",
		Payload: model.Profile{
			Name:"妈妈催我找对象",
			Gender:"女",
			Age:20,
			Height:165,
			Weight:0,
			Income:"3001-5000元",
			Marriage:"未婚",
			Education:"中专",
			Occupation:"四川成都",
			Hokou:"--",
			Xinzuo:"牡羊座",
			House:"--",
			Car:"未购车",
		},
	}

	if expectedProfile != profile{
		t.Errorf("expected :%v;but was %v",expectedProfile,profile)
	}
}
