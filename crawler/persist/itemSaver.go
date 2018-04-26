package persist

import (
	"context"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
)



/*老师关于ItemSaver的解释
channel在go语言中是一等公民。也就是可以像其他比如整数，字符串这样传来传去。
这里调用ItemSaver之后他生成并且返回一个channel，背后开了一个goroutine来从这个channel接收并且处理数据。
由于之前说到的函数闭包的原因，这个goroutine，以及里面所引用的变量，在ItemSaver返回之后仍然继续在工作。
然后拿着返回值也就是拿着这个channel的人可以往里放item。这样的逻辑很合理。*/

func ItemSaver() chan interface{}  {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("ItemSaver got #%d: %v",itemCount,item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver :error Saving item %v :%v",item,err)
			}
		}
	}()
	return  out
}

func save(item interface{}) (str string,err error)  {
	client, err := elastic.NewClient(/*url.URL{"http://192.168.99.100:9200"},*/
		elastic.SetURL("http://192.168.99.100:9200"),
		//Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return "",err
	}
	response, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "",err
	}
	fmt.Printf("%+v",response)

	return response.Id,err
}
