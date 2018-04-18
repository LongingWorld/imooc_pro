package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding  {
	bytes, e := r.Peek(1024)
	if e != nil {
		log.Printf("Fetcher error:%v",e)
		return unicode.UTF8  //return default encoding.Encoding
	}
	encoding,_,_ := charset.DetermineEncoding(bytes, "")
	return encoding
}


var rateLimit = time.Tick(100*time.Millisecond)
func FetcherURL(url string) ([]byte ,error) {
	<-rateLimit
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error: status code", resp.StatusCode)
		return nil,fmt.Errorf("wrong status code:%d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	//get encoding of the html
	encode := determineEncoding(bodyReader)

	//手工转码
	//utf8reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8reader := transform.NewReader(bodyReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8reader)
}
