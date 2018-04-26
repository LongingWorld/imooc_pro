package persist

import (
	"context"
	"encoding/json"
	"testing"

	"GitHub/imooc_pro/crawler/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	expectedProfile := model.Profile{
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
	}
	id,err := save(expectedProfile)
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
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s",result.Source)
	var actual model.Profile
	e := json.Unmarshal(*result.Source, &actual)
	if e != nil {
		panic(e)
	}
	if expectedProfile != actual {
		t.Errorf("Got %v ; expected %v",actual,expectedProfile)
	}
}
