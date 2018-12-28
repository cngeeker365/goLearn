package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityReg = `<a href="http://album.zhenai.com/u/([0-9]+)" target="_blank">([^>]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityReg)
	matches := reg.FindAllSubmatch(contents, -1)
	
	result := engine.ParseResult{}
	
	for _,m := range matches {
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}