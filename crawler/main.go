package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"regexp"
)

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	encoding,_,_ := charset.DetermineEncoding(bytes, "")
	return encoding
}

func printCityURL(context []byte)  {
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matchs := reg.FindAllSubmatch(context,-1)
	for _,m := range matchs  {
		//for _,subMatch := range m {
		//	fmt.Printf("%s ",subMatch)
		//}
		fmt.Printf("City:%s , CityURL:%s \n",m[2],m[1])
		fmt.Println()
	}
	fmt.Printf("Match number is :%d\n",len(matchs))
}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	//get encoding of the html
	encode := determineEncoding(resp.Body)

	//手工转码
	//utf8reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8reader := transform.NewReader(resp.Body, encode.NewDecoder())

	bytes, e := ioutil.ReadAll(utf8reader)
	if e != nil {
		panic(e)
	}

	printCityURL(bytes)
	//fmt.Printf("%s \n", bytes)
}

