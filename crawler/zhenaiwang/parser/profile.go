package parser

import (
	"regexp"
	"strconv"

	"GitHub/imooc_pro/crawler/engine"
	"GitHub/imooc_pro/crawler/model"

)

var  AgeRegexp  = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var GenderRegexp = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var MarriageRegexp = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var HeightRegexp = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([0-9]+)CM</span></td>`)
var WeightRegexp = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var IncomeRegexp = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var EduRegexp = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var OccupationReg = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var HokouRegexp = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var XinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var HouseRegexp = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var CarRegexp = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var GuessRegexp = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var UrlRegexp = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)


func ParseProfile(contents []byte,url string,name string) engine.ParserResult  {
	profile := model.Profile{}
	profile.Name =name
	profile.Age = extractInt(contents,AgeRegexp)
	profile.Gender = extractString(contents,GenderRegexp)
	profile.Marriage = extractString(contents,MarriageRegexp)
	profile.Height = extractInt(contents,HeightRegexp)
	profile.Weight = extractInt(contents,WeightRegexp)
	profile.Income = extractString(contents,IncomeRegexp)
	profile.Education = extractString(contents,EduRegexp)
	profile.Occupation = extractString(contents,OccupationReg)
	profile.Hokou = extractString(contents,HokouRegexp)
	profile.Xinzuo = extractString(contents,XinzuoReg)
	profile.House = extractString(contents,HouseRegexp)
	profile.Car = extractString(contents,CarRegexp)

	//result := engine.ParserResult{}
	//result.Items = append(result.Items,profile)
	result := engine.ParserResult{
		Items:[]engine.Item{
			{
				URL:url,
				Type:"zhenai",
				Id:extractString([]byte(url),UrlRegexp),
				Payload:profile,
			},
		},
	}

	matchs := GuessRegexp.FindAllSubmatch(contents, -1)
	for _,m :=range matchs  {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests,engine.Request{
			URL:url,
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return ParseProfile(bytes,url,name)
			},
		})
	}

	return result
}

func extractString(contents []byte,reg *regexp.Regexp) string  {
	re := reg.FindSubmatch(contents)
	if len(re) >=2{
		//fmt.Printf("contents: %s\n",string(re[1]))
		return string(re[1])
	}else {
		return ""
	}

}

func extractInt(contents []byte,reg *regexp.Regexp) int  {
	re := reg.FindSubmatch(contents)
	//fmt.Printf("cnts :%s\n",string(re[1]))
	if len(re) >=2{
		ints ,err :=strconv.Atoi(string(re[1]))
		if err == nil{
			//fmt.Printf("numbers:%d\n",ints)
			return ints
		}else{
			return 0
		}

	}else {
		return 0
	}

}
