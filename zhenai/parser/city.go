package parser

import (
	"crawler/engine"
	"regexp"
)

var  cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
func ParseCity(contens []byte)engine.ParseResult{
	matches:=cityRe.FindAllSubmatch(contens,-1)

	result :=engine.ParseResult{}
	for _,m:=range matches{
		name :=string(m[2])
		result.Items=append(result.Items,"User "+name)
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
			return 	ParseProfile(c,name)
			},
		})
	}
	return result
}