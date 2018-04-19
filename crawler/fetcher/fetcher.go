package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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


var rateLimit = time.Tick(10*time.Millisecond)

//User agent Proxy
var userAgents = []string {
	"Mozilla/5.0(Macintosh;U;IntelMacOSX10_6_8;en-us)AppleWebKit/534.50(KHTML,likeGecko)Version/5.1Safari/534.50",
	"Mozilla/5.0(Windows;U;WindowsNT6.1;en-us)AppleWebKit/534.50(KHTML,likeGecko)Version/5.1Safari/534.50",
	"Mozilla/5.0(compatible;MSIE9.0;WindowsNT6.1;Trident/5.0)",	// IE9
	"Mozilla/4.0(compatible;MSIE8.0;WindowsNT6.0;Trident/4.0)",	// IE8
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT6.0)",	// IE7
	"Mozilla/4.0(compatible;MSIE6.0;WindowsNT5.1)",	// IE6
	"Mozilla/5.0(Macintosh;IntelMacOSX10.6;rv:2.0.1)Gecko/20100101Firefox/4.0.1",
	"Mozilla/5.0(WindowsNT6.1;rv:2.0.1)Gecko/20100101Firefox/4.0.1",
	"Opera/9.80(Macintosh;IntelMacOSX10.6.8;U;en)Presto/2.8.131Version/11.11",
	"Opera/9.80(WindowsNT6.1;U;en)Presto/2.8.131Version/11.11",
	"Mozilla/5.0(Macintosh;IntelMacOSX10_7_0)AppleWebKit/535.11(KHTML,likeGecko)Chrome/17.0.963.56Safari/535.11",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;Maxthon2.0)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;TencentTraveler4.0)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;TheWorld)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;Trident/4.0;SE2.XMetaSr1.0;SE2.XMetaSr1.0;.NETCLR2.0.50727;SE2.XMetaSr1.0)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;360SE)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1;AvantBrowser)",
	"Mozilla/4.0(compatible;MSIE7.0;WindowsNT5.1)",
	"Mozilla/5.0(iPhone;U;CPUiPhoneOS4_3_3likeMacOSX;en-us)AppleWebKit/533.17.9(KHTML,likeGecko)Version/5.0.2Mobile/8J2Safari/6533.18.5",
	"Mozilla/5.0(iPod;U;CPUiPhoneOS4_3_3likeMacOSX;en-us)AppleWebKit/533.17.9(KHTML,likeGecko)Version/5.0.2Mobile/8J2Safari/6533.18.5",
	"Mozilla/5.0(iPad;U;CPUOS4_3_3likeMacOSX;en-us)AppleWebKit/533.17.9(KHTML,likeGecko)Version/5.0.2Mobile/8J2Safari/6533.18.5",
	"Mozilla/5.0(Linux;U;Android2.3.7;en-us;NexusOneBuild/FRF91)AppleWebKit/533.1(KHTML,likeGecko)Version/4.0MobileSafari/533.1",
	"MQQBrowser/26Mozilla/5.0(Linux;U;Android2.3.7;zh-cn;MB200Build/GRJ22;CyanogenMod-7)AppleWebKit/533.1(KHTML,likeGecko)Version/4.0MobileSafari/533.1",
	"Opera/9.80(Android2.3.4;Linux;OperaMobi/build-1107180945;U;en-GB)Presto/2.8.149Version/11.10",
	"Mozilla/5.0(Linux;U;Android3.0;en-us;XoomBuild/HRI39)AppleWebKit/534.13(KHTML,likeGecko)Version/4.0Safari/534.13",
	"Mozilla/5.0(BlackBerry;U;BlackBerry9800;en)AppleWebKit/534.1+(KHTML,likeGecko)Version/6.0.0.337MobileSafari/534.1+",
	"Mozilla/5.0(hp-tablet;Linux;hpwOS/3.0.0;U;en-US)AppleWebKit/534.6(KHTML,likeGecko)wOSBrowser/233.70Safari/534.6TouchPad/1.0",
	"Mozilla/4.0(compatible;MSIE6.0;)Opera/UCWEB7.0.2.37/28/999",
}


func FetcherURL(url string) ([]byte ,error) {
	<-rateLimit
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil,err
	//}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	//
	uA_nums := rand.Intn(len(userAgents))
	//log.Printf("nums: %d Agent:%s",uA_nums,userAgents[uA_nums])
	request.Header.Set("User-Agent",userAgents[uA_nums])

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
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
