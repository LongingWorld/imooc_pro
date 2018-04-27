package persist

import (
	"context"
	"encoding/json"
	"testing"

	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
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
	err := save(expectedProfile)
	if err != nil {
		panic(err)
	}

	//TODO:Try to start up elastic search. Here using docker go client.
	cli,err := elastic.NewClient(
		elastic.SetURL("http://192.168.99.100:9200"),
		//Must turn off sniff in docker
		elastic.SetSniff(false)	)
	if err != nil{
		panic(err)
	}
	result, err := cli.Get().
		Index("dating_profile").
		Type(expectedProfile.Type).
		Id(expectedProfile.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s",result.Source)
	var actual engine.Item
	e := json.Unmarshal(*result.Source, &actual)
	if e != nil {
		panic(e)
	}

	actualProfile,err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}

	actual.Payload = actualProfile

	if expectedProfile != actual {
		t.Errorf("Got %v ; expected %v",actual,expectedProfile)
	}
}
