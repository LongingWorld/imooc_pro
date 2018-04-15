package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		log.Printf("Fetcher error:%v",e)
		return unicode.UTF8  //return default encoding.Encoding
	}
	encoding,_,_ := charset.DetermineEncoding(bytes, "")
	return encoding
}

func FetcherURL(url string) ([]byte ,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error: status code", resp.StatusCode)
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}

	//get encoding of the html
	encode := determineEncoding(resp.Body)

	//手工转码
	//utf8reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8reader := transform.NewReader(resp.Body, encode.NewDecoder())

	return ioutil.ReadAll(utf8reader)
}
